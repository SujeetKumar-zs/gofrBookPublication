package author

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
		desc         string
		id           int
		rows         *sqlmock.Rows
		resExistence bool
		errExistence error
	}{
		{"success:exist", 1, sqlmock.NewRows([]string{"AuthorId", "FirstName", "LastName", "DateOfBirth", "PenName"}).
			AddRow(1, "sujeet", "kumar", "06/04/2001", "tony"), true, nil},
		{"failure:do not exist", 10, sqlmock.NewRows([]string{}), false, nil},
		{"failure:error in checking existence", 10, sqlmock.NewRows([]string{}), false, errors.New("error")},
	}
	for _, tc := range testCases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			log.Printf("error mocking:%v", err)
		}
		mock.ExpectQuery("select * from Author where AuthorId=?").WillReturnRows(tc.rows).WillReturnError(tc.errExistence)
		mockDB := New(db)
		res, err := mockDB.CheckExistence(context.TODO(), tc.id)
		if err != nil {
			log.Printf("err :%v", err)
		}
		assert.Equal(t, res, tc.resExistence)
	}
}

func TestPostAuthor(t *testing.T) {

	testcases := []struct {
		desc     string
		req      entities.Author
		response entities.Author
		err      error
	}{
		{"success:posted", entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "Tony"},
			entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "Tony"}, nil},
		{"failure:not posted", entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "Tony"},
			entities.Author{}, errors.New(" Duplicate entry '1' for key 'Author.PRIMARY'")},
	}
	for i, tc := range testcases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			log.Printf("error in mocking:%v", err)
		}

		mock.ExpectExec("insert into Author(AuthorId,FirstName,LastName,DateOfBirth,PenName) values(?,?,?,?,?)").
			WithArgs(tc.req.AuthorID, tc.req.FirstName, tc.req.LastName, tc.req.DateOfBirth, tc.req.PenName).
			WillReturnResult(sqlmock.NewResult(1, 1)).WillReturnError(tc.err)
		mockDB := New(db)
		res, err := mockDB.PostAuthor(context.TODO(), tc.req)
		fmt.Println(err)
		if res != tc.response && err != tc.err {
			t.Errorf("[testcase:%d] desc:%v actualError:%v expectedError:%v actualResult:%v expectedResult:%v", i, tc.desc, err, tc.err, res, tc.response)
		}

	}
}

func TestDeleteAuthor(t *testing.T) {
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

		mock.ExpectExec("delete from Author where AuthorId=?").WithArgs(tc.req).WillReturnResult(tc.res).WillReturnError(tc.err)
		mockDB := New(db)
		res, err := mockDB.DeleteAuthor(context.TODO(), tc.req)
		if res != tc.response && err != tc.err {
			t.Errorf("testcase:%d desc:%v actualoutput:%v actualerror:%v expectedOutput:%v expectederror:%v ", i, tc.desc, res, err, tc.response, tc.err)
		}
	}
}

func TestPutAuthor(t *testing.T) {
	testcases := []struct {
		desc           string
		req            entities.Author
		response       entities.Author
		rows           *sqlmock.Rows
		lastInsertedId int64
		rowsAffected   int64
		err            error
	}{
		{"success:updated successfully", entities.Author{AuthorID: 1, FirstName: "shiva", LastName: "kumar", DateOfBirth: "06/07/2000", PenName: "shiva"},
			entities.Author{AuthorID: 1, FirstName: "shiva", LastName: "kumar", DateOfBirth: "06/07/2000", PenName: "shiva"}, sqlmock.NewRows([]string{"1", "shiva", "kumar", "06/07/2000", "shiva"}), 1, 1, nil},
		{"failure:id does not exist", entities.Author{AuthorID: 9, FirstName: "shiva", LastName: "kumar", DateOfBirth: "06/07/2000", PenName: "shiva"},
			entities.Author{}, sqlmock.NewRows([]string{}), 0, 0, errors.New("sql: no rows in result set")},
	}

	for i, tc := range testcases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			log.Printf("error mocking:%v", err)
		}

		mock.ExpectExec("update Author set FirstName=?,LastName=?,DateOfBirth=?,PenName=? where AuthorId=?").
			WithArgs(tc.req.FirstName, tc.req.LastName, tc.req.DateOfBirth, tc.req.PenName, tc.req.AuthorID).WillReturnResult(sqlmock.NewResult(tc.lastInsertedId, tc.rowsAffected)).WillReturnError(tc.err)
		mockDB := New(db)
		res, err := mockDB.PutAuthor(context.TODO(), tc.req)

		if res != tc.response && err != tc.err {
			t.Errorf("testcase:%d desc:%v actualResult:%v actualError:%v expectedResponse:%v expectedError:%v", i, tc.desc, res, err, tc.response, tc.err)
		}
	}
}
func TestFetchingAuthor(t *testing.T) {
	testCases := []struct {
		desc     string
		id       int
		response entities.Author
		rows     *sqlmock.Rows
	}{
		{"success:author exist", 1, entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "tony"},
			sqlmock.NewRows([]string{"AuthorId", "FirstName", "LastName", "DateOfBirth", "PenName"}).AddRow(1, "sujeet", "kumar", "06/04/2001", "tony")},
		{"failure:author does not exist", 20, entities.Author{},
			sqlmock.NewRows([]string{}).AddRow()},
	}
	for _, tc := range testCases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			log.Printf("error :%v", err)
		}
		mock.ExpectQuery("SELECT * FROM Author where AuthorId=?").WithArgs(tc.id).WillReturnRows(tc.rows)
		mockDB := New(db)
		_, res := mockDB.FetchingAuthor(context.TODO(), tc.id)
		assert.Equal(t, res, tc.response)
	}
}

func TestIncludeAuthor(t *testing.T) {
	testCases := []struct {
		desc     string
		req      int
		response entities.Author
		row      *sqlmock.Rows
		err      error
	}{
		{"success:fetched author", 1, entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "tony"},
			sqlmock.NewRows([]string{"AuthorId", "FirstName", "LastName", "DateOfBirth", "PenName"}).AddRow(1, "sujeet", "kumar", "06/04/2001", "tony"), nil},
		{"failure:error scanning", 1, entities.Author{AuthorID: 1, FirstName: "sujeet", LastName: "kumar", DateOfBirth: "06/04/2001", PenName: "tony"},
			sqlmock.NewRows([]string{"AuthorId", "FirstName", "LastName", "DateOfBirth", "PenName"}).AddRow("abc", "sujeet", "kumar", "06/04/2001", "tony"), errors.New("error")},
	}
	for _, tc := range testCases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			log.Printf("error:%v", err)
		}
		mock.ExpectQuery("select * from Author where authorId=?").WithArgs(tc.req).WillReturnRows(tc.row).WillReturnError(tc.err)
		mockDB := New(db)
		res, err := mockDB.IncludeAuthor(context.TODO(), tc.req)
		if err != tc.err && res != tc.response {
			t.Errorf("error")
		}
	}
}
