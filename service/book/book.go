package book

import (
	"context"
	"errors"
	"log"
	"strconv"
	"strings"

	"Project/datastore"
	"Project/entities"
)

type book struct {
	dsBook   datastore.Book
	dsAuthor datastore.Author
}

func New(dsBook datastore.Book, dsAuthor datastore.Author) book {
	return book{dsBook: dsBook, dsAuthor: dsAuthor}
}

// PostBook :checking existence of book and posting details of books
func (b book) PostBook(ctx context.Context, book entities.Books) (entities.Books, error) {
	if !CheckPublication(book.Publications) || !CheckPublishDate(book.PublishedDate) {
		return entities.Books{}, errors.New("invalid book details")
	}
	res, err := b.dsBook.CheckExistence(ctx, book.BookID)
	if err != nil {
		return entities.Books{}, err
	}
	if res == false {
		return b.dsBook.PostBook(ctx, book)
	}
	return entities.Books{}, errors.New("book exist cant post")
}

//GetAllBook :fetching all book ,fetching only with title ,including author details
func (b book) GetAllBook(ctx context.Context, title, includeAuthor string) ([]entities.Books, error) {
	var (
		books []entities.Books
		err   error
	)

	if title != "" {
		books, err = b.dsBook.GetBookByTitle(ctx, title)
		if err != nil {
			return []entities.Books{}, err
		}
	} else {
		books, err = b.dsBook.GetAllBook(ctx)
		if err != nil {
			return []entities.Books{}, err
		}
	}

	if err != nil {
		return []entities.Books{}, err
	}

	if includeAuthor == "true" {
		for i := range books {
			author, err2 := b.dsAuthor.IncludeAuthor(ctx, books[i].AuthorID)
			if err2 != nil {
				log.Print(err)
				return []entities.Books{}, err
			}

			books[i].Author = author
		}
	}

	return books, nil

}

// GetBookByID :checking existence and fetching book by id
func (b book) GetBookByID(ctx context.Context, id int) (entities.Books, error) {
	res, err := b.dsBook.CheckExistence(ctx, id)
	if err != nil {
		log.Printf("error checcking existence:%v", err)
		return entities.Books{}, err
	}

	if res == true {
		return b.dsBook.GetBookByID(ctx, id)
	}

	return entities.Books{}, errors.New("book does not exist")
}

// PutBook :checking existence and updating details of book
func (b book) PutBook(ctx context.Context, id int, book entities.Books) (entities.Books, error) {

	if !CheckPublication(book.Publications) || !CheckPublishDate(book.PublishedDate) {
		return entities.Books{}, errors.New("invalid book details")
	}

	res, err := b.dsBook.CheckExistence(ctx, id)
	if err != nil {
		log.Printf("error checcking existence:%v", err)
		return entities.Books{}, err
	}
	if res == true {
		return b.dsBook.PutBook(ctx, id, book)
	}

	return entities.Books{}, errors.New(" Book with given id does not exist")
}

// DeleteBook :checking existence and deleting book details
func (b book) DeleteBook(ctx context.Context, id int) (int64, error) {
	res, err := b.dsBook.CheckExistence(ctx, id)
	if err != nil {
		return 0, err
	}
	if res == true {
		return b.dsBook.DeleteBook(ctx, id)
	}
	return 0, errors.New("book does not exist")

}

// CheckPublishDate :validates the published date
func CheckPublishDate(PublishDate string) bool {
	p := strings.Split(PublishDate, "/")
	day, _ := strconv.Atoi(p[0])
	month, _ := strconv.Atoi(p[1])
	year, _ := strconv.Atoi(p[2])

	switch {
	case day < 0 || day > 31:
		return false
	case month < 0 || month > 12:
		return false
	case year > 2022 || year < 1880:
		return false
	}

	return true
}

// CheckPublication : validates particular publications only
func CheckPublication(publication entities.Publication) bool {
	strings.ToLower(string(publication))

	if !(publication == entities.Arihant || publication == entities.Scholastic || publication == entities.Penguin) {
		return false
	}
	return true
}
