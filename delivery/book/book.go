package book

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"Project/entities"
	"Project/service"

	"github.com/gorilla/mux"
)

type bookHandler struct {
	svcBook service.Book
}

func New(book service.Book) bookHandler {
	return bookHandler{svcBook: book}
}

// PostBook :posting/inserting book
func (b bookHandler) PostBook(w http.ResponseWriter, req *http.Request) {

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("failed for:%v", err)
	}
	var book entities.Books
	err = json.Unmarshal(body, &book)
	if err != nil {
		log.Printf("error:%v", err)
		return
	}

	if book.BookID <= 0 || book.Title == "" || book.AuthorID <= 0 || book.Author.FirstName == "" || book.Author.AuthorID <= 0 || book.AuthorID != book.Author.AuthorID {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte("invalid book details"))
		if err != nil {
			log.Printf("error:%v", err)
			return
		}

		return
	}
	ctx := context.TODO()
	_, err = b.svcBook.PostBook(ctx, book)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			log.Printf("error:%v", err)
			return
		}

		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(body)
	if err != nil {
		log.Printf("error:%v", err)
		return
	}

}

// GetAllBook :fetching all books and book with titles only and with included author details
func (b bookHandler) GetAllBook(w http.ResponseWriter, req *http.Request) {
	title := req.URL.Query().Get("title")
	includeAuthor := req.URL.Query().Get("includeAuthor")
	ctx := req.Context()

	requestedBooks, err := b.svcBook.GetAllBook(ctx, title, includeAuthor)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	gotBooks, err := json.Marshal(requestedBooks)
	if err != nil {
		log.Printf("failed for:%v", err)
	}
	bytes.NewBuffer(gotBooks)

	_, err = w.Write(gotBooks)
	if err != nil {
		log.Printf("failed for:%v", err)
	}

}

// GetBookById :fetching book by specific id
func (b bookHandler) GetBookById(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)

	strings.ToLower(params["id"])
	id, err := strconv.Atoi(params["id"])
	if err != nil || id <= 0 {
		w.WriteHeader(http.StatusBadRequest)
	}
	ctx := context.TODO()
	res, err := b.svcBook.GetBookByID(ctx, id)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			log.Printf("error:%v", err)
			return
		}
		return
	}
	data, err := json.Marshal(res)
	if err != nil {

		log.Printf("error marshal:%v", err)
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	if err != nil {
		log.Printf("error:%v", err)
		return
	}

}

//PutBook :Updates the book by id
func (b bookHandler) PutBook(w http.ResponseWriter, req *http.Request) {

	body := req.Body
	params := mux.Vars(req)

	data, err := ioutil.ReadAll(body)
	if err != nil {
		log.Printf("failed for:%v", err)
		return
	}

	var book entities.Books
	err = json.Unmarshal(data, &book)
	if err != nil {
		log.Printf("error:%v", err)
		return
	}

	strings.ToLower(params["id"])
	id, err := strconv.Atoi(params["id"])
	if err != nil || id <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte("bad params: "))
		if err != nil {
			log.Printf("error:%v", err)
			return
		}
		return
	}

	if book.BookID <= 0 || book.Title == "" || book.AuthorID <= 0 || book.AuthorID != book.Author.AuthorID {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte("invalid book details"))
		if err != nil {
			log.Printf("error:%v", err)
			return
		}
		return
	}
	ctx := context.TODO()
	res, err := b.svcBook.PutBook(ctx, id, book)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			log.Printf("error:%v", err)
			return
		}
		return
	}
	data, err = json.Marshal(res)
	if err != nil {
		fmt.Println("marshal")
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(data)
	if err != nil {
		log.Printf("error:%v", err)
	}

}

// DeleteBook :deletes book by id
func (b bookHandler) DeleteBook(w http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)

	id, err := strconv.Atoi(params["id"])
	if err != nil || id <= 0 {

		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte("bad param "))
		if err != nil {
			log.Printf("error:%v", err)
			return
		}

		return
	}

	ctx := context.TODO()
	_, err = b.svcBook.DeleteBook(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			log.Printf("error:%v", err)
		}

	}

	w.WriteHeader(http.StatusNoContent)
}
