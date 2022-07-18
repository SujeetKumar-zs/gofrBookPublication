package datastore

import (
	"Project/entities"
	"context"
)

type Book interface {
	PostBook(ctx context.Context, b entities.Books) (entities.Books, error)
	GetAllBook(ctx context.Context) ([]entities.Books, error)
	GetBookByID(ctx context.Context, id int) (entities.Books, error)
	PutBook(ctx context.Context, id int, b entities.Books) (entities.Books, error)
	DeleteBook(ctx context.Context, id int) (int64, error)
	CheckExistence(ctx context.Context, id int) (bool, error)
	GetBookByTitle(ctx context.Context, title string) ([]entities.Books, error)
}

type Author interface {
	PostAuthor(ctx context.Context, a entities.Author) (entities.Author, error)
	DeleteAuthor(ctx context.Context, id int) (int64, error)
	PutAuthor(ctx context.Context, a entities.Author) (entities.Author, error)
	CheckExistence(ctx context.Context, id int) (bool, error)
	FetchingAuthor(ctx context.Context, id int) (int, entities.Author)
	IncludeAuthor(ctx context.Context, id int) (entities.Author, error)
}
