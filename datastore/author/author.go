package author

import (
	"context"
	"database/sql"
	"log"

	"Project/entities"
)

type author struct {
	db *sql.DB
}

func New(db *sql.DB) author {
	return author{db: db}
}

// CheckExistence :checking existence of author with given id
func (a author) CheckExistence(ctx context.Context, id int) (bool, error) {
	row, err := a.db.Query("select * from Author where AuthorId=?", id)
	if err != nil || !row.Next() {
		return false, err
	}

	return true, nil

}

// PostAuthor :inserting author into database
func (a author) PostAuthor(ctx context.Context, author entities.Author) (entities.Author, error) {
	_, err := a.db.ExecContext(ctx, "insert into Author(AuthorId,FirstName,LastName,DateOfBirth,PenName) values(?,?,?,?,?)",
		author.AuthorID, author.FirstName, author.LastName, author.DateOfBirth, author.PenName)
	if err != nil {

		return entities.Author{}, err
	}

	return author, nil
}

// DeleteAuthor :deleting author details from database
func (a author) DeleteAuthor(ctx context.Context, id int) (int64, error) {

	res, err := a.db.Exec("delete from Author where AuthorId=?", id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {

		return 0, err
	}
	return rowsAffected, nil

}

// PutAuthor :updating author details in database
func (a author) PutAuthor(ctx context.Context, author entities.Author) (entities.Author, error) {

	_, err := a.db.ExecContext(ctx, "update Author set FirstName=?,LastName=?,DateOfBirth=?,PenName=? where AuthorId=?",
		author.FirstName, author.LastName, author.DateOfBirth, author.PenName, author.AuthorID)
	if err != nil {
		return entities.Author{}, err
	}

	return author, nil
}

// FetchingAuthor :fetching author details from the database
func (a author) FetchingAuthor(ctx context.Context, id int) (int, entities.Author) {

	Row := a.db.QueryRow("SELECT * FROM Author where AuthorId=?", id)

	var author entities.Author
	if err := Row.Scan(&author.AuthorID, &author.FirstName, &author.LastName, &author.DateOfBirth, &author.PenName); err != nil {

		return 0, entities.Author{}

	}
	return author.AuthorID, author
}

// IncludeAuthor :fetching author details and using it to include with book detail
func (a author) IncludeAuthor(ctx context.Context, id int) (entities.Author, error) {
	Row := a.db.QueryRow("select * from Author where authorId=?", id)

	var author entities.Author

	if err := Row.Scan(&author.AuthorID, &author.FirstName, &author.LastName, &author.DateOfBirth, &author.PenName); err != nil {
		log.Printf("failed: %v\n", err)
		return entities.Author{}, err
	}

	return author, nil
}
