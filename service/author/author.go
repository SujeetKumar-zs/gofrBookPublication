package author

import (
	"context"
	"errors"
	"log"

	"Project/datastore"
	"Project/entities"
)

type author struct {
	ds datastore.Author
}

func New(ds datastore.Author) author {
	return author{ds: ds}
}

// PostAuthor :posting author with checking existence of author in database
func (a author) PostAuthor(ctx context.Context, author entities.Author) (entities.Author, error) {
	res, err := a.ds.CheckExistence(ctx, author.AuthorID)
	if err != nil {
		log.Printf("error checking existence")
		return entities.Author{}, err
	}
	if res != true {
		auth, err := a.ds.PostAuthor(ctx, author)
		if err != nil {
			return entities.Author{}, err
		}
		return auth, nil
	}

	return entities.Author{}, errors.New("author exist cant post")
}

// PutAuthor :checking existence of author with id and updating details
func (a author) PutAuthor(ctx context.Context, id int, author entities.Author) (entities.Author, error) {
	res, err := a.ds.CheckExistence(ctx, id)
	if err != nil {
		log.Printf("error in checkexistence")
		return entities.Author{}, err
	}
	if res == true {
		auth, err := a.ds.PutAuthor(ctx, author)
		if err != nil {
			return entities.Author{}, err
		}
		return auth, nil
	}
	return entities.Author{}, errors.New("author does not exist,cant update")
}

// DeleteAuthor :checking existence of author with id and deleting author details
func (a author) DeleteAuthor(ctx context.Context, id int) (int64, error) {
	res, err := a.ds.CheckExistence(ctx, id)
	if err != nil {
		return 0, err
	}
	if res == true {
		return a.ds.DeleteAuthor(ctx, id)
	}
	return 0, err
}
