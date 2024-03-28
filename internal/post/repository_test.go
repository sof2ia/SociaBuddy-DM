package post

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
	output   []Post
	hasError error
}

type argCreate struct {
	name     string
	newPost  Post
	output   *Post
	hasError error
}

type argID struct {
	name     string
	id       int
	output   *Post
	hasError error
}

type argIDUser struct {
	name     string
	idUser   int
	output   []Post
	hasError error
}

type argDate struct {
	name     string
	date     time.Time
	output   []Post
	hasError error
}

type argTitle struct {
	name     string
	title    string
	output   []Post
	hasError error
}

type argEdit struct {
	name       string
	editedPost Post
	id         int
	output     *Post
	hasError   error
}

type argDelete struct {
	name     string
	id       int
	hasError error
}

func TestGetPosts(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("the creation of mock is failed %v", err)
	}
	defer func(mockDB *sql.DB) {
		_ = mockDB.Close()

	}(mockDB)
	rep := NewRepository(mockDB)

	//format := "Mon Jan _2 15:04:05 2006"
	timeNow := time.Date(2023, 11, 13, 0, 0, 0, 0, time.Local)
	//customDate, err := time.Parse(format, timeNow)

	result := sqlmock.NewRows([]string{
		"ID", "IDUser", "DatePost", "Title", "Content",
	}).AddRow(1, 2, timeNow, "title1", "content1")
	mock.ExpectQuery("SELECT \\* FROM Posts").WillReturnRows(result)

	test := []argGet{
		{name: "GetPosts() is succeed",
			output: []Post{
				{ID: 1,
					IDUser:  2,
					Date:    timeNow,
					Title:   "title1",
					Content: "content1",
				},
			},

			hasError: nil,
		},
		{
			name:     "GetPost() when there is no result",
			output:   nil,
			hasError: errors.New("no posts in database"),
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			posts, err := rep.GetPosts()
			log.Printf("users: %+v, err: %+v", posts, err)
			if !reflect.DeepEqual(posts, tt.output) {
				t.Fatalf("expected %+v, got %+v", tt.output, posts)
			}
			if (err != nil && tt.hasError == nil) || (err == nil && tt.hasError != nil) {
				t.Fatalf("expeced error %+v, got %+v", tt.hasError, err)
			}
		})
	}
}

func TestCreatePost(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("the creation of mock is failed %v", err)
	}
	defer func(mockDB *sql.DB) {
		_ = mockDB.Close()

	}(mockDB)

	customDate := time.Now().In(time.Local)
	log.Printf("test: %v", customDate)
	rep := NewRepository(mockDB)
	mock.ExpectExec("INSERT INTO Posts").WithArgs(2, customDate, "title1", "content1").WillReturnResult(sqlmock.NewResult(1, 1))
	result := sqlmock.NewRows([]string{
		"ID", "IDUser", "DatePost", "Title", "Content",
	}).AddRow(1, 2, customDate, "title1", "content1")
	mock.ExpectQuery("SELECT \\* FROM Posts WHERE ID = ?").WithArgs(1).WillReturnRows(result)

	test := []argCreate{
		{name: "CreatePost() is succeed",
			newPost: Post{
				ID:      1,
				IDUser:  2,
				Date:    customDate,
				Title:   "title1",
				Content: "content1",
			},
			output: &Post{
				ID:      1,
				IDUser:  2,
				Date:    customDate,
				Title:   "title1",
				Content: "content1",
			},

			hasError: nil,
		},
		{
			name: "CreatePost() when there is no result",
			newPost: Post{
				ID:      1,
				IDUser:  2,
				Date:    customDate,
				Title:   "title1",
				Content: "content1",
			},
			output:   nil,
			hasError: errors.New("publication has not succeed"),
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			posts, err := rep.CreatePost(tt.newPost)
			log.Printf("posts: %+v, err: %+v", posts, err)
			if !reflect.DeepEqual(posts, tt.output) {
				t.Fatalf("expected %+v, got %+v", tt.output, posts)
			}
			if (err != nil && tt.hasError == nil) || (err == nil && tt.hasError != nil) {
				t.Fatalf("expeced error %+v, got %+v", tt.hasError, err)
			}
		})
	}
}
func TestGetPostByID(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("the creation of mock is failed %v", err)
	}
	defer func(mockDB *sql.DB) {
		_ = mockDB.Close()

	}(mockDB)
	//format := "Mon Jan _2 15:04:05 2006"
	timeNow := time.Date(2023, 11, 13, 0, 0, 0, 0, time.Local)

	rep := NewRepository(mockDB)
	result := sqlmock.NewRows([]string{
		"ID", "IDUser", "DatePost", "Title", "Content",
	}).AddRow(1, 2, timeNow, "title1", "content1")
	mock.ExpectQuery("SELECT \\* FROM Posts WHERE ID = ?").WithArgs(1).WillReturnRows(result)

	test := []argID{
		{name: "GetPostsByID() is succeed",
			id: 1,
			output: &Post{
				ID:      1,
				IDUser:  2,
				Date:    timeNow,
				Title:   "title1",
				Content: "content1",
			},

			hasError: nil,
		},
		{
			name:     "GetPostByID() has not succeed",
			id:       2,
			output:   nil,
			hasError: errors.New("no posts with this id was found"),
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			posts, err := rep.GetPostByID(tt.id)
			log.Printf("posts: %+v, err: %+v", posts, err)
			if !reflect.DeepEqual(posts, tt.output) {
				t.Fatalf("expected %+v, got %+v", tt.output, posts)
			}
			if (err != nil && tt.hasError == nil) || (err == nil && tt.hasError != nil) {
				t.Fatalf("expeced error %+v, got %+v", tt.hasError, err)
			}
		})
	}
}

func TestGetPostByUserID(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("the creation of mock is failed %v", err)
	}
	defer func(mockDB *sql.DB) {
		_ = mockDB.Close()

	}(mockDB)
	//format := "Mon Jan _2 15:04:05 2006"
	timeNow := time.Date(2023, 11, 13, 0, 0, 0, 0, time.Local)

	rep := NewRepository(mockDB)
	result := sqlmock.NewRows([]string{
		"ID", "IDUser", "DatePost", "Title", "Content",
	}).AddRow(1, 2, timeNow, "title1", "content1")
	mock.ExpectQuery("SELECT \\* FROM Posts WHERE IDUser = ?").WithArgs(2).WillReturnRows(result)

	test := []argIDUser{
		{name: "GetPostsByUserID() is succeed",
			idUser: 2,
			output: []Post{
				{
					ID:      1,
					IDUser:  2,
					Date:    timeNow,
					Title:   "title1",
					Content: "content1",
				},
			},

			hasError: nil,
		},
		{
			name:     "GetPostByUserID() has not succeed",
			idUser:   1,
			output:   nil,
			hasError: errors.New("no posts with this id was found"),
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			posts, err := rep.GetPostByUserID(tt.idUser)
			log.Printf("posts: %+v, err: %+v", posts, err)
			if !reflect.DeepEqual(posts, tt.output) {
				t.Fatalf("expected %+v, got %+v", tt.output, posts)
			}
			if (err != nil && tt.hasError == nil) || (err == nil && tt.hasError != nil) {
				t.Fatalf("expeced error %+v, got %+v", tt.hasError, err)
			}
		})
	}
}

func TestGetPostByDate(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("the creation of mock is failed %v", err)
	}
	defer func(mockDB *sql.DB) {
		_ = mockDB.Close()

	}(mockDB)
	//format := "Mon Jan _2 15:04:05 2006"
	timeNow := time.Date(2023, 11, 13, 0, 0, 0, 0, time.Local)

	rep := NewRepository(mockDB)
	result := sqlmock.NewRows([]string{
		"ID", "IDUser", "Date", "Title", "Content",
	}).AddRow(1, 2, timeNow, "title1", "content1")
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM Posts WHERE DatePost LIKE '?%'`)).WithArgs(time.Date(2023, 11, 13, 0, 0, 0, 0, time.Local).Format("2006-01-02")).WillReturnRows(result)

	test := []argDate{
		{name: "GetPostsByDate() is succeed",
			date: time.Date(2023, 11, 13, 0, 0, 0, 0, time.Local),
			output: []Post{
				{ID: 1,
					IDUser:  2,
					Date:    timeNow,
					Title:   "title1",
					Content: "content1",
				},
			},

			hasError: nil,
		},
		{
			name:     "GetPostByDate() when there is no result",
			output:   nil,
			hasError: errors.New("no posts on this date in database"),
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			posts, err := rep.GetPostByDate(tt.date)
			log.Printf("posts: %+v, err: %+v", posts, err)
			if !reflect.DeepEqual(posts, tt.output) {
				t.Fatalf("expected %+v, got %+v", tt.output, posts)
			}
			if (err != nil && tt.hasError == nil) || (err == nil && tt.hasError != nil) {
				t.Fatalf("expeced error %+v, got %+v", tt.hasError, err)
			}
		})
	}
}
func TestGetPostByTitle(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("the creation of mock is failed %v", err)
	}
	defer func(mockDB *sql.DB) {
		_ = mockDB.Close()

	}(mockDB)
	//format := "Mon Jan _2 15:04:05 2006"
	timeNow := time.Date(2023, 11, 13, 0, 0, 0, 0, time.Local)

	rep := NewRepository(mockDB)
	result := sqlmock.NewRows([]string{
		"ID", "IDUser", "DatePost", "Title", "Content",
	}).AddRow(1, 2, timeNow, "title1", "content1")
	mock.ExpectQuery("SELECT \\* FROM Posts WHERE Title =?").WithArgs("title1").WillReturnRows(result)

	test := []argTitle{
		{name: "GetPosts() is succeed",
			title: "title1",
			output: []Post{
				{ID: 1,
					IDUser:  2,
					Date:    timeNow,
					Title:   "title1",
					Content: "content1",
				},
			},

			hasError: nil,
		},
		{
			name:     "GetPost() when there is no result",
			title:    "title2",
			output:   nil,
			hasError: errors.New("no posts in database"),
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			posts, err := rep.GetPostByTitle(tt.title)
			log.Printf("posts: %+v, err: %+v", posts, err)
			if !reflect.DeepEqual(posts, tt.output) {
				t.Fatalf("expected %+v, got %+v", tt.output, posts)
			}
			if (err != nil && tt.hasError == nil) || (err == nil && tt.hasError != nil) {
				t.Fatalf("expeced error %+v, got %+v", tt.hasError, err)
			}
		})
	}
}

func TestEditPost(t *testing.T) {
	mockDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("the creation of mock is failed %v", err)
	}
	defer func(mockDB *sql.DB) {
		_ = mockDB.Close()

	}(mockDB)
	//format := "Mon Jan _2 15:04:05 2006"
	customDate := time.Now().In(time.Local)
	log.Printf("test: %v", customDate)
	rep := NewRepository(mockDB)
	mock.ExpectExec("UPDATE Posts SET IDUser = ?, DatePost = ?, Title = ?, Content = ? WHERE ID = ?").WithArgs(2, customDate, "title1", "content1", 1).WillReturnResult(sqlmock.NewResult(1, 1))
	test := []argEdit{
		{name: "GetPosts() is succeed",
			editedPost: Post{
				ID:      1,
				IDUser:  2,
				Date:    customDate,
				Title:   "title1",
				Content: "content1",
			},
			id: 1,
			output: &Post{
				ID:      1,
				IDUser:  2,
				Date:    customDate,
				Title:   "title1",
				Content: "content1",
			},
			hasError: nil,
		},
		{
			name: "EditPost() is failed",
			editedPost: Post{
				ID:      1,
				IDUser:  2,
				Date:    customDate,
				Title:   "title1",
				Content: "content1",
			},
			id:       1,
			output:   nil,
			hasError: errors.New("the post wasn't updated"),
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			post, err := rep.EditPost(tt.editedPost, tt.id)
			log.Printf("post: %+v, err: %+v", post, err)
			if !reflect.DeepEqual(post, tt.output) {
				t.Fatalf("expected %+v, got %+v", tt.output, post)
			}
			if (err != nil && tt.hasError == nil) || (err == nil && tt.hasError != nil) {
				t.Fatalf("expeced error %+v, got %+v", tt.hasError, err)
			}
		})
	}
}

func TestDeletePost(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("the creation of mock is failed %v", err)
	}
	defer func(mockDB *sql.DB) {
		_ = mockDB.Close()

	}(mockDB)
	rep := NewRepository(mockDB)
	mock.ExpectExec("DELETE FROM Posts WHERE ID = ?").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	test := []argDelete{
		{name: "DeletePost() is succeed",
			id: 1,

			hasError: nil,
		},
		{
			name:     "DeletePost() is failed",
			id:       1,
			hasError: errors.New("the post wasn't deleted"),
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			err := rep.DeletePost(tt.id)
			log.Printf("err: %v", err)
			if (err != nil && tt.hasError == nil) || (err == nil && tt.hasError != nil) {
				t.Fatalf("expeced error %+v, got %+v", tt.hasError, err)
			}
		})
	}
}
