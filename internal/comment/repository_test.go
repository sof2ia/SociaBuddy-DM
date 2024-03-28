package comment

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"log"
	"reflect"
	"testing"
	"time"
)

type argGet struct {
	name     string
	output   []Comment
	hasError error
}

type argCreate struct {
	name     string
	newCom   Comment
	output   *Comment
	hasError error
}

type argID struct {
	name     string
	id       int
	output   *Comment
	hasError error
}

type argIDList struct {
	name     string
	id       int
	output   []Comment
	hasError error
}

type argDate struct {
	name     string
	date     time.Time
	output   []Comment
	hasError error
}

type argEdit struct {
	name      string
	editedCom Comment
	id        int
	output    *Comment
	hasError  error
}

type argDelete struct {
	name     string
	id       int
	hasError error
}

func TestCreateCom(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("the creation of mock is failed %v", err)
	}
	defer func(mockDB *sql.DB) {
		_ = mockDB.Close()

	}(mockDB)

	customDate := time.Now().In(time.Local)
	rep := NewRepository(mockDB)
	mock.ExpectExec("INSERT INTO Comment").WithArgs(2, 1, customDate, "content1").WillReturnResult(sqlmock.NewResult(1, 1))
	result := sqlmock.NewRows([]string{
		"ID", "IDPost", "IDUser", "DateComment", "Content",
	}).AddRow(1, 2, 1, customDate, "content1")
	mock.ExpectQuery("SELECT \\* FROM Comment WHERE ID = ?").WithArgs(1).WillReturnRows(result)
	test := []argCreate{
		{
			name: "CreateCom() is succeed",
			newCom: Comment{
				ID:          1,
				IDPost:      2,
				IDUser:      1,
				DateComment: customDate,
				Content:     "content1",
			},
			output: &Comment{
				ID:          1,
				IDPost:      2,
				IDUser:      1,
				DateComment: customDate,
				Content:     "content1",
			},
			hasError: nil,
		},
		{
			name: "CreateComment() when there is no result",
			newCom: Comment{
				ID:          1,
				IDPost:      2,
				IDUser:      1,
				DateComment: customDate,
				Content:     "content1",
			},
			output:   nil,
			hasError: errors.New("comment has not written"),
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			comments, err := rep.CreateCom(tt.newCom)
			log.Printf("comments: %+v, err: %+v", comments, err)
			if !reflect.DeepEqual(comments, tt.newCom) {
				fmt.Printf("comments: %s\n", reflect.TypeOf(comments))
				fmt.Printf("tt.newCom: %s\n", reflect.TypeOf(tt.newCom))
				fmt.Printf("tt.output: %s\n", reflect.TypeOf(tt.output))
				t.Fatalf("expected %+v, got %+v", tt.output, comments)

			}
			if (err != nil && tt.hasError == nil) || (err == nil && tt.hasError != nil) {
				t.Fatalf("expeced error %+v, got %+v", tt.hasError, err)
			}
		})
	}
}
