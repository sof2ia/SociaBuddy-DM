package comment

import (
	"database/sql"
	"log"
	"time"
)

type Repository interface {
	CreateCom(com Comment, idPost int) (*Comment, error)
	GetCom() ([]Comment, error)
	GetComByID(idCom int) (*Comment, error)
	GetComByPostID(idPost int) ([]Comment, error)
	GetComByUserID(idUser int) ([]Comment, error)
	GetComByDate(date time.Time, idPost int) ([]Comment, error)
	EditCom(com Comment, idCom int, idPost int) (*Comment, error)
	DeleteCom(idCom int) error
}

type repository struct {
	db *sql.DB
}

func (r *repository) CreateCom(com Comment, idPost int) (*Comment, error) {
	res, err := r.db.Exec(`INSERT INTO Comment (IDPost, IDUser, DateComment, Content)
VALUES (?, ?, ?, ?)`, idPost, com.IDUser, com.DateComment, com.Content)
	if err != nil {
		return nil, err
	}

	idCom, err := res.LastInsertId()
	newCom, err := r.GetComByID(int(idCom))
	if err != nil {
		return nil, err
	}
	return newCom, nil
}

func (r *repository) GetCom() ([]Comment, error) {
	comments, err := r.db.Query("SELECT * FROM Comment")
	if err != nil {
		return nil, err
	}
	var listCom []Comment
	for comments.Next() {
		var com Comment
		err := comments.Scan(
			&com.ID,
			&com.IDPost,
			&com.IDUser,
			&com.DateComment,
			&com.Content,
		)
		if err != nil {
			return nil, err
		}
		listCom = append(listCom, com)
	}
	return listCom, nil
}

func (r *repository) GetComByID(idCom int) (*Comment, error) {
	rows, err := r.db.Query("SELECT * FROM Comment WHERE ID = ?", idCom)
	if err != nil {
		return nil, err
	}
	var com Comment
	for rows.Next() {
		err := rows.Scan(
			&com.ID,
			&com.IDPost,
			&com.IDUser,
			&com.DateComment,
			&com.Content,
		)
		if err != nil {
			return nil, err
		}
		return &com, nil
	}
	return nil, nil
}

func (r *repository) GetComByPostID(idPost int) ([]Comment, error) {
	rows, err := r.db.Query("SELECT * FROM Comment WHERE IDPost = ?", idPost)
	if err != nil {
		return nil, err
	}
	var listCom []Comment
	for rows.Next() {
		var com Comment
		err := rows.Scan(
			&com.ID,
			&com.IDPost,
			&com.IDUser,
			&com.DateComment,
			&com.Content,
		)
		if err != nil {
			return nil, err
		}
		listCom = append(listCom, com)
	}
	return listCom, nil
}

func (r *repository) GetComByUserID(idUser int) ([]Comment, error) {
	rows, err := r.db.Query("SELECT * FROM Comment WHERE IDUser = ?", idUser)
	if err != nil {
		return nil, err
	}
	var listCom []Comment
	for rows.Next() {
		var com Comment
		err := rows.Scan(
			&com.ID,
			&com.IDPost,
			&com.IDUser,
			&com.DateComment,
			&com.Content,
		)
		if err != nil {
			return nil, err
		}
		listCom = append(listCom, com)
	}
	return listCom, nil
}

func (r *repository) GetComByDate(date time.Time, idPost int) ([]Comment, error) {
	dateFormat := date.Format("2006-01-02")
	log.Println(dateFormat)
	rows, err := r.db.Query(`SELECT * FROM Comment WHERE strftime('%Y-%m-%d', DateComment) = ? AND IDPost = ?`, date.Format("2006-01-02"), idPost)
	if err != nil {
		return nil, err
	}
	var listCom []Comment
	for rows.Next() {
		var com Comment
		err := rows.Scan(
			&com.ID,
			&com.IDPost,
			&com.IDUser,
			&com.DateComment,
			&com.Content,
		)
		if err != nil {
			return nil, err
		}
		listCom = append(listCom, com)
	}
	return listCom, nil
}
func (r *repository) EditCom(com Comment, idCom int, idPost int) (*Comment, error) {
	_, err := r.db.Exec(`UPDATE Comment SET IDPost = ?, IDUser = ? , DateComment = ?, Content = ?
WHERE ID = ?`, idPost, com.IDUser, com.DateComment, com.Content, idCom)
	if err != nil {
		return nil, err
	}

	editedCom, err := r.GetComByID(idCom)
	if err != nil {
		return nil, err
	}

	return editedCom, nil
}
func (r *repository) DeleteCom(idCom int) error {
	_, err := r.db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		return err
	}
	_, err = r.db.Exec("DELETE FROM Comment WHERE ID = ?", idCom)
	if err != nil {
		return err
	}
	return nil
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}
