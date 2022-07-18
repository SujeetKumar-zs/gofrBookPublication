package author

import (
	"errors"
	"log"
	"strconv"

	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"Project/entities"
	"Project/service"
)

type authorHandler struct {
	svcAuthor service.Author
}

func New(svcAuthor service.Author) authorHandler {
	return authorHandler{svcAuthor: svcAuthor}
}

// PostAuthor :posting/inserting author
func (a authorHandler) PostAuthor(c *gofr.Context) (interface{}, error) {

	var author entities.Author

	if err := c.Bind(&author); err != nil {
		log.Printf("error:%v", err)
		return nil, err
	}

	if author.AuthorID <= 0 || author.FirstName == "" {
		return nil, errors.New("invalid author details")
	}

	_, err := a.svcAuthor.PostAuthor(c, author)
	if err != nil {
		return nil, err
	}

	return author, nil
}

//PutAuthor :update the author
func (a authorHandler) PutAuthor(c *gofr.Context) (interface{}, error) {
	id := c.PathParam("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	var author entities.Author

	if err := c.Bind(&author); err != nil {
		return nil, err
	}
	if author.AuthorID <= 0 || author.FirstName == "" || ID <= 0 {
		return nil, errors.New("invalid author details")
	}

	_, err = a.svcAuthor.PutAuthor(c, ID, author)
	if err != nil {
		return nil, err
	}
	return author, nil
}

// DeleteAuthor :delete author with specific id
func (a authorHandler) DeleteAuthor(c *gofr.Context) (interface{}, error) {
	id := c.PathParam("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("error:%v", err)
		return nil, err
	}
	if ID <= 0 {
		return nil, errors.New("invalid id")
	}
	res, err := a.svcAuthor.DeleteAuthor(c, ID)
	if err != nil {
		return nil, err
	}
	if res == 1 {
		return res, nil
	}
	return nil, errors.New("author does not exist")
}
