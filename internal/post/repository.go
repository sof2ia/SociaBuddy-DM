package post

import (
	"database/sql"
	"log"
	"time"
)

type Repository interface {
	CreatePost(post Post) (*Post, error)
	GetPosts() ([]Post, error)
	GetPostByID(idPost int) (*Post, error)
	GetPostByUserID(idUser int) ([]Post, error)
	GetPostByDate(date time.Time) ([]Post, error)
	GetPostByTitle(title string) ([]Post, error)
	EditPost(post Post, idPost int) (*Post, error)
	DeletePost(idPost int) error
}

type repository struct {
	db *sql.DB
}

func (r *repository) CreatePost(post Post) (*Post, error) {
	res, err := r.db.Exec(`INSERT INTO Posts (IDUser, DatePost, Title, Content)
	VALUES (?,?,?,?)`, post.IDUser, post.Date, post.Title, post.Content)

	if err != nil {
		return nil, err
	}
	idPost, err := res.LastInsertId()
	newPost, err := r.GetPostByID(int(idPost))
	if err != nil {
		return nil, err
	}
	return newPost, nil
}

func (r *repository) GetPosts() ([]Post, error) {
	posts, err := r.db.Query("SELECT * FROM Posts")
	if err != nil {
		return nil, err
	}
	var listPosts []Post
	for posts.Next() {
		var post Post
		err := posts.Scan(
			&post.ID,
			&post.IDUser,
			&post.Date,
			&post.Title,
			&post.Content,
		)
		if err != nil {
			return nil, err
		}
		listPosts = append(listPosts, post)
	}
	return listPosts, nil
}

func (r *repository) GetPostByID(idPost int) (*Post, error) {
	rows, err := r.db.Query("SELECT * FROM Posts WHERE ID = ?", idPost)
	if err != nil {
		return nil, err
	}
	var post Post
	for rows.Next() {
		err := rows.Scan(
			&post.ID,
			&post.IDUser,
			&post.Date,
			&post.Title,
			&post.Content,
		)
		if err != nil {
			return nil, err
		}
		return &post, nil
	}
	return nil, nil
}

func (r *repository) GetPostByUserID(idUser int) ([]Post, error) {
	rows, err := r.db.Query("SELECT * FROM Posts WHERE IDUser = ?", idUser)
	if err != nil {
		return nil, err
	}
	var listPosts []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(
			&post.ID,
			&post.IDUser,
			&post.Date,
			&post.Title,
			&post.Content,
		)
		if err != nil {
			return nil, err
		}
		listPosts = append(listPosts, post)
	}
	return listPosts, nil
}

func (r *repository) GetPostByDate(date time.Time) ([]Post, error) {
	dateFormat := date.Format("2006-01-02")
	log.Println(dateFormat)
	rows, err := r.db.Query(`SELECT * FROM Posts WHERE strftime('%Y-%m-%d', DatePost) = ?`, date.Format("2006-01-02"))
	if err != nil {
		return nil, err
	}
	var listPosts []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(
			&post.ID,
			&post.IDUser,
			&post.Date,
			&post.Title,
			&post.Content,
		)
		if err != nil {
			return nil, err
		}
		listPosts = append(listPosts, post)
	}
	return listPosts, nil
}

func (r *repository) GetPostByTitle(title string) ([]Post, error) {
	rows, err := r.db.Query("SELECT * FROM Posts WHERE Title = ?", title)
	if err != nil {
		return nil, err
	}
	var listPosts []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(
			&post.ID,
			&post.IDUser,
			&post.Date,
			&post.Title,
			&post.Content,
		)
		if err != nil {
			return nil, err
		}
		listPosts = append(listPosts, post)
	}
	return listPosts, nil
}

func (r *repository) EditPost(post Post, idPost int) (*Post, error) {
	_, err := r.db.Exec(`UPDATE Posts SET IDUser = ?, DatePost = ?, Title = ?, Content = ?
			WHERE ID = ?`, post.IDUser, post.Date, post.Title, post.Content, idPost)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *repository) DeletePost(idPost int) error {
	_, err := r.db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		return err
	}
	_, err = r.db.Exec("DELETE FROM Posts WHERE ID = ?", idPost)
	if err != nil {
		return err
	}
	return nil
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}
