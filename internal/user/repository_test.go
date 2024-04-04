package user

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"log"
	"reflect"
	"testing"
)

type argGet struct {
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

type argCreate struct {
	name     string
	newUser  User
	output   *User
	hasError error
}

type argUpdate struct {
	name     string
	newUser  User
	id       int
	output   *User
	hasError error
}

type argDelete struct {
	name     string
	id       int
	hasError error
}

type argFollower struct {
	name        string
	idFollower  int
	idFollowing int
	hasError    error
}

type argGetFollow struct {
	name     string
	id       int
	output   []User
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

	test := []argGet{
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
			hasError: errors.New("the email is not found"),
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

func TestCreateUser(t *testing.T) {
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
	mock.ExpectExec("INSERT INTO Users").WithArgs("Name First", 35, "123.345.567-89", "name.first@gmail.com", "+55 11 12345 6789", "12246-260", "Brasil", "SP", "São José dos Campos", "Parque Residencial Aquarius", "Avenida Salmão", "456", "C").WillReturnResult(sqlmock.NewResult(1, 1))
	result := sqlmock.NewRows([]string{
		"ID", "Name", "Age", "DocumentNumber", "Email", "Phone", "ZipCode", "Country", "State", "City", "Neighborhood", "Street", "Number", "Complement",
	}).AddRow(1, "Name First", 35, "123.345.567-89", "name.first@gmail.com", "+55 11 12345 6789", "12246-260", "Brasil", "SP", "São José dos Campos", "Parque Residencial Aquarius", "Avenida Salmão", "456", "C")
	mock.ExpectQuery("SELECT \\* FROM Users WHERE ID = ?").WithArgs(1).WillReturnRows(result)
	test := []argCreate{
		{
			name: "CreateUser() is succeed",
			newUser: User{
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
					Complement:   "C",
				},
			},
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
					Complement:   "C",
				},
			},
			hasError: nil,
		},
		{
			name: "CreateUser() is failed",
			newUser: User{
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
					Complement:   "C",
				},
			},
			output:   nil,
			hasError: errors.New("new user wasn't found"),
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			users, err := rep.CreateUser(tt.newUser)
			log.Printf("user: %v, err: %v", users, err)
			if !reflect.DeepEqual(users, tt.output) {
				t.Fatalf("expected %+v, got %+v", tt.output, users)
			}
			if (err != nil && tt.hasError == nil) || (err == nil && tt.hasError != nil) {
				t.Fatalf("expeced error %+v, got %+v", tt.hasError, err)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	mockDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("the creation of mock is failed: %+v", mockDB)
	}
	defer func(mockDB *sql.DB) {
		err := mockDB.Close()
		if err != nil {
		}
	}(mockDB)
	rep := NewRepository(mockDB)
	mock.ExpectExec("UPDATE Users SET Name = ?, Age = ?, DocumentNumber = ?, Email = ?, Phone = ?, ZipCode = ?, Country = ?, State = ?, City = ?, Neighborhood = ?, Street = ?, Number = ?, Complement = ? WHERE ID = ?").WithArgs("Name First", 35, "123.345.567-89", "name.first@gmail.com", "+55 11 12345 6789", "12246-260", "Brasil", "SP", "São José dos Campos", "Parque Residencial Aquarius", "Avenida Salmão", "456", "C", 1).WillReturnResult(sqlmock.NewResult(1, 1))
	result := sqlmock.NewRows([]string{
		"ID", "Name", "Age", "DocumentNumber", "Email", "Phone", "ZipCode", "Country", "State", "City", "Neighborhood", "Street", "Number", "Complement",
	}).AddRow(1, "Name First", 35, "123.345.567-89", "name.first@gmail.com", "+55 11 12345 6789", "12246-260", "Brasil", "SP", "São José dos Campos", "Parque Residencial Aquarius", "Avenida Salmão", "456", "C")
	mock.ExpectQuery("SELECT \\* FROM Users WHERE ID = ?").WithArgs(1).WillReturnRows(result)

	test := []argUpdate{
		{
			name: "UpdateUser() is succeed",
			newUser: User{
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
					Complement:   "C",
				},
			},
			id: 1,
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
					Complement:   "C",
				},
			},
			hasError: nil,
		},
		{
			name: "UpdateUser() is failed",
			newUser: User{
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
					Complement:   "C",
				},
			},
			id:       1,
			output:   nil,
			hasError: errors.New("the user wasn't updated"),
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			users, err := rep.UpdateUser(tt.newUser, tt.id)
			log.Printf("user: %v, err: %v", users, err)
			if !reflect.DeepEqual(users, tt.output) {
				t.Fatalf("expected %+v, got %+v", tt.output, users)
			}
			if (err != nil && tt.hasError == nil) || (err == nil && tt.hasError != nil) {
				t.Fatalf("expeced error %+v, got %+v", tt.hasError, err)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
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
	mock.ExpectExec("PRAGMA foreign_keys = ON").WithoutArgs().WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("DELETE FROM Users WHERE ID = ?").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
	test := []argDelete{
		{
			name:     "DeleteUser() is succeed",
			id:       1,
			hasError: nil,
		},
		{
			name:     "DeleteUser() is failed",
			id:       1,
			hasError: errors.New("the user wasn't deleted"),
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			err := rep.DeleteUser(tt.id)
			log.Printf("err: %v", err)
			if (err != nil && tt.hasError == nil) || (err == nil && tt.hasError != nil) {
				t.Fatalf("expeced error %+v, got %+v", tt.hasError, err)
			}
		})
	}
}

func TestFollowUser(t *testing.T) {
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
	mock.ExpectExec("INSERT INTO Connection").WithArgs(1, 2).WillReturnResult(sqlmock.NewResult(1, 1))
	result := sqlmock.NewRows([]string{"ID", "idFollower", "idFollowing"}).AddRow(1, 1, 2)
	mock.ExpectQuery("SELECT \\* FROM Connection WHERE IdFollower = \\?").WithArgs(1).WillReturnRows(result)
	test := []argFollower{
		{
			name:        "FollowUser() is succeed",
			idFollower:  1,
			idFollowing: 2,
			hasError:    nil,
		},
		{
			name:        "FollowUser() is failed",
			idFollower:  1,
			idFollowing: 3,
			hasError:    errors.New("the connection is incorrect"),
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			err := rep.FollowUser(tt.idFollower, tt.idFollowing)
			log.Printf("err: %v", err)
			if (err != nil && tt.hasError == nil) || (err == nil && tt.hasError != nil) {
				t.Fatalf("expeced error %+v, got %+v", tt.hasError, err)
			}
		})
	}
}

func TestDeleteConnection(t *testing.T) {
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
	mock.ExpectExec("DELETE FROM Connection WHERE IdFollower = \\? AND IdFollowing = \\?").WithArgs(1, 2).WillReturnResult(sqlmock.NewResult(1, 1))
	test := []argFollower{
		{
			name:        "DeleteConnection() is succeed",
			idFollower:  1,
			idFollowing: 2,
			hasError:    nil,
		},
		{
			name:        "DeleteConnection() is failed",
			idFollower:  1,
			idFollowing: 2,
			hasError:    errors.New("the connection wasn't deleted"),
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			err := rep.DeleteConnection(tt.idFollower, tt.idFollowing)
			log.Printf("err: %v", err)
			if (err != nil && tt.hasError == nil) || (err == nil && tt.hasError != nil) {
				t.Fatalf("expeced error %+v, got %+v", tt.hasError, err)
			}
		})
	}
}

func TestGetFollowingByUserID(t *testing.T) {
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
	}).AddRow(2, "Name First", 35, "123.345.567-89", "name.first@gmail.com", "+55 11 12345 6789", "12246-260", "Brasil", "SP", "São José dos Campos", "Parque Residencial Aquarius", "Avenida Salmão", "456", "C")
	mock.ExpectQuery("SELECT Users.\\* FROM Users INNER JOIN Connection ON Users.ID = Connection.idFollowing WHERE Connection.idFollower = \\? ").WithArgs(3).WillReturnRows(result)
	test := []argGetFollow{
		{
			name: "GetFollowingByUserID() is succeed",
			id:   3,
			output: []User{
				{
					ID:             2,
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
						Complement:   "C",
					},
				},
			},
			hasError: nil,
		},
		{
			name:     "GetFollowingByUserID() is failed - has no following",
			id:       7,
			output:   nil,
			hasError: errors.New("the id has no following "),
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			users, err := rep.GetFollowingByUserID(tt.id)
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

func TestGetUserFollowers(t *testing.T) {
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

	result1 := sqlmock.NewRows([]string{
		"ID", "Name", "Age", "DocumentNumber", "Email", "Phone", "ZipCode", "Country", "State", "City", "Neighborhood", "Street", "Number", "Complement",
	}).AddRow(2, "Name First", 35, "123.345.567-89", "name.first@gmail.com", "+55 11 12345 6789", "12246-260", "Brasil", "SP", "São José dos Campos", "Parque Residencial Aquarius", "Avenida Salmão", "456", "C")
	mock.ExpectQuery("SELECT Users.\\* FROM Users INNER JOIN Connection ON Users.ID = Connection.idFollower WHERE Connection.idFollowing = \\? ").WithArgs(1).WillReturnRows(result1)

	test := []argGetFollow{
		{
			name: "GetUserFollowers() is succeed",
			id:   1,
			output: []User{
				{
					ID:             2,
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
						Complement:   "C",
					},
				},
			},
			hasError: nil,
		},
		{
			name:     "GetUserFollowers() is failed - has no followers",
			id:       5,
			output:   nil,
			hasError: errors.New("the id has no followers"),
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			users, err := rep.GetUserFollowers(tt.id)
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
