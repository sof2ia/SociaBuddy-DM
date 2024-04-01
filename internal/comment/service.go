package comment

import (
	"socialBuddy/internal/post"
	"socialBuddy/internal/user"
	"time"
)

type service struct {
	ComRepository  Repository
	PostRepository post.Service
	UserService    user.Service
}

type Service interface {
	CreateCom(com Comment) (*Comment, error)
	GetCom() ([]Comment, error)
	GetComByID(idCom int) (*Comment, error)
	GetComByPostID(idPost int) ([]Comment, error)
	GetComByUserID(idUser int) ([]Comment, error)
	GetComByDate(date time.Time) ([]Comment, error)
	EditCom(com Comment, idCom int) (*Comment, error)
	DeleteCom(idCom int) error
}

func (s *service) CreateCom(com Comment) (*Comment, error) {
	err := ValidateIDPost(com.IDPost, s.PostRepository)
	if err != nil {
		return nil, err
	}

	err = ValidateIDUser(com.IDUser, s.UserService)
	if err != nil {
		return nil, err
	}

	com.DateComment = time.Now()
	newPost, err := s.ComRepository.CreateCom(com)
	if err != nil {
		return nil, err
	}
	return newPost, nil
}

func (s *service) GetCom() ([]Comment, error) {
	comments, err := s.ComRepository.GetCom()
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (s *service) GetComByID(idCom int) (*Comment, error) {
	comment, err := s.ComRepository.GetComByID(idCom)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (s *service) GetComByPostID(idPost int) ([]Comment, error) {
	comments, err := s.ComRepository.GetComByPostID(idPost)
	if err != nil {
		return nil, err
	}
	return comments, nil
}
func (s *service) GetComByUserID(idUser int) ([]Comment, error) {
	comments, err := s.ComRepository.GetComByUserID(idUser)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (s *service) GetComByDate(date time.Time) ([]Comment, error) {
	comments, err := s.ComRepository.GetComByDate(date)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (s *service) EditCom(com Comment, idCom int) (*Comment, error) {
	com.DateComment = time.Now()
	comment, err := s.ComRepository.EditCom(com, idCom)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (s *service) DeleteCom(idCom int) error {
	err := s.ComRepository.DeleteCom(idCom)
	if err != nil {
		return err
	}
	return nil
}
func NewService(comRepository Repository, postService post.Service, userService user.Service) Service {
	return &service{comRepository, postService, userService}
}
