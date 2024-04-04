package main

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"socialBuddy/internal/comment"
	"socialBuddy/internal/post"
	"socialBuddy/internal/user"
)

func main() {
	file := "../internal/database/socialbuddy.db"

	db, err := startSqlite("Users", "Posts", "Comment", file)
	if err != nil {
		log.Fatal(err)
		return
	}

	repUser := user.NewRepository(db)
	cli := http.DefaultClient
	fac := user.NewFacade("https://viacep.com.br", cli)
	servUser := user.NewService(repUser, fac)
	serUser := user.NewServer(servUser)

	repPost := post.NewRepository(db)
	servPost := post.NewService(repPost, servUser)
	serPost := post.NewServer(servPost)

	repCom := comment.NewRepository(db)
	servCom := comment.NewService(repCom, servPost, servUser)
	serCom := comment.NewServer(servCom)

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/v1/user", serUser.GetUsers)
	router.Get("/v1/user/{id}", serUser.GetUserByID)
	router.Get("/v1/user/email/{email}", serUser.GetUserByEmail)
	router.Post("/v1/user", serUser.CreateUser)
	router.Put("/v1/user/{id}", serUser.UpdateUser)
	router.Delete("/v1/user/{id}", serUser.DeleteUser)

	router.Put("/v1/user/{id}/following/{following_id}", serUser.FollowUser)
	router.Delete("/v1/user/{id}/following/{following_id}", serUser.DeleteConnection)
	router.Get("/v1/user/{id}/following", serUser.GetFollow)

	router.Get("/v1/post", serPost.GetPosts)
	router.Get("/v1/post/{id}", serPost.GetPostByID)
	router.Get("/v1/post/id/{id_user}", serPost.GetPostByUserID)
	router.Get("/v1/post/title/{title}", serPost.GetPostByTitle)
	router.Get("/v1/post/date/{date}", serPost.GetPostByDate)
	router.Post("/v1/post", serPost.CreatePost)
	router.Put("/v1/post/{id}", serPost.EditPost)
	router.Delete("/v1/post/{id}", serPost.DeletePost)

	router.Get("/v1/comment", serCom.GetCom)
	router.Get("/v1/post/{id_post}/comment", serCom.GetComByPostID)
	router.Get("/v1/user/{id_user}/comment", serCom.GetComByUserID)
	router.Get("/v1/post/{id_post}/comment/{id}", serCom.GetComByID)
	router.Get("/v1/post/{id_post}/date/{date}/comment", serCom.GetComByDate)
	router.Post("/v1/post/{id_post}/comment", serCom.CreateCom)
	router.Put("/v1/post/{id_post}/comment/{id}", serCom.EditCom)
	router.Delete("/v1/post/{id_post}/comment/{id}", serCom.DeleteCom)

	log.Println("server's running on the port: 8081")
	err = http.ListenAndServe(":8081", router)

	if err != nil {
		log.Fatal(err)
		return
	}
}
func startSqlite(tableName1 string, tableName2 string, tableName3 string, file string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	_, err = db.Exec(fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
    												ID INTEGER PRIMARY KEY AUTOINCREMENT,
    												Name 		 TEXT,
    												Age 		 TEXT,
    												DocumentNumber TEXT,
    												Email 		 TEXT,
    												Phone 		 TEXT,
    												ZipCode 	 TEXT,
    												Country 	 TEXT, 
    												State        TEXT,
    												City 		 TEXT, 
    												Neighborhood TEXT
													Street       TEXT
													Number       TEXT
                              						Complement   TEXT
    												)    `, tableName1))
	if err != nil {
		log.Println(err)
		err := db.Close()
		if err != nil {
			return nil, err
		}
		return nil, err
	}
	_, err = db.Exec(fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
    												ID INTEGER PRIMARY KEY AUTOINCREMENT,
    												IDUser INTEGER,
    												DatePost DATE,
    												Title TEXT,
    												Content TEXT,
    												FOREIGN KEY (IDUser) REFERENCES Users(ID)
    												)    `, tableName2))
	if err != nil {
		log.Println(err)
		err := db.Close()
		if err != nil {
			return nil, err
		}

		return nil, err
	}
	_, err = db.Exec(fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
  													ID INTEGER PRIMARY KEY AUTOINCREMENT,
   													IDPost INTEGER,
    												IDUser INTEGER,
   													DateComment DATE,
    												Content TEXT,
    												FOREIGN KEY (IDPost) REFERENCES Posts(ID)
   													FOREIGN KEY (IDUser) REFERENCES Users(ID)
													)`, tableName3))
	if err != nil {
		log.Println(err)
		err := db.Close()
		if err != nil {
			return nil, err
		}

		return nil, err
	}
	return db, nil
}
