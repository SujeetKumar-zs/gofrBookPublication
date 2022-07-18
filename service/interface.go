package service

import (
	"context"

	"Project/entities"
)

type Book interface {
	PostBook(ctx context.Context, b entities.Books) (entities.Books, error)
	GetAllBook(ctx context.Context, title, includeAuthor string) ([]entities.Books, error)
	GetBookByID(ctx context.Context, id int) (entities.Books, error)
	PutBook(ctx context.Context, id int, b entities.Books) (entities.Books, error)
	DeleteBook(ctx context.Context, id int) (int64, error)
}

type Author interface {
	PostAuthor(ctx context.Context, a entities.Author) (entities.Author, error)
	DeleteAuthor(ctx context.Context, id int) (int64, error)
	PutAuthor(ctx context.Context, id int, a entities.Author) (entities.Author, error)
}
