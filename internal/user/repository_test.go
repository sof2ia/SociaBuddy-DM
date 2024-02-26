package user

import (
	"github.com/DATA-DOG/go-sqlmock"
	"reflect"
	"testing"
)

type arg struct {
	name     string
	input    []User
	hasError error
}

func TestGetUser(t *testing.T) {

	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("the creation of mock is failed %v", err)
	}
	defer mockDB.Close()
	rep := NewRepository(mockDB)
	result := sqlmock.NewRows([]string{
		"ID", "Name", "Age", "DocumentNumber", "Email", "Phone", "ZipCode", "Country", "State", "City", "Neighborhood", "Street", "Number", "Complement",
	}).AddRow(1, "Name First", 35, "123.345.567-89", "nome.first@gmail.com", "+55 11 12345 6789", "12246-260", "Brasil", "SP", "São José dos Campos", "Parque Residencial Aquarius", "Avenida Salmão", "456", "C")
	mock.ExpectQuery("SELECT * FROM Users").WillReturnRows(result)

	test := []arg{
		{name: "GetUsers() from database succeed",
			input: []User{
				{ID: 1,
					Name:           "Name First",
					Age:            35,
					DocumentNumber: "123.345.567-89",
					Email:          "nome.first@gmail.com",
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
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			users, err := rep.GetUsers()
			if !reflect.DeepEqual(users, tt.input) {
				t.Fatalf("expected %+v, got %+v", tt.input, users)
			}
			if (err != nil && tt.hasError == nil) || (err == nil && tt.hasError != nil) {
				t.Fatalf("expeced error %+v, got %+v", tt.hasError, err)
			}
		})
	}
}
