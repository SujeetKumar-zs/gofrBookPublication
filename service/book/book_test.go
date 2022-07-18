package book

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"Project/datastore"
	"Project/entities"

	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
)

var author = entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "tony"}

func TestPostBook(t *testing.T) {
	testCases := []struct {
		desc           string
		input          entities.Books
		expectedOutput entities.Books
		resExistence   bool
		errExistence   error
		errPostBook    error
	}{
		{"success:posted successfully", entities.Books{BookID: 1, AuthorID: 1, Title: "titan", Publications: "penguin", PublishedDate: "09/03/1990",
			Author: entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "tony"}},
			entities.Books{BookID: 1, AuthorID: 1, Title: "titan", Publications: "penguin", PublishedDate: "09/03/1990",
				Author: entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "tony"}}, false, nil, nil},
		{"success:posted success fully", entities.Books{BookID: 2, AuthorID: 1, Title: "titan", Publications: "penguin", PublishedDate: "09/03/1990",
			Author: entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "tony"}},
			entities.Books{}, true, nil, nil},
		{"failure:err checking existence", entities.Books{BookID: 3, AuthorID: 1, Title: "titan", Publications: "penguin", PublishedDate: "09/03/1990",
			Author: entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "tony"}},
			entities.Books{}, false, errors.New("error"), nil},
		{"failure:invalid publication", entities.Books{BookID: 4, AuthorID: 1, Title: "vbhjn", Publications: "kiran", PublishedDate: "12/04/1998",
			Author: entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "sk"}}, entities.Books{}, false, nil, nil},
		{"failure:invalid publish date", entities.Books{BookID: 5, AuthorID: 1, Title: "hdfvg", Publications: "penguin", PublishedDate: "-2/04/1998",
			Author: entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "sk"}}, entities.Books{}, false, nil, nil},
		{"failure:invalid publish date", entities.Books{BookID: 6, AuthorID: 1, Title: "jkgcdf", Publications: "penguin", PublishedDate: "45/04/1998",
			Author: entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "sk"}}, entities.Books{}, false, nil, nil},
		{"failure:invalid publish month", entities.Books{BookID: 7, AuthorID: 1, Title: "cdrfg", Publications: "penguin", PublishedDate: "07/-4/1998",
			Author: entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "sk"}}, entities.Books{}, false, nil, nil},
		{"failure:invalid publish month", entities.Books{BookID: 8, AuthorID: 1, Title: "jhvdx", Publications: "penguin", PublishedDate: "07/17/1998",
			Author: entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "sk"}}, entities.Books{}, false, nil, nil},
		{"failure:invalid publish year", entities.Books{BookID: 9, AuthorID: 1, Title: "nbv", Publications: "penguin", PublishedDate: "07/04/1790",
			Author: entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "sk"}}, entities.Books{}, false, nil, nil},
		{"failure:invalid publish year", entities.Books{BookID: 10, AuthorID: 1, Title: "jhf", Publications: "penguin", PublishedDate: "07/04/2045",
			Author: entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "sk"}}, entities.Books{}, false, nil, nil},
	}
	ctrl := gomock.NewController(t)
	mockBookDatastore := datastore.NewMockBook(ctrl)
	mockAuthorDatastore := datastore.NewMockAuthor(ctrl)
	mock := New(mockBookDatastore, mockAuthorDatastore)
	for i, tc := range testCases {
		mockBookDatastore.EXPECT().CheckExistence(context.TODO(), tc.input.BookID).Return(tc.resExistence, tc.errExistence).AnyTimes()
		mockBookDatastore.EXPECT().PostBook(context.TODO(), tc.input).Return(tc.expectedOutput, tc.errPostBook).AnyTimes()
		res, _ := mock.PostBook(context.TODO(), tc.input)
		//assert.Equal(t, res, tc.expectedOutput)
		if res != tc.expectedOutput {
			t.Errorf("testcase:%d got:%v want:%v", i, res, tc.expectedOutput)
		}

	}
}

func TestGetAllBook(t *testing.T) {
	testCases := []struct {
		desc             string
		title            string
		includeAuthor    string
		response         []entities.Books
		errGetByTitle    error
		errGetAll        error
		errIncludeAuthor error
	}{
		{"success:fetched successfully", "", " ", []entities.Books{
			{BookID: 1, AuthorID: 1, Title: "Titan", Publications: "penguin", PublishedDate: "16/07/1990",
				Author: entities.Author{}},
			{BookID: 2, AuthorID: 1, Title: "wrath", Publications: "scholastic", PublishedDate: "15/08/1995",
				Author: entities.Author{}},
		}, nil, nil, nil},
		{"success:fetched with titles", "titan", " ", []entities.Books{
			{BookID: 3, AuthorID: 1, Title: "Titan", Publications: "penguin", PublishedDate: "16/07/1990",
				Author: entities.Author{}},
			{BookID: 4, AuthorID: 1, Title: "Titan", Publications: "penguin", PublishedDate: "15/08/1995",
				Author: entities.Author{}},
		}, nil, nil, nil},
		{"success:fetched with author included ", "", "true", []entities.Books{
			{BookID: 1, AuthorID: 1, Title: "Titan", Publications: "penguin", PublishedDate: "16/07/1990",
				Author: entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "tony"}},
			{BookID: 2, AuthorID: 1, Title: "wrath", Publications: "scholastic", PublishedDate: "15/08/1995",
				Author: entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "tony"}},
		}, nil, nil, nil},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockBookDatastore := datastore.NewMockBook(ctrl)
	mockAuthorDatastore := datastore.NewMockAuthor(ctrl)
	mock := New(mockBookDatastore, mockAuthorDatastore)
	for i, tc := range testCases {
		mockBookDatastore.EXPECT().GetBookByTitle(context.TODO(), tc.title).Return(tc.response, tc.errGetByTitle).AnyTimes()
		mockBookDatastore.EXPECT().GetAllBook(context.TODO()).Return(tc.response, tc.errGetAll).AnyTimes()
		mockAuthorDatastore.EXPECT().IncludeAuthor(context.TODO(), author.AuthorID).Return(author, tc.errIncludeAuthor).AnyTimes()
		res, _ := mock.GetAllBook(context.TODO(), tc.title, tc.includeAuthor)
		fmt.Println(i)
		assert.Equal(t, res, tc.response)

	}
}

func TestGetAllBook1(t *testing.T) {

	testCases := []struct {
		desc             string
		title            string
		includeAuthor    string
		response         []entities.Books
		errGetByTitle    error
		errGetAll        error
		errIncludeAuthor error
	}{

		{"failure:err in get by title", "titan", "", []entities.Books{}, errors.New("error"),
			nil, nil},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockBookDatastore := datastore.NewMockBook(ctrl)
	mockAuthorDatastore := datastore.NewMockAuthor(ctrl)
	mock := New(mockBookDatastore, mockAuthorDatastore)
	for i, tc := range testCases {
		mockBookDatastore.EXPECT().GetBookByTitle(context.TODO(), tc.title).Return(tc.response, tc.errGetByTitle).AnyTimes()
		mockBookDatastore.EXPECT().GetAllBook(context.TODO()).Return(tc.response, tc.errGetAll).AnyTimes()
		mockAuthorDatastore.EXPECT().IncludeAuthor(context.TODO(), author.AuthorID).Return(author, tc.errIncludeAuthor).AnyTimes()
		res, _ := mock.GetAllBook(context.TODO(), tc.title, tc.includeAuthor)
		fmt.Println(i)
		assert.Equal(t, res, tc.response)

	}
}

func TestGetAllBook2(t *testing.T) {

	testCases := []struct {
		desc             string
		title            string
		includeAuthor    string
		response         []entities.Books
		errGetByTitle    error
		errGetAll        error
		errIncludeAuthor error
	}{

		{"failure:err in getAll", "", "", []entities.Books{}, nil,
			errors.New("error"), nil},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockBookDatastore := datastore.NewMockBook(ctrl)
	mockAuthorDatastore := datastore.NewMockAuthor(ctrl)
	mock := New(mockBookDatastore, mockAuthorDatastore)
	for i, tc := range testCases {
		mockBookDatastore.EXPECT().GetBookByTitle(context.TODO(), tc.title).Return(tc.response, tc.errGetByTitle).AnyTimes()
		mockBookDatastore.EXPECT().GetAllBook(context.TODO()).Return(tc.response, tc.errGetAll).AnyTimes()
		mockAuthorDatastore.EXPECT().IncludeAuthor(context.TODO(), author.AuthorID).Return(author, tc.errIncludeAuthor).AnyTimes()
		res, _ := mock.GetAllBook(context.TODO(), tc.title, tc.includeAuthor)
		fmt.Println(i)
		assert.Equal(t, res, tc.response)

	}
}
func TestGetAllBook3(t *testing.T) {

	testCases := []struct {
		desc             string
		title            string
		includeAuthor    string
		responseGetAll   []entities.Books
		response         []entities.Books
		errGetByTitle    error
		errGetAll        error
		errIncludeAuthor error
	}{

		{"failure:err include author", "", "true", []entities.Books{{BookID: 1, AuthorID: 1, Title: "Titan", Publications: "penguin", PublishedDate: "16/07/1990",
			Author: entities.Author{}},
			{BookID: 2, AuthorID: 1, Title: "wrath", Publications: "scholastic", PublishedDate: "15/08/1995",
				Author: entities.Author{}}}, []entities.Books{}, nil,
			nil, errors.New("error")},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockBookDatastore := datastore.NewMockBook(ctrl)
	mockAuthorDatastore := datastore.NewMockAuthor(ctrl)
	mock := New(mockBookDatastore, mockAuthorDatastore)
	for i, tc := range testCases {
		mockBookDatastore.EXPECT().GetBookByTitle(context.TODO(), tc.title).Return(tc.responseGetAll, tc.errGetByTitle).AnyTimes()
		mockBookDatastore.EXPECT().GetAllBook(context.TODO()).Return(tc.responseGetAll, tc.errGetAll).AnyTimes()
		mockAuthorDatastore.EXPECT().IncludeAuthor(context.TODO(), author.AuthorID).Return(author, tc.errIncludeAuthor).AnyTimes()
		res, _ := mock.GetAllBook(context.TODO(), tc.title, tc.includeAuthor)
		fmt.Println(i)
		assert.Equal(t, res, tc.response)

	}
}
func TestGetBookByID(t *testing.T) {
	testCases := []struct {
		desc         string
		input        int
		response     entities.Books
		resExistence bool
		errExistence error
		errGetById   error
	}{
		{"success:fetched successfully", 1, entities.Books{BookID: 1, AuthorID: 1, Title: "Titan", Publications: "penguin", PublishedDate: "16/07/1990",
			Author: entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "Kumar", DateOfBirth: "06/04/2001", PenName: "Tony"}}, true, nil, nil},
		{"failure:res existence is false ", 2, entities.Books{}, false, nil, nil},
		{"failure:res err in checking existence ", 3, entities.Books{}, false, errors.New("error"), nil},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockBookDatastore := datastore.NewMockBook(ctrl)
	mockAuthorDatastore := datastore.NewMockAuthor(ctrl)
	mock := New(mockBookDatastore, mockAuthorDatastore)
	for _, tc := range testCases {
		mockBookDatastore.EXPECT().CheckExistence(context.TODO(), tc.input).Return(tc.resExistence, tc.errExistence).AnyTimes()
		mockBookDatastore.EXPECT().GetBookByID(context.TODO(), tc.input).Return(tc.response, tc.errGetById).AnyTimes()
		res, _ := mock.GetBookByID(context.TODO(), tc.input)
		assert.Equal(t, res, tc.response)
	}
}

func TestPutBook(t *testing.T) {
	testCases := []struct {
		desc           string
		id             int
		input          entities.Books
		expectedOutput entities.Books
		resExistence   bool
		errExistence   error
		errPutBook     error
	}{
		{"success:updated successfully", 2, entities.Books{BookID: 2, AuthorID: 1, Title: "avatar", Publications: "arihant", PublishedDate: "10/08/1998",
			Author: entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "Tony"}},
			entities.Books{BookID: 2, AuthorID: 1, Title: "avatar", Publications: "arihant", PublishedDate: "10/08/1998",
				Author: entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "Tony"}}, true, nil, nil},
		{"resExistence is false", 3, entities.Books{BookID: 2, AuthorID: 1, Title: "avatar", Publications: "arihant", PublishedDate: "10/08/1998",
			Author: entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "Tony"}},
			entities.Books{}, false, nil, nil},
		{"error in existence", 4, entities.Books{BookID: 2, AuthorID: 1, Title: "avatar", Publications: "arihant", PublishedDate: "10/08/1998",
			Author: entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "Tony"}},
			entities.Books{}, true, errors.New("error"), nil},
		{"failure:invalid publication", 4, entities.Books{BookID: 4, AuthorID: 1, Title: "vbhjn", Publications: "kiran", PublishedDate: "12/04/1998",
			Author: entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "sk"}}, entities.Books{}, false, nil, nil},
		{"failure:invalid publish date", 5, entities.Books{BookID: 5, AuthorID: 1, Title: "hdfvg", Publications: "penguin", PublishedDate: "-2/04/1998",
			Author: entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "sk"}}, entities.Books{}, false, nil, nil},
		{"failure:invalid publish date", 6, entities.Books{BookID: 6, AuthorID: 1, Title: "jkgcdf", Publications: "penguin", PublishedDate: "45/04/1998",
			Author: entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "sk"}}, entities.Books{}, false, nil, nil},
		{"failure:invalid publish month", 7, entities.Books{BookID: 7, AuthorID: 1, Title: "cdrfg", Publications: "penguin", PublishedDate: "07/-4/1998",
			Author: entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "sk"}}, entities.Books{}, false, nil, nil},
		{"failure:invalid publish month", 8, entities.Books{BookID: 8, AuthorID: 1, Title: "jhvdx", Publications: "penguin", PublishedDate: "07/17/1998",
			Author: entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "sk"}}, entities.Books{}, false, nil, nil},
		{"failure:invalid publish year", 9, entities.Books{BookID: 9, AuthorID: 1, Title: "nbv", Publications: "penguin", PublishedDate: "07/04/1790",
			Author: entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "sk"}}, entities.Books{}, false, nil, nil},
		{"failure:invalid publish year", 10, entities.Books{BookID: 10, AuthorID: 1, Title: "jhf", Publications: "penguin", PublishedDate: "07/04/2045",
			Author: entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "sk"}}, entities.Books{}, false, nil, nil},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockBookDatastore := datastore.NewMockBook(ctrl)
	mockAuthorDatastore := datastore.NewMockAuthor(ctrl)
	mock := New(mockBookDatastore, mockAuthorDatastore)
	for _, tc := range testCases {
		mockBookDatastore.EXPECT().CheckExistence(context.TODO(), tc.id).Return(tc.resExistence, tc.errExistence).AnyTimes()
		mockBookDatastore.EXPECT().PutBook(context.TODO(), tc.id, tc.input).Return(tc.expectedOutput, tc.errPutBook).AnyTimes()
		res, _ := mock.PutBook(context.TODO(), tc.id, tc.input)
		assert.Equal(t, res, tc.expectedOutput)
	}
}

func TestDeleteBook(t *testing.T) {
	testCases := []struct {
		desc           string
		input          int
		expectedOutput int64
		resExistence   bool
		errExistence   error
		errDeleteBook  error
	}{
		{"success:deleted successfully", 2, 1, true, nil, nil},
		{"failure:err in existence checking", 3, 0, true, errors.New("error"), nil},
		{"failure:book do not exist", 4, 0, false, nil, nil},
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockBookDatastore := datastore.NewMockBook(ctrl)
	mockAuthorDatastore := datastore.NewMockAuthor(ctrl)
	mock := New(mockBookDatastore, mockAuthorDatastore)

	for _, tc := range testCases {
		mockBookDatastore.EXPECT().CheckExistence(context.TODO(), tc.input).Return(tc.resExistence, tc.errExistence).AnyTimes()
		mockBookDatastore.EXPECT().DeleteBook(context.TODO(), tc.input).Return(tc.expectedOutput, tc.errDeleteBook).AnyTimes()
		res, _ := mock.DeleteBook(context.TODO(), tc.input)
		assert.Equal(t, res, tc.expectedOutput)
	}
}
