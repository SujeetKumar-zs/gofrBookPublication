package author

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"Project/entities"
	"Project/service"

	"github.com/gorilla/mux"

	"github.com/golang/mock/gomock"
)

type author struct {
	svc service.Author
}

func TestPostAuthor(t *testing.T) {
	testcases := []struct {
		desc         string
		target       string
		req          entities.Author
		response     entities.Author
		errIoRead    error
		errSvcAuthor error
		errMarshal   error
		expected     int
	}{
		{desc: "success:author posted", target: "author", req: entities.Author{
			AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "30/04/2001", PenName: "sk"}, response: entities.Author{
			AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "30/04/2001", PenName: "sk"},
			errIoRead: nil, errSvcAuthor: nil, errMarshal: nil, expected: http.StatusCreated},
		{"failure:invalid firstname", "author", entities.Author{
			AuthorID: 3, FirstName: "", LastName: "mrinal", DateOfBirth: "20/05/1990", PenName: "Dark horse"},
			entities.Author{}, nil, nil, nil, http.StatusBadRequest},
		{"failure:invalid id", "author", entities.Author{
			AuthorID: -3, FirstName: "minnal", LastName: "mrinal", DateOfBirth: "20/05/1990", PenName: "Dark horse"},
			entities.Author{}, nil, nil, nil, http.StatusBadRequest},

		{desc: "failure:error in svc author", target: "author", req: entities.Author{
			AuthorID: 5, FirstName: "shiv", LastName: "kumar", DateOfBirth: "30/04/2001", PenName: "sk"},
			response: entities.Author{}, errIoRead: nil, errSvcAuthor: errors.New("error"), errMarshal: nil, expected: http.StatusBadRequest},
	}
	ctrl := gomock.NewController(t)
	mockBookService := service.NewMockAuthor(ctrl)
	mock := New(mockBookService)

	for i, tc := range testcases {

		author, err := json.Marshal(tc.req)
		if err != nil {
			fmt.Println("error:", err)
		}

		req := httptest.NewRequest(http.MethodPost, "localhost:8000/"+tc.target, bytes.NewBuffer(author))
		w := httptest.NewRecorder()
		mockBookService.EXPECT().PostAuthor(context.TODO(), tc.req).Return(tc.response, tc.errSvcAuthor).AnyTimes()
		mock.PostAuthor(w, req)
		res := w.Result().StatusCode

		if res != tc.expected {
			t.Errorf("testcase:%d desc:%s got:%v epected:%v", i, tc.desc, res, tc.expected)
		}
	}
}

func TestDeleteAuthor(t *testing.T) {
	testcases := []struct {
		desc            string
		id              int
		resDeleteAuthor int64
		errDeleteAuthor error
		expected        int
	}{
		{"valid authorId", 1, 1, nil, http.StatusNoContent},
		{"invalid authorId", -4, 0, nil, http.StatusBadRequest},
		{"invalid authorId", 5, 0, nil, http.StatusBadRequest},
		{"invalid authorId", 6, 1, errors.New("error"), http.StatusBadRequest},
	}
	ctrl := gomock.NewController(t)
	mockBookService := service.NewMockAuthor(ctrl)
	mock := New(mockBookService)

	for i, tc := range testcases {
		id := strconv.Itoa(tc.id)
		req := httptest.NewRequest("DELETE", "localhost:8000/author/{id}"+id, nil)
		req = mux.SetURLVars(req, map[string]string{"id": id})
		w := httptest.NewRecorder()
		//mockAuthor := New(mockAuthorServices{})
		//
		//mockAuthor.DeleteAuthor(w, req)

		mockBookService.EXPECT().DeleteAuthor(context.TODO(), tc.id).Return(tc.resDeleteAuthor, tc.errDeleteAuthor).AnyTimes()
		mock.DeleteAuthor(w, req)

		res := w.Result().StatusCode
		if tc.expected != res {
			t.Errorf("testcase:%d got:%v want:%v", i, res, tc.desc)
		}
	}
}

func TestPutAuthor(t *testing.T) {
	testCases := []struct {
		desc      string
		req       entities.Author
		resSvcPut entities.Author
		errSvcPut error
		expected  int
	}{
		{"success:updated ", entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "Tony"},
			entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "Tony"}, nil, http.StatusOK},
		{"failure:invalid id", entities.Author{AuthorID: -1, FirstName: "shiv", LastName: "kumar", DateOfBirth: "07/05/1999", PenName: "shiv"}, entities.Author{},
			nil, http.StatusBadRequest},
		{"failure:missing first name", entities.Author{AuthorID: 2, FirstName: "", LastName: "halwai", DateOfBirth: "07/03/1989", PenName: "halwai"}, entities.Author{},
			nil, http.StatusBadRequest},
		{"failure:err in svc put author", entities.Author{AuthorID: 3, FirstName: "shiva", LastName: "halwai", DateOfBirth: "07/03/1989", PenName: "halwai"}, entities.Author{},
			errors.New("error"), http.StatusBadRequest},
	}
	ctrl := gomock.NewController(t)
	mockAuthorServices := service.NewMockAuthor(ctrl)
	mock := New(mockAuthorServices)
	for i, tc := range testCases {
		author, err := json.Marshal(tc.req)
		if err != nil {
			log.Printf("error in marshaling:%v", err)
		}
		req := httptest.NewRequest(http.MethodPut, "localhost:8000/author", bytes.NewReader(author))
		w := httptest.NewRecorder()
		//mockAuthor := New(mockAuthorServices{})
		//mockAuthor.PutAuthor(w, req)
		mockAuthorServices.EXPECT().PutAuthor(context.TODO(), tc.req).Return(tc.resSvcPut, tc.errSvcPut).AnyTimes()
		mock.PutAuthor(w, req)
		res := w.Result().StatusCode
		if res != tc.expected {
			t.Errorf("testcase:%d desc:%s got:%v want:%v", i, tc.desc, res, tc.expected)
		}
	}
}
