package main

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"socialBuddy/internal/user"
)

func main() {
	file := "../internal/database/socialbuddy.db"
	db, err := startSqlite("Users", file)
	if err != nil {
		log.Fatal(err)
		return
	}

	rep := user.NewRepository(db)
	cli := http.DefaultClient
	fac := user.NewFacade(cli)
	serv := user.NewService(rep, fac)
	ser := user.NewServer(serv)

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/v1/user", ser.GetUsers)
	router.Get("/v1/user/{id}", ser.GetUserByID)
	router.Get("/v1/user/email/{email}", ser.GetUserByEmail)
	router.Post("/v1/user", ser.CreateUser)
	router.Put("/v1/user/{id}", ser.UpdateUser)
	router.Delete("/v1/user/{id}", ser.DeleteUser)
	log.Println("server's running on the port: 8081")
	err = http.ListenAndServe(":8081", router)

	if err != nil {
		log.Fatal(err)
		return
	}
}
func startSqlite(tableName string, file string) (*sql.DB, error) {
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
