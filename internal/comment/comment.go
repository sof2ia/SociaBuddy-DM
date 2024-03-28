package comment

import (
	"errors"
	"socialBuddy/internal/post"
	"socialBuddy/internal/user"
	"time"
)

type Comment struct {
	ID          int
	IDPost      int
	IDUser      int
	DateComment time.Time
	Content     string
}

func ValidateIDPost(idPost int, servicePost post.Service) error {
	userPost, err := servicePost.GetPostByID(idPost)
	if err != nil {
		return err
	}
	if userPost == nil {
		return errors.New("the post is not in database")
	}
	return nil
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
