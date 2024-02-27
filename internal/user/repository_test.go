package user

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"log"
	"reflect"
	"testing"
)

type arg struct {
	name     string
	output   []User
	hasError error
}

type argID struct {
	name     string
	id       int
	output   *User
	hasError error
}

type argEmail struct {
	name     string
	email    string
	output   *User
	hasError error
}

func TestGetUser(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("the creation of mock is failed %v", err)
	}
	defer func(mockDB *sql.DB) {
		err := mockDB.Close()
		if err != nil {

		}
	}(mockDB)
	rep := NewRepository(mockDB)
	result := sqlmock.NewRows([]string{
		"ID", "Name", "Age", "DocumentNumber", "Email", "Phone", "ZipCode", "Country", "State", "City", "Neighborhood", "Street", "Number", "Complement",
	}).AddRow(1, "Name First", 35, "123.345.567-89", "name.first@gmail.com", "+55 11 12345 6789", "12246-260", "Brasil", "SP", "São José dos Campos", "Parque Residencial Aquarius", "Avenida Salmão", "456", "C")
	mock.ExpectQuery("SELECT \\* FROM Users").WillReturnRows(result)

	test := []arg{
		{name: "GetUsers() from database succeed",
			output: []User{
				{ID: 1,
					Name:           "Name First",
					Age:            35,
					DocumentNumber: "123.345.567-89",
					Email:          "name.first@gmail.com",
					Phone:          "+55 11 12345 6789",
					Address: Address{
						ZipCode:      "12246-260",
						Country:      "Brasil",
						State:        "SP",
						City:         "São José dos Campos",
						Neighborhood: "Parque Residencial Aquarius",
						Street:       "Avenida Salmão",
						Number:       "456",
						Complement:   "C"},
				},
			},

			hasError: nil,
		},
		{
			name:     "GetUsers() when there is no result",
			output:   nil,
			hasError: errors.New("no users in database"),
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			users, err := rep.GetUsers()
			log.Printf("users: %+v, err: %+v", users, err)
			if !reflect.DeepEqual(users, tt.output) {
				t.Fatalf("expected %+v, got %+v", tt.output, users)
			}
			if (err != nil && tt.hasError == nil) || (err == nil && tt.hasError != nil) {
				t.Fatalf("expeced error %+v, got %+v", tt.hasError, err)
			}
		})
	}
}

func TestGetUserByID(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("the creation of mock is failed: %+v", mockDB)
	}
	defer func(mockDB *sql.DB) {
		err := mockDB.Close()
		if err != nil {

		}
	}(mockDB)
	rep := NewRepository(mockDB)
	result := sqlmock.NewRows([]string{
		"ID", "Name", "Age", "DocumentNumber", "Email", "Phone", "ZipCode", "Country", "State", "City", "Neighborhood", "Street", "Number", "Complement",
	}).AddRow(1, "Name First", 35, "123.345.567-89", "name.first@gmail.com", "+55 11 12345 6789", "12246-260", "Brasil", "SP", "São José dos Campos", "Parque Residencial Aquarius", "Avenida Salmão", "456", "C")
	mock.ExpectQuery("SELECT \\* FROM Users WHERE ID = ?").WithArgs(1).WillReturnRows(result)
	test := []argID{
		{
			name: "GetUserByID() is succeed",
			id:   1,
			output: &User{
				ID:             1,
				Name:           "Name First",
				Age:            35,
				DocumentNumber: "123.345.567-89",
				Email:          "name.first@gmail.com",
				Phone:          "+55 11 12345 6789",
				Address: Address{
					ZipCode:      "12246-260",
					Country:      "Brasil",
					State:        "SP",
					City:         "São José dos Campos",
					Neighborhood: "Parque Residencial Aquarius",
					Street:       "Avenida Salmão",
					Number:       "456",
					Complement:   "C"},
			},
			hasError: nil,
		},
		{
			name:     "GetUserByID() when ID wasn't found",
			id:       2,
			output:   nil,
			hasError: errors.New("the ID is not found"),
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			users, err := rep.GetUserByID(tt.id)
			log.Printf("users: %+v, err: %+v", users, err)
			if !reflect.DeepEqual(users, tt.output) {
				t.Fatalf("expected %+v, got %+v", tt.output, users)
			}
			if (err != nil && tt.hasError == nil) || (err == nil && tt.hasError != nil) {
				t.Fatalf("expeced error %+v, got %+v", tt.hasError, err)
			}
		})
	}
}

func TestGetUserByEmail(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("the creation of mock is failed: %+v", mockDB)
	}
	defer func(mockDB *sql.DB) {
		err := mockDB.Close()
		if err != nil {

		}
	}(mockDB)
	rep := NewRepository(mockDB)
	result := sqlmock.NewRows([]string{
		"ID", "Name", "Age", "DocumentNumber", "Email", "Phone", "ZipCode", "Country", "State", "City", "Neighborhood", "Street", "Number", "Complement",
	}).AddRow(1, "Name First", 35, "123.345.567-89", "name.first@gmail.com", "+55 11 12345 6789", "12246-260", "Brasil", "SP", "São José dos Campos", "Parque Residencial Aquarius", "Avenida Salmão", "456", "C")
	mock.ExpectQuery("SELECT \\* FROM Users WHERE Email = ?").WithArgs("name.first@gmail.com").WillReturnRows(result)
	test := []argEmail{
		{
			name:  "GetUserByEmail() is succeed",
			email: "name.first@gmail.com",
			output: &User{
				ID:             1,
				Name:           "Name First",
				Age:            35,
				DocumentNumber: "123.345.567-89",
				Email:          "name.first@gmail.com",
				Phone:          "+55 11 12345 6789",
				Address: Address{
					ZipCode:      "12246-260",
					Country:      "Brasil",
					State:        "SP",
					City:         "São José dos Campos",
					Neighborhood: "Parque Residencial Aquarius",
					Street:       "Avenida Salmão",
					Number:       "456",
					Complement:   "C"},
			},
			hasError: nil,
		},
		{
			name:     "GetUserByID() when ID wasn't found",
			email:    "name.second@gmail.com",
			output:   nil,
			hasError: errors.New("the ID is not found"),
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			users, err := rep.GetUserByEmail(tt.email)
			log.Printf("users: %+v, err: %+v", users, err)
			if !reflect.DeepEqual(users, tt.output) {
				t.Fatalf("expected %+v, got %+v", tt.output, users)
			}
			if (err != nil && tt.hasError == nil) || (err == nil && tt.hasError != nil) {
				t.Fatalf("expeced error %+v, got %+v", tt.hasError, err)
			}
		})
	}
}
