package book

import (
	"context"
	"database/sql"
	"log"

	"Project/entities"
)

type book struct {
	db *sql.DB
}

func New(db *sql.DB) book {
	return book{db: db}
}

// CheckExistence :checking existence of book in database
func (b book) CheckExistence(ctx context.Context, id int) (bool, error) {
	row, err := b.db.Query("select * from Books where BookId=?", id)
	if err != nil || !row.Next() {
		return false, err
	}

	return true, nil
}

// PostBook :inserting book into database
func (b book) PostBook(ctx context.Context, book entities.Books) (entities.Books, error) {
	_, err := b.db.Exec("insert into Books values (?,?,?,?,?) ", book.BookID, book.AuthorID, book.Title, book.Publications, book.PublishedDate)
	if err != nil {
		log.Printf("error inserting:%v", err)
		return entities.Books{}, err
	}

	return book, nil
}

// GetAllBook : fetching all books from the database
func (b book) GetAllBook(ctx context.Context) ([]entities.Books, error) {
	Rows, err := b.db.Query("select  * from Books")

	if err != nil {
		log.Printf("failed for:%v", err)
		return []entities.Books{}, err
	}
	defer Rows.Close()

	var books []entities.Books

	for Rows.Next() {
		var book entities.Books
		err := Rows.Scan(&book.BookID, &book.AuthorID, &book.Title, &book.Publications, &book.PublishedDate)
		if err != nil {
			log.Printf("failed for:%v", err)
			return []entities.Books{}, err
		}
		books = append(books, book)
	}

	return books, nil
}

// GetBookByID : fetching book of a specific id from the database
func (b book) GetBookByID(ctx context.Context, id int) (entities.Books, error) {
	row, err := b.db.Query("select * from Books where BookId=?", id)
	if err != nil {

		return entities.Books{}, err
	}
	var newBook entities.Books
	for row.Next() {
		if err = row.Scan(&newBook.BookID, &newBook.AuthorID, &newBook.Title, &newBook.Publications, &newBook.PublishedDate); err != nil {
			return entities.Books{}, err
		}
	}

	return newBook, nil
}

// PutBook :updating book details into the database
func (b book) PutBook(ctx context.Context, id int, book entities.Books) (entities.Books, error) {

	_, err := b.db.Exec("update Books set BookId=? ,Title=?, Publications=?, PublishedDate=? where BookId=? ", book.BookID, book.Title, book.Publications, book.PublishedDate, id)
	if err != nil {
		log.Printf("error updating book %v", err)
		return entities.Books{}, err
	}
	return book, nil
}

// DeleteBook : deleting book of a specific id from the database
func (b book) DeleteBook(ctx context.Context, id int) (int64, error) {

	res, err := b.db.Exec("delete from Books where BookId=?", id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {

		return 0, err
	}
	return rowsAffected, nil
}

// GetBookByTitle : fetching all books with the provided title
func (b book) GetBookByTitle(ctx context.Context, title string) ([]entities.Books, error) {

	row, err := b.db.QueryContext(ctx, "select * from Books where Title=?", title)
	if err != nil {
		log.Printf("error :%v", err)
		return []entities.Books{}, err
	}
	var books []entities.Books
	for row.Next() {
		var book entities.Books
		if err := row.Scan(&book.BookID, &book.AuthorID, &book.Title, &book.Publications, &book.PublishedDate); err != nil {
			log.Printf("error :%v", err)
			return []entities.Books{}, err
		}
		books = append(books, book)
	}
	return books, err
}
