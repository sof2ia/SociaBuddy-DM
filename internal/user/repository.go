package user

import (
	"database/sql"
)

type Repository interface {
	CreateUser(user User) (*User, error)
	GetUsers() ([]User, error)
	GetUserByID(idUser int) (*User, error)
	GetUserByEmail(emailUser string) (*User, error)
	UpdateUser(user User, idUser int) (*User, error)
	DeleteUser(idUser int) error
}
type repository struct {
	db *sql.DB
}

func (r *repository) CreateUser(user User) (*User, error) {
	res, err := r.db.Exec(`INSERT INTO Users ("Name", "Age", "DocumentNumber", "Email", 
                   "Phone", "ZipCode", "Country", "State", "City", "Neighborhood", "Street", "Number", "Complement")
					VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, user.Name, user.Age, user.DocumentNumber,
		user.Email, user.Phone, user.Address.ZipCode, user.Address.Country, user.Address.State,
		user.Address.City, user.Address.Neighborhood, user.Address.Street, user.Address.Number, user.Address.Complement)
	if err != nil {
		return nil, err
	}
	idUser, err := res.LastInsertId()
	newUser, err := r.GetUserByID(int(idUser))
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (r *repository) GetUsers() ([]User, error) {
	users, err := r.db.Query("SELECT * FROM Users")
	if err != nil {
		return nil, err
	}
	var listUser []User
	for users.Next() {
		var user User
		err := users.Scan(
			&user.ID,
			&user.Name,
			&user.Age,
			&user.DocumentNumber,
			&user.Email,
			&user.Phone,
			&user.Address.ZipCode,
			&user.Address.Country,
			&user.Address.State,
			&user.Address.City,
			&user.Address.Neighborhood,
			&user.Address.Street,
			&user.Address.Number,
			&user.Address.Complement,
		)
		if err != nil {
			return nil, err
		}
		listUser = append(listUser, user)
	}
	return listUser, nil
}

func (r *repository) GetUserByID(idUser int) (*User, error) {
	rows, err := r.db.Query("SELECT * FROM Users WHERE ID = ?", idUser)
	if err != nil {
		return nil, err
	}
	var user User
	if rows.Next() {
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Age,
			&user.DocumentNumber,
			&user.Email,
			&user.Phone,
			&user.Address.ZipCode,
			&user.Address.Country,
			&user.Address.State,
			&user.Address.City,
			&user.Address.Neighborhood,
			&user.Address.Street,
			&user.Address.Number,
			&user.Address.Complement,
		)
		if err != nil {
			return nil, err
		}
	}
	return &user, nil
}
func (r *repository) GetUserByEmail(emailUser string) (*User, error) {
	rows, err := r.db.Query("SELECT * FROM Users WHERE Email = ?", emailUser)
	if err != nil {
		return nil, err
	}
	var user User
	if rows.Next() {
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Age,
			&user.DocumentNumber,
			&user.Email,
			&user.Phone,
			&user.Address.ZipCode,
			&user.Address.Country,
			&user.Address.State,
			&user.Address.City,
			&user.Address.Neighborhood,
			&user.Address.Street,
			&user.Address.Number,
			&user.Address.Complement,
		)
		if err != nil {
			return nil, err
		}
	}
	return &user, nil
}

func (r *repository) UpdateUser(user User, idUser int) (*User, error) {
	_, err := r.db.Exec(`UPDATE Users SET Name = ?, Age = ?, DocumentNumber = ?, Email = ?, 
            Phone = ?, ZipCode = ?, Country = ?, State = ?, City = ?, Neighborhood = ?, Street = ?, Number = ?, Complement = ?
			WHERE ID = ?`, user.Name, user.Age, user.DocumentNumber,
		user.Email, user.Phone, user.Address.ZipCode, user.Address.Country, user.Address.State,
		user.Address.City, user.Address.Neighborhood, user.Address.Street, user.Address.Number, user.Address.Complement, idUser)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) DeleteUser(idUser int) error {
	_, err := r.db.Exec("DELETE FROM Users WHERE ID = ?", idUser)
	if err != nil {
		return err
	}
	return nil
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}
