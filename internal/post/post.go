package post

import (
	"errors"
	"socialBuddy/internal/user"
	"time"
)

type Post struct {
	ID      int
	IDUser  int
	Date    time.Time
	Title   string
	Content string
}

func ValidateIDUser(idUser int, serviceUser user.Service) error {
	userPost, err := serviceUser.GetUserByID(idUser)
	if err != nil {
		return err
	}
	if userPost == nil {
		return errors.New("the user is not in database")
	}
	return nil
}
