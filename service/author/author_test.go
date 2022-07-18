package author

import (
	"context"
	"errors"
	"testing"

	"Project/datastore"
	"Project/entities"

	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
)

func TestPostAuthor(t *testing.T) {
	testCases := []struct {
		desc         string
		req          entities.Author
		response     entities.Author
		id           int
		res          bool
		errExistence error
		errPost      error
	}{
		{"success:author does not exist", entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "Tony"},
			entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "Tony"}, 1, false, nil, nil},
		{"failure:author does not exist", entities.Author{AuthorID: 20, FirstName: "shiv", LastName: "kumar", DateOfBirth: "09/03/2000", PenName: "shiv"},
			entities.Author{}, 20, false, errors.New("author does not exist"), nil},
		{"failure:author does not exist", entities.Author{AuthorID: 21, FirstName: "shiv", LastName: "kumar", DateOfBirth: "09/03/2000", PenName: "shiv"},
			entities.Author{}, 21, false, nil, errors.New("error")},
		{"failure:author exist", entities.Author{AuthorID: 2, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "Tony"},
			entities.Author{}, 2, true, nil, nil},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAuthorDatastore := datastore.NewMockAuthor(ctrl)
	mock := New(mockAuthorDatastore)

	for i, tc := range testCases {
		//call when we need method from another only

		mockAuthorDatastore.EXPECT().CheckExistence(context.TODO(), tc.id).Return(tc.res, tc.errExistence).AnyTimes()
		mockAuthorDatastore.EXPECT().PostAuthor(context.TODO(), tc.req).Return(tc.response, tc.errPost).AnyTimes()

		res, _ := mock.PostAuthor(context.TODO(), tc.req)

		if res != tc.response {
			t.Errorf("testcase:%v got:%v want:%v", i, res, tc.response)
		}

	}
}

func TestPutAuthor(t *testing.T) {
	testCases := []struct {
		desc           string
		input          entities.Author
		expectedOutput entities.Author
		resExistence   bool
		errExistence   error
		errPutAuthor   error
	}{
		{"success:valid details", entities.Author{AuthorID: 1, FirstName: "shiv", LastName: "kumar", DateOfBirth: "06/05/2001", PenName: "shiv"},
			entities.Author{AuthorID: 1, FirstName: "shiv", LastName: "kumar", DateOfBirth: "06/05/2001", PenName: "shiv"}, true, nil, nil},
		{"failure:author does not exist", entities.Author{AuthorID: -2, FirstName: "shiv", LastName: "kumar", DateOfBirth: "09/03/2000", PenName: "shiv"},
			entities.Author{}, false, nil, nil},
		{"failure:error in existence checking", entities.Author{AuthorID: 2, FirstName: "shiv", LastName: "kumar", DateOfBirth: "09/03/2000", PenName: "shiv"},
			entities.Author{}, true, errors.New("invalid author details"), nil},
		{"failure:error in put author", entities.Author{AuthorID: 3, FirstName: "", LastName: "kumar", DateOfBirth: "09/03/2000", PenName: "shiv"},
			entities.Author{}, true, nil, errors.New("error")},
	}
	ctrl := gomock.NewController(t)
	mockAuthorDatastore := datastore.NewMockAuthor(ctrl)
	mock := New(mockAuthorDatastore)

	for _, tc := range testCases {
		mockAuthorDatastore.EXPECT().CheckExistence(context.TODO(), tc.input.AuthorID).Return(tc.resExistence, tc.errExistence).AnyTimes()
		mockAuthorDatastore.EXPECT().PutAuthor(context.TODO(), tc.input).Return(tc.expectedOutput, tc.errPutAuthor).AnyTimes()
		result, _ := mock.PutAuthor(context.TODO(), tc.input)
		assert.Equal(t, result, tc.expectedOutput)
	}
}
func TestDeleteAuthor(t *testing.T) {
	testcases := []struct {
		desc         string
		req          int
		response     int64 //no of rows affected
		res          bool
		errExistence error
		errDelete    error
	}{
		{"success:successfully deleted", 1, 1, true, nil, nil},
		{"failure:error in checking existence", 10, 0, true, errors.New("error"), nil},
		{"failure:error in checking existence", 11, 0, false, nil, nil},
	}
	ctrl := gomock.NewController(t)
	mockAuthorDatastore := datastore.NewMockAuthor(ctrl)
	mock := New(mockAuthorDatastore)

	for _, tc := range testcases {
		mockAuthorDatastore.EXPECT().CheckExistence(context.TODO(), tc.req).Return(tc.res, tc.errExistence).AnyTimes()
		mockAuthorDatastore.EXPECT().DeleteAuthor(context.TODO(), tc.req).Return(tc.response, tc.errDelete).AnyTimes()
		res, _ := mock.DeleteAuthor(context.TODO(), tc.req)
		assert.Equal(t, res, tc.response)
	}
}
