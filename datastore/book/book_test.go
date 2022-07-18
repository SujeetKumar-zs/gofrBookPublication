package book

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"log"
	"testing"

	"Project/entities"

	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCheckExistence(t *testing.T) {
	testCases := []struct {
		desc     string
		id       int
		rows     *sqlmock.Rows
		response bool
		err      error
	}{
		{"success:book exist", 1, sqlmock.NewRows([]string{"BookId", "AuthorId", "Title", "Publications", "PublishedDate"}).
			AddRow(1, 1, "titan", "penguin", "12/07/1998"), true, nil},
		{"failure:book does not exist", 10, sqlmock.NewRows([]string{}), false, nil},
	}
	for i, tc := range testCases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			log.Printf("error mocking:%v", err)
		}
		mock.ExpectQuery("select * from Books where BookId=?").WithArgs(tc.id).
			WillReturnRows(tc.rows).WillReturnError(tc.err)
		mockDB := New(db)
		res, err := mockDB.CheckExistence(context.TODO(), tc.id)
		if res != tc.response && err != tc.err {
			t.Errorf("testcase:%d gotResult:%v gotError:%v expectedResult:%v expectedError:%v", i, res, err, tc.response, tc.err)
		}
	}
}

func TestPostBook(t *testing.T) {
	testCases := []struct {
		desc           string
		req            entities.Books
		response       entities.Books
		lastInsertedId int64
		rowsAffected   int64
		err            error
	}{
		{"success:posted successfully", entities.Books{BookID: 1, AuthorID: 1, Title: "Titan", Publications: "penguin", PublishedDate: "16/07/1990",
			Author: entities.Author{}},
			entities.Books{BookID: 1, AuthorID: 1, Title: "Titan", Publications: "penguin", PublishedDate: "16/07/1990",
				Author: entities.Author{}}, 1, 1, nil,
		},
		{"success:posted successfully", entities.Books{BookID: 2, AuthorID: 1, Title: "wrath", Publications: "scholastic", PublishedDate: "15/08/1995",
			Author: entities.Author{}},
			entities.Books{BookID: 2, AuthorID: 1, Title: "wrath", Publications: "scholastic", PublishedDate: "15/08/1995",
				Author: entities.Author{}}, 2, 1, nil,
		},

		{"failure:duplicate entry", entities.Books{BookID: 2, AuthorID: 1, Title: "wrath", Publications: "scholastic", PublishedDate: "15/08/1995",
			Author: entities.Author{}},
			entities.Books{}, 0, 0, fmt.Errorf("Duplicate entry '2' for key 'PRIMARY' "),
		},
	}

	for i, tc := range testCases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			log.Printf("error in mocking:%v", err)
		}
		mock.ExpectExec("insert into Books values (?,?,?,?,?) ").WithArgs(tc.req.BookID, tc.req.AuthorID, tc.req.Title, tc.req.Publications, tc.req.PublishedDate).
			WillReturnResult(sqlmock.NewResult(1, 1)).WillReturnError(tc.err)
		mockDB := New(db)
		res, err := mockDB.PostBook(context.TODO(), tc.req)
		if res != tc.response && err != tc.err {
			t.Errorf("testcase:%d desc:%v actualResult:%v expectedResult:%v actualError:%v expectedError:%v", i, tc.desc, res, tc.response, err, tc.err)
		}
	}
}

func TestGetAllBook(t *testing.T) {
	testCases := []struct {
		desc     string
		response []entities.Books
		rows     *sqlmock.Rows
		err      error
	}{
		{"success:fetched all successfully", []entities.Books{
			{BookID: 1, AuthorID: 1, Title: "Titan", Publications: "penguin", PublishedDate: "16/07/1990",
				Author: entities.Author{}},
			{BookID: 2, AuthorID: 1, Title: "wrath", Publications: "scholastic", PublishedDate: "15/08/1995",
				Author: entities.Author{}},
		}, sqlmock.NewRows([]string{"BookId", "AuthorId", "Title", "Publications", "PublishedDate"}).AddRow(1, 1, "Titan", "penguin", "16/07/1990").
			AddRow(2, 1, "wrath", "scholastic", "15/08/1995"), nil},
		{"failure:error scanning", []entities.Books{}, sqlmock.NewRows([]string{"BookId", "AuthorId", "Title", "Publications", "PublishedDate"}).
			AddRow("abc", 1, "Titan", "penguin", "16/07/1990"), nil},
		{"failure:error select all", []entities.Books{}, sqlmock.NewRows([]string{}), errors.New("error")},
	}
	for _, tc := range testCases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			log.Printf("error mocking:%v", err)
		}

		mock.ExpectQuery("select  * from Books").WillReturnRows(tc.rows).WillReturnError(tc.err)
		mockDB := New(db)
		res, err := mockDB.GetAllBook(context.TODO())
		if err != nil {
			log.Printf("err:%v", err)
		}
		fmt.Println(res)
		fmt.Println(tc.response)
		assert.Equal(t, tc.response, res)

	}

}

func TestGetBookByID(t *testing.T) {
	testCases := []struct {
		desc           string
		req            int
		response       entities.Books
		row            *sqlmock.Rows
		lastInsertedId int64
		rowsAffected   int64
		err            error
	}{
		{"success:fetched successfully", 1, entities.Books{BookID: 1, AuthorID: 1, Title: "Titan", Publications: "penguin", PublishedDate: "16/07/1990",
			Author: entities.Author{}},
			sqlmock.NewRows([]string{"BookId", "AuthorId", "Title", "Publications", "PublishedDate"}).AddRow(1, 1, "Titan", "penguin", "16/07/1990"),
			0, 0, nil},
		{"error:scanning", 1, entities.Books{},
			sqlmock.NewRows([]string{"BookId", "AuthorId", "Title", "Publications", "PublishedDate"}).AddRow("abc", 1, "Titan", "penguin", "16/07/1990"),
			0, 0, nil},
		{"failure:id does not exist", 6, entities.Books{}, sqlmock.NewRows([]string{}), 0, 0, errors.New("sql: no rows in result set")},
	}

	for i, tc := range testCases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			log.Printf("error mocking:%v", err)
		}
		mock.ExpectQuery("select * from Books where BookId=?").WithArgs(tc.req).
			WillReturnRows(tc.row).WillReturnError(tc.err)
		mockDB := New(db)
		res, err := mockDB.GetBookByID(context.TODO(), tc.req)
		if res != tc.response && err != tc.err {
			t.Errorf("testcase:%d desc:%v actualResult:%v actualError:%v expectedResponse:%v expectedError:%v", i, tc.desc, res, err, tc.response, tc.err)
		}
	}
}

func TestPutBook(t *testing.T) {
	testCases := []struct {
		desc           string
		input          int
		req            entities.Books
		response       entities.Books
		lastInsertedId int64
		rowsAffected   int64
		err            error
	}{
		{"success:updated successfully", 2, entities.Books{BookID: 1, AuthorID: 1, Title: "avatar", Publications: "arihant", PublishedDate: "10/08/1998",
			Author: entities.Author{}}, entities.Books{BookID: 1, AuthorID: 1, Title: "avatar", Publications: "arihant", PublishedDate: "10/08/1998",
			Author: entities.Author{}}, 1, 1, nil},
		{"failure:id does not exist", 9, entities.Books{BookID: 9, AuthorID: 1, Title: "avatar", Publications: "arihant", PublishedDate: "10/08/1998",
			Author: entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "Tony"}},
			entities.Books{}, 0, 0, fmt.Errorf("id does not exist")},
	}
	for i, tc := range testCases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			log.Printf("error mocking:%v", err)
		}

		mock.ExpectExec("update Books set BookId=? ,Title=?, Publications=?, PublishedDate=? where BookId=? ").WithArgs(tc.req.BookID, tc.req.Title, tc.req.Publications, tc.req.PublishedDate, tc.input).
			WillReturnResult(sqlmock.NewResult(tc.lastInsertedId, tc.rowsAffected)).WillReturnError(tc.err)

		mockDB := New(db)

		res, err := mockDB.PutBook(context.TODO(), tc.input, tc.req)
		if res != tc.response && err != tc.err {
			t.Errorf("testcase:%d desc:%v actualResult:%v actualError:%v expectedResponse:%v expectedError:%v ", i, tc.desc, res, err, tc.response, tc.err)
		}
	}
}

func TestDeleteBook(t *testing.T) {
	testCases := []struct {
		desc     string
		req      int
		response int64
		res      driver.Result
		err      error
	}{
		{"success:deleted successfully", 2, 1, sqlmock.NewResult(0, 1), nil},

		{"failure:inserted id error", 4, 0, sqlmock.NewErrorResult(errors.New("error")), nil},
		{"failure id does not exist", 20, 1, sqlmock.NewResult(0, 0), errors.New("error")},
	}

	for i, tc := range testCases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			log.Printf("error mocking:%v", err)
		}

		mock.ExpectExec("delete from Books where BookId=?").WithArgs(tc.req).WillReturnResult(tc.res).WillReturnError(tc.err)
		mockDB := New(db)
		res, err := mockDB.DeleteBook(context.TODO(), tc.req)
		if res != tc.response && err != tc.err {
			t.Errorf("testcase:%d desc:%v actualoutput:%v actualerror:%v expectedOutput:%v expectederror:%v ", i, tc.desc, res, err, tc.response, tc.err)
		}
	}
}

func TestGetBookByTitle(t *testing.T) {
	testCases := []struct {
		desc     string
		title    string
		response []entities.Books
		rows     *sqlmock.Rows
		err      error
	}{
		{"sucessfully fetched all", "titan", []entities.Books{
			{BookID: 1, AuthorID: 1, Title: "titan", Publications: "penguin", PublishedDate: "12/06/1990", Author: entities.Author{}},
			{BookID: 2, AuthorID: 1, Title: "wrath", Publications: "scholastic", PublishedDate: "02/06/1980", Author: entities.Author{}},
		}, sqlmock.NewRows([]string{"BookId", "AuthorId", "Title", "Publications", "PublishedDate"}).
			AddRow(1, 1, "titan", "penguin", "12/06/1990").AddRow(2, 1, "wrath", "scholastic", "02/06/1980"), nil},
		{"failure:error query", "titan", []entities.Books{}, sqlmock.NewRows([]string{}), errors.New("error")},
		{"failure:error scanning", "wrath", []entities.Books{}, sqlmock.NewRows([]string{"BookId", "AuthorId", "Title", "Publications", "PublishedDate"}).
			AddRow("abc", 1, "titan", "penguin", "12/06/1990"), nil},
	}
	for _, tc := range testCases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			log.Printf("error :%v", err)
		}
		mock.ExpectQuery("select * from Books where Title=?").WithArgs(tc.title).WillReturnRows(tc.rows).WillReturnError(tc.err)
		mockDB := New(db)

		res, _ := mockDB.GetBookByTitle(context.TODO(), tc.title)

		assert.Equal(t, res, tc.response)
	}
}
