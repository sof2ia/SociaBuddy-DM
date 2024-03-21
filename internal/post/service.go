package post

import "time"

type service struct {
	PostRepository Repository
}

type Service interface {
	CreatePost(post Post) (*Post, error)
	GetPosts() ([]Post, error)
	GetPostByID(idPost int) (*Post, error)
	GetPostByUserID(idUser int) ([]Post, error)
	GetPostByDate(date time.Time) ([]Post, error)
	GetPostByTitle(title string) ([]Post, error)
	EditPost(post Post, idPost int) (*Post, error)
	DeletePost(idPost int) error
	DeleteAllPostsByUserID(idUser int) error
}

func (s *service) CreatePost(post Post) (*Post, error) {
	post.Date = time.Now()
	newPost, err := s.PostRepository.CreatePost(post)
	if err != nil {
		return nil, err
	}
	return newPost, nil
}

func (s *service) GetPosts() ([]Post, error) {
	posts, err := s.PostRepository.GetPosts()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *service) GetPostByID(idPost int) (*Post, error) {
	post, err := s.PostRepository.GetPostByID(idPost)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *service) GetPostByUserID(idUser int) ([]Post, error) {
	posts, err := s.PostRepository.GetPostByUserID(idUser)
	if err != nil {
		return nil, err
	}
	return posts, err
}

func (s *service) GetPostByDate(date time.Time) ([]Post, error) {
	posts, err := s.PostRepository.GetPostByDate(date)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *service) GetPostByTitle(title string) ([]Post, error) {
	post, err := s.PostRepository.GetPostByTitle(title)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *service) EditPost(editPost Post, idPost int) (*Post, error) {
	editPost.Date = time.Now()
	post, err := s.PostRepository.EditPost(editPost, idPost)
	if err != nil {
		return nil, err
	}
	return post, nil

}

func (s *service) DeletePost(idPost int) error {
	err := s.PostRepository.DeletePost(idPost)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteAllPostsByUserID(idUser int) error {
	err := s.PostRepository.DeleteAllPostsByUserID(idUser)
	if err != nil {
		return err
	}
	return nil
}

func NewService(postRepository Repository) Service {
	return &service{postRepository}
}
