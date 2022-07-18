package book

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"

	"Project/entities"
	"Project/service"

	"github.com/gorilla/mux"

	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"
)

type book struct {
	book service.Book
}

func TestPostBook(t *testing.T) {
	testCases := []struct {
		desc        string
		req         entities.Books
		resPostBook entities.Books
		errPostBook error
		expected    int
	}{
		{"success:posted", entities.Books{BookID: 1, AuthorID: 1, Title: "titan", Publications: "penguin", PublishedDate: "12/04/1998",
			Author: entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "sk"}}, entities.Books{BookID: 1, AuthorID: 1, Title: "titan", Publications: "penguin", PublishedDate: "12/04/1998",
			Author: entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "sk"}}, nil, http.StatusCreated},
		{"failure:invalid book id", entities.Books{BookID: -1, AuthorID: 1, Title: "titan", Publications: "penguin", PublishedDate: "12/04/1998",
			Author: entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "sk"}}, entities.Books{}, nil, http.StatusBadRequest},
		{"failure:missing title", entities.Books{BookID: 2, AuthorID: 1, Title: "", Publications: "penguin", PublishedDate: "12/04/1998",
			Author: entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "sk"}}, entities.Books{}, nil, http.StatusBadRequest},
		{"failure:invalid author id", entities.Books{BookID: 3, AuthorID: 1, Title: "jgg", Publications: "penguin", PublishedDate: "12/04/1998",
			Author: entities.Author{AuthorID: -1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "sk"}}, entities.Books{}, nil, http.StatusBadRequest},
		{"failure:error in svc post", entities.Books{BookID: 11, AuthorID: 1, Title: "titan", Publications: "penguin", PublishedDate: "07/04/2010",
			Author: entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "sk"}}, entities.Books{}, errors.New("error"), http.StatusBadRequest},
	}
	ctrl := gomock.NewController(t)
	mockServiceBook := service.NewMockBook(ctrl)
	mock := New(mockServiceBook)
	for _, tc := range testCases {
		body, err := json.Marshal(tc.req)
		if err != nil {
			log.Printf("error marshaling %v", err)
		}
		req := httptest.NewRequest(http.MethodPost, "/book", bytes.NewBuffer(body))
		w := httptest.NewRecorder()
		mockServiceBook.EXPECT().PostBook(context.TODO(), tc.req).Return(tc.resPostBook, tc.errPostBook).AnyTimes()
		//mockBook := New(mockBookServices{})
		//mockBook.PostBook(w, req)
		mock.PostBook(w, req)
		res := w.Result().StatusCode

		assert.Equal(t, res, tc.expected)
	}
}

func TestGetAll(t *testing.T) {
	testCases := []struct {
		desc          string
		title         string
		includeAuthor string
		errSvcGetAll  error
		response      []entities.Books
	}{
		{"success:fetched all", "", "false", nil, []entities.Books{{BookID: 1, AuthorID: 1, Title: "deciding decade", Publications: "penguin", PublishedDate: "20/03/2010",
			Author: entities.Author{}},
			{BookID: 2, AuthorID: 1, Title: "wrath", Publications: "penguin", PublishedDate: "20/08/2018",
				Author: entities.Author{}}},
		},
		{"success:fetched all with title", "titan", "false", nil, []entities.Books{{BookID: 2, AuthorID: 1, Title: "titan", Publications: "penguin", PublishedDate: "20/03/2010",
			Author: entities.Author{}},
			{BookID: 2, AuthorID: 1, Title: "titan", Publications: "penguin", PublishedDate: "20/08/2018",
				Author: entities.Author{}}},
		},
		{"success:fetched all with author included", "titan", "true", nil, []entities.Books{{BookID: 2, AuthorID: 1, Title: "titan", Publications: "penguin", PublishedDate: "20/03/2010",
			Author: entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "tony"}},
			{BookID: 2, AuthorID: 1, Title: "wrath", Publications: "penguin", PublishedDate: "20/08/2018",
				Author: entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "tony"}}},
		},
		{"success:fetched all with title", "", "", errors.New("error"), []entities.Books{{}}},
	}
	ctrl := gomock.NewController(t)
	mockServiceBook := service.NewMockBook(ctrl)
	mock := New(mockServiceBook)

	for i, tc := range testCases {
		req := httptest.NewRequest(http.MethodGet, "localhost:8000/book?"+"title="+tc.title+"&"+"includeAuthor="+tc.includeAuthor, nil)
		w := httptest.NewRecorder()
		mockServiceBook.EXPECT().GetAllBook(context.TODO(), tc.title, tc.includeAuthor).Return(tc.response, tc.errSvcGetAll).AnyTimes()
		mock.GetAllBook(w, req)
		data, err := io.ReadAll(w.Body)
		if err != nil {

			return
		}
		var book []entities.Books
		err = json.Unmarshal(data, &book)
		if err != nil {

			return
		}
		//assert.Equal(t, book, tc.response)
		if !(reflect.DeepEqual(book, tc.response)) {
			t.Errorf("test:%d got:%v  want:%v", i, book, tc.response)
		}
	}
}

func TestPutBook(t *testing.T) {
	testcases := []struct {
		desc       string
		endpoint   int
		body       entities.Books
		resPutBook entities.Books
		errPutBook error
		expected   int
	}{
		{"success:book Updated", 1, entities.Books{BookID: 1, AuthorID: 1, Title: "deciding decade",
			Publications: "penguin", PublishedDate: "20/03/2010", Author: entities.Author{}}, entities.Books{BookID: 1, AuthorID: 1, Title: "deciding decade",
			Publications: "penguin", PublishedDate: "20/03/2010", Author: entities.Author{}}, nil, http.StatusOK},
		{"failure:invalid book id", 4, entities.Books{BookID: -4, AuthorID: 1, Title: "deciding decade",
			Publications: "penguin", PublishedDate: "20/03/2010", Author: entities.Author{}}, entities.Books{}, nil, http.StatusBadRequest},
		{"failure:missing title", 5, entities.Books{BookID: 4, AuthorID: 1, Title: "",
			Publications: "penguin", PublishedDate: "20/03/2010", Author: entities.Author{}}, entities.Books{}, nil, http.StatusBadRequest},
		{"failure:invalid author id", 6, entities.Books{BookID: 4, AuthorID: -1, Title: "hf",
			Publications: "penguin", PublishedDate: "20/03/2010", Author: entities.Author{}}, entities.Books{}, nil, http.StatusBadRequest},
		{"failure:invalid publish date", 7, entities.Books{BookID: 4, AuthorID: 1, Title: "jhgf",
			Publications: "penguin", PublishedDate: "-7/03/2010", Author: entities.Author{}}, entities.Books{}, nil, http.StatusBadRequest},
		{"failure:invalid publish date", 8, entities.Books{BookID: 4, AuthorID: 1, Title: "hv",
			Publications: "penguin", PublishedDate: "34/03/2010", Author: entities.Author{}}, entities.Books{}, nil, http.StatusBadRequest},
		{"failure:invalid publish month", 9, entities.Books{BookID: 4, AuthorID: 1, Title: "jhvgcf",
			Publications: "penguin", PublishedDate: "34/-3/2010", Author: entities.Author{}}, entities.Books{}, nil, http.StatusBadRequest},
		{"failure:invalid publish month", 10, entities.Books{BookID: 4, AuthorID: 1, Title: "jhvvh",
			Publications: "penguin", PublishedDate: "34/13/2010", Author: entities.Author{}}, entities.Books{}, nil, http.StatusBadRequest},
		{"failure:invalid publish year", 11, entities.Books{BookID: 4, AuthorID: 1, Title: "jhcfg",
			Publications: "penguin", PublishedDate: "34/13/1770", Author: entities.Author{}}, entities.Books{}, nil, http.StatusBadRequest},
		{"failure:invalid publish year", 12, entities.Books{BookID: 4, AuthorID: 1, Title: "hvtf",
			Publications: "penguin", PublishedDate: "34/13/2045", Author: entities.Author{}}, entities.Books{}, nil, http.StatusBadRequest},
		{"failure:error put book", 13, entities.Books{BookID: 5, AuthorID: 1, Title: "jhfyg",
			Publications: "penguin", PublishedDate: "04/11/2010", Author: entities.Author{}}, entities.Books{}, errors.New("error"), http.StatusBadRequest},
		{"failure:err strconv ", -14, entities.Books{BookID: 6, AuthorID: 1, Title: "jkghfgh",
			Publications: "penguin", PublishedDate: "04/03/2010", Author: entities.Author{}}, entities.Books{}, nil, http.StatusBadRequest},
	}
	ctrl := gomock.NewController(t)
	mockServiceBook := service.NewMockBook(ctrl)
	mock := New(mockServiceBook)

	for i, tc := range testcases {
		book, err := json.Marshal(tc.body)
		if err != nil {
			log.Printf("error encoding %v", err)
		}

		req := httptest.NewRequest(http.MethodPut, "localhost:8000/book/{id}"+strconv.Itoa(tc.endpoint), bytes.NewReader(book))
		w := httptest.NewRecorder()
		req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(tc.endpoint)})
		//mockBook := New(mockBookServices{})
		//mockBook.PutBook(w, req)
		mockServiceBook.EXPECT().PutBook(context.TODO(), tc.endpoint, tc.body).Return(tc.resPutBook, tc.errPutBook).AnyTimes()
		mock.PutBook(w, req)
		res := w.Result()

		//assert.Equal(t, tc.expected, res.StatusCode)
		if res.StatusCode != tc.expected {
			t.Errorf("TestCase:%v got:%v want:%v", i, res.StatusCode, tc.expected)
		}

	}
}

func TestDeleteBook(t *testing.T) {
	testcases := []struct {
		desc          string
		target        int
		resDeleteBook int64
		errDeleteBook error
		expected      int
	}{
		{"valid id", 4, 1, nil, http.StatusNoContent},
		{"invalid id", -4, 0, nil, http.StatusBadRequest},
		{"valid id", 5, 0, errors.New("error "), http.StatusBadRequest},
	}
	ctrl := gomock.NewController(t)
	mockBookServices := service.NewMockBook(ctrl)
	mock := New(mockBookServices)

	for _, tc := range testcases {
		id := strconv.Itoa(tc.target)
		req := httptest.NewRequest(http.MethodDelete, "https://localhost:8000/book/{id}"+id, nil)
		w := httptest.NewRecorder()
		req = mux.SetURLVars(req, map[string]string{"id": id})

		mockBookServices.EXPECT().DeleteBook(context.TODO(), tc.target).Return(tc.resDeleteBook, tc.errDeleteBook).AnyTimes()
		mock.DeleteBook(w, req)
		res := w.Result().StatusCode

		assert.Equal(t, res, tc.expected)
		//if res != tc.expected {
		//	t.Errorf("testCase:%v got:%v want:%v", i, res, tc.expected)
		//}

	}
}

func TestGetBookByID(t *testing.T) {
	testCases := []struct {
		desc       string
		id         int
		resGetBook entities.Books
		errGetBook error
		expected   int
	}{
		{"sucess:valid id fetched book", 1, entities.Books{BookID: 1, AuthorID: 1, Title: "titan", Publications: "penguin", PublishedDate: "01/07/1998", Author: entities.Author{}}, nil, http.StatusOK},
		{"failure:invalid id", -1, entities.Books{}, nil, http.StatusBadRequest},
		{"failure:error in get book", 2, entities.Books{}, errors.New("error"), http.StatusBadRequest},
	}
	ctrl := gomock.NewController(t)
	mockServiceBook := service.NewMockBook(ctrl)
	mock := New(mockServiceBook)

	for _, tc := range testCases {
		req := httptest.NewRequest(http.MethodGet, "localhost:8000/book"+strconv.Itoa(tc.id), nil)
		req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(tc.id)})
		w := httptest.NewRecorder()
		//mockBook := New(mockBookServices{})
		//mockBook.GetBookById(w, req)
		mockServiceBook.EXPECT().GetBookByID(context.TODO(), tc.id).Return(tc.resGetBook, tc.errGetBook).AnyTimes()
		mock.GetBookById(w, req)
		res := w.Result().StatusCode

		assert.Equal(t, res, tc.expected)
	}
}
