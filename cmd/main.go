package main

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"socialBuddy/internal/post"
	"socialBuddy/internal/user"
)

func main() {
	file := "../internal/database/socialbuddy.db"

	db, err := startSqliteUser("Users", file)
	if err != nil {
		log.Fatal(err)
		return
	}
	postDB, err := startSqlitePost("Post", file)
	if err != nil {
		log.Fatal(err)
		return
	}

	repUser := user.NewRepository(db)
	cli := http.DefaultClient
	fac := user.NewFacade("https://viacep.com.br", cli)
	servUser := user.NewService(repUser, fac)
	serUser := user.NewServer(servUser)
	log.Println("module user")

	repPost := post.NewRepository(postDB)
	log.Println("newRepository")
	servPost := post.NewService(repPost)
	log.Println("newService")
	serPost := post.NewServer(servPost)
	log.Println("newServer")

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

	log.Println("server's running on the port: 8081")
	err = http.ListenAndServe(":8081", router)

	if err != nil {
		log.Fatal(err)
		return
	}
}
func startSqliteUser(tableName string, file string) (*sql.DB, error) {
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
    												)    `, tableName))
	if err != nil {
		err := db.Close()
		if err != nil {
			return nil, err
		}
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}

func startSqlitePost(tableName string, file string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	_, err = db.Exec(fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
    												ID INTEGER PRIMARY KEY AUTOINCREMENT,
    												IDUser INTEGER FOREIGN KEY,
    												DatePost DATE,
    												Title TEXT,
    												Content TEXT,
    												)    `, tableName))
	if err != nil {
		err := db.Close()
		if err != nil {
			return nil, err
		}
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}
