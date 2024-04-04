package comment

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"log"
	"reflect"
	"regexp"
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
	idPost   int
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
	idPost   int
	output   []Comment
	hasError error
}

type argEdit struct {
	name      string
	idPost    int
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
			name:   "CreateCom() is succeed",
			idPost: 2,
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
			name:   "CreateComment() when there is no result",
			idPost: 1,
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
			comments, err := rep.CreateCom(tt.newCom, tt.idPost)
			log.Printf("comments: %+v, err: %+v", comments, err)
			if !reflect.DeepEqual(comments, tt.output) {
				t.Fatalf("expected %+v, got %+v", tt.output, comments)
			}
			if (err != nil && tt.hasError == nil) || (err == nil && tt.hasError != nil) {
				t.Fatalf("expeced error %+v, got %+v", tt.hasError, err)
			}
		})
	}
}

func TestGetCom(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("the creation of mock is failed %v", err)
	}
	defer func(mockDB *sql.DB) {
		_ = mockDB.Close()

	}(mockDB)
	rep := NewRepository(mockDB)
	timeNow := time.Date(2023, 11, 13, 0, 0, 0, 0, time.Local)

	result := sqlmock.NewRows([]string{
		"ID", "IDPost", "IDUser", "DateComment", "Content",
	}).AddRow(1, 2, 1, timeNow, "content1")
	mock.ExpectQuery("SELECT \\* FROM Comment").WillReturnRows(result)
	test := []argGet{
		{
			name: "GetComments() is succeed",
			output: []Comment{
				{
					ID:          1,
					IDPost:      2,
					IDUser:      1,
					DateComment: timeNow,
					Content:     "content1",
				},
			},
			hasError: nil,
		},
		{
			name:     "GetComments() when there is no result",
			output:   nil,
			hasError: errors.New("no comments in database"),
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			comments, err := rep.GetCom()
			log.Printf("comments: %+v, err: %+v", comments, err)
			if !reflect.DeepEqual(comments, tt.output) {
				t.Fatalf("expected %+v, got %+v", tt.output, comments)
			}
			if (err != nil && tt.hasError == nil) || (err == nil && tt.hasError != nil) {
				t.Fatalf("expeced error %+v, got %+v", tt.hasError, err)
			}
		})
	}
}
func TestGetComByID(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("the creation of mock is failed %v", err)
	}
	defer func(mockDB *sql.DB) {
		_ = mockDB.Close()

	}(mockDB)
	rep := NewRepository(mockDB)
	timeNow := time.Date(2023, 11, 13, 0, 0, 0, 0, time.Local)

	result := sqlmock.NewRows([]string{
		"ID", "IDPost", "IDUser", "DateComment", "Content",
	}).AddRow(1, 2, 1, timeNow, "content1")
	mock.ExpectQuery("SELECT \\* FROM Comment WHERE ID = ?").WithArgs(1).WillReturnRows(result)
	test := []argID{
		{
			name: "GetCommentByID() is succeed",
			id:   1,
			output: &Comment{

				ID:          1,
				IDPost:      2,
				IDUser:      1,
				DateComment: timeNow,
				Content:     "content1",
			},
			hasError: nil,
		},
		{
			name:     "GetCommentByID() is failed",
			id:       3,
			output:   nil,
			hasError: errors.New("no comment with this id in database"),
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			comment, err := rep.GetComByID(tt.id)
			log.Printf("comments: %+v, err: %+v", comment, err)
			if !reflect.DeepEqual(comment, tt.output) {
				t.Fatalf("expected %+v, got %+v", tt.output, comment)
			}
			if (err != nil && tt.hasError == nil) || (err == nil && tt.hasError != nil) {
				t.Fatalf("expeced error %+v, got %+v", tt.hasError, err)
			}
		})
	}
}

func TestGetComByPostID(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("the creation of mock is failed %v", err)
	}
	defer func(mockDB *sql.DB) {
		_ = mockDB.Close()

	}(mockDB)
	rep := NewRepository(mockDB)
	timeNow := time.Date(2023, 11, 13, 0, 0, 0, 0, time.Local)

	result := sqlmock.NewRows([]string{
		"ID", "IDPost", "IDUser", "DateComment", "Content",
	}).AddRow(1, 2, 1, timeNow, "content1")
	mock.ExpectQuery("SELECT \\* FROM Comment WHERE IDPost = ?").WithArgs(2).WillReturnRows(result)
	test := []argIDList{
		{
			name: "GetComByPostID() is succeed",
			id:   2,
			output: []Comment{
				{
					ID:          1,
					IDPost:      2,
					IDUser:      1,
					DateComment: timeNow,
					Content:     "content1",
				},
			},
			hasError: nil,
		},
		{
			name:     "GetComByPostID() is failed",
			id:       1,
			output:   nil,
			hasError: errors.New("the post has no comments in database"),
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			comments, err := rep.GetComByPostID(tt.id)
			log.Printf("comments: %+v, err: %+v", comments, err)
			if !reflect.DeepEqual(comments, tt.output) {
				t.Fatalf("expected %+v, got %+v", tt.output, comments)
			}
			if (err != nil && tt.hasError == nil) || (err == nil && tt.hasError != nil) {
				t.Fatalf("expeced error %+v, got %+v", tt.hasError, err)
			}
		})
	}
}

func TestComByUserID(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("the creation of mock is failed %v", err)
	}
	defer func(mockDB *sql.DB) {
		_ = mockDB.Close()

	}(mockDB)
	rep := NewRepository(mockDB)
	timeNow := time.Date(2023, 11, 13, 0, 0, 0, 0, time.Local)

	result := sqlmock.NewRows([]string{
		"ID", "IDPost", "IDUser", "DateComment", "Content",
	}).AddRow(1, 2, 1, timeNow, "content1")
	mock.ExpectQuery("SELECT \\* FROM Comment WHERE IDUser = ?").WithArgs(1).WillReturnRows(result)
	test := []argIDList{
		{
			name: "GetComByUserID() is succeed",
			id:   1,
			output: []Comment{
				{
					ID:          1,
					IDPost:      2,
					IDUser:      1,
					DateComment: timeNow,
					Content:     "content1",
				},
			},
			hasError: nil,
		},
		{
			name:     "GetComByUserID() is failed",
			id:       3,
			output:   nil,
			hasError: errors.New("the user has no comments in database"),
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			comments, err := rep.GetComByUserID(tt.id)
			log.Printf("comments: %+v, err: %+v", comments, err)
			if !reflect.DeepEqual(comments, tt.output) {
				t.Fatalf("expected %+v, got %+v", tt.output, comments)
			}
			if (err != nil && tt.hasError == nil) || (err == nil && tt.hasError != nil) {
				t.Fatalf("expeced error %+v, got %+v", tt.hasError, err)
			}
		})
	}
}
func TestGetComByDate(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("the creation of mock is failed %v", err)
	}
	defer func(mockDB *sql.DB) {
		_ = mockDB.Close()

	}(mockDB)
	timeNow := time.Date(2023, 11, 13, 0, 0, 0, 0, time.Local)
	rep := NewRepository(mockDB)
	result := sqlmock.NewRows([]string{
		"ID", "IDPost", "IDUser", "DateComment", "Content",
	}).AddRow(1, 2, 1, timeNow, "content1")
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM Comment WHERE strftime('%Y-%m-%d', DateComment) = ? AND IDPost = ?")).WithArgs(time.Date(2023, 11, 13, 0, 0, 0, 0, time.Local).Format("2006-01-02"), 2).WillReturnRows(result)

	tests := []argDate{
		{name: "GetComByDate() is succeed",
			date:   time.Date(2023, 11, 13, 0, 0, 0, 0, time.Local),
			idPost: 2,
			output: []Comment{
				{
					ID:          1,
					IDPost:      2,
					IDUser:      1,
					DateComment: timeNow,
					Content:     "content1",
				},
			},
			hasError: nil,
		},
		{
			name:     "GetComByDate() is failed",
			idPost:   2,
			date:     time.Date(2023, 11, 13, 0, 0, 0, 0, time.Local),
			output:   nil,
			hasError: errors.New("no comments on this date in database"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comments, err := rep.GetComByDate(tt.date, tt.idPost)
			log.Printf("comments: %+v, err: %+v", comments, err)
			if !reflect.DeepEqual(comments, tt.output) {
				t.Fatalf("expected %+v, got %+v", tt.output, comments)
			}
			if (err != nil && tt.hasError == nil) || (err == nil && tt.hasError != nil) {
				t.Fatalf("expeced error %+v, got %+v", tt.hasError, err)
			}
		})
	}
}

func TestEditCom(t *testing.T) {
	mockDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("the creation of mock is failed %v", err)
	}
	defer func(mockDB *sql.DB) {
		_ = mockDB.Close()
	}(mockDB)
	customDate := time.Now().In(time.Local)
	rep := NewRepository(mockDB)
	mock.ExpectExec("UPDATE Comment SET IDPost = ?, IDUser = ? , DateComment = ?, Content = ? WHERE ID = ?").WithArgs(2, 1, customDate, "content1", 1).WillReturnResult(sqlmock.NewResult(1, 1))
	result := sqlmock.NewRows([]string{
		"ID", "IDPost", "IDUser", "DateComment", "Content",
	}).AddRow(1, 2, 1, customDate, "content1")
	mock.ExpectQuery("SELECT \\* FROM Comment WHERE ID = ?").WithArgs(1).WillReturnRows(result)
	test := []argEdit{
		{
			name:   "EditComment() is succeed",
			idPost: 2,
			editedCom: Comment{
				ID:          1,
				IDPost:      2,
				IDUser:      1,
				DateComment: customDate,
				Content:     "content1",
			},
			id: 1,
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
			name:   "EditComment() is failed",
			idPost: 1,
			editedCom: Comment{
				ID:          1,
				IDPost:      2,
				IDUser:      1,
				DateComment: customDate,
				Content:     "content1",
			},
			id:       2,
			output:   nil,
			hasError: errors.New("the comment wasn't updated"),
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			comment, err := rep.EditCom(tt.editedCom, tt.id, tt.idPost)
			log.Printf("comments: %+v, err: %+v", comment, err)
			if !reflect.DeepEqual(comment, tt.output) {
				t.Fatalf("expected %+v, got %+v", tt.output, comment)
			}
			if (err != nil && tt.hasError == nil) || (err == nil && tt.hasError != nil) {
				t.Fatalf("expeced error %+v, got %+v", tt.hasError, err)
			}
		})
	}
}

func TestDeleteCom(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("the creation of mock is failed %v", err)
	}
	defer func(mockDB *sql.DB) {
		_ = mockDB.Close()

	}(mockDB)
	rep := NewRepository(mockDB)
	mock.ExpectExec("PRAGMA foreign_keys = ON").WithoutArgs().WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("DELETE FROM Comment WHERE ID = ?").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	test := []argDelete{
		{
			name:     "DeleteCom() is succeed",
			id:       1,
			hasError: nil,
		},
		{
			name:     "DeleteCom() is failed",
			id:       1,
			hasError: errors.New("the comment wasn't deleted"),
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			err := rep.DeleteCom(tt.id)
			log.Printf("err:%v", err)
			if (err != nil && tt.hasError == nil) || (err == nil && tt.hasError != nil) {
				t.Fatalf("expeced error %+v, got %+v", tt.hasError, err)
			}
		})
	}
}
