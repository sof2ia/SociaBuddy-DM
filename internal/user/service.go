package user

import (
	"errors"
	"log"
	"time"
)

type service struct {
	UserRepository Repository
	UserFacade     Facade
}

type Service interface {
	CreateUser(user User) (*User, error)
	GetUsers() ([]User, error)
	GetUserByID(idUser int) (*User, error)
	GetUserByEmail(emailUser string) (*User, error)
	UpdateUser(user User, idUser int) (*User, error)
	DeleteUser(idUser int) error
	FollowUser(idFollower int, idFollowing int) error
	DeleteConnection(idFollower int, idFollowing int) error
	GetFollowingByUserID(idUser int) ([]User, error)
	GetUserFollowers(idUser int) ([]User, error)
}

func (s *service) CreateUser(user User) (*User, error) {

	err := nameValidation(user.Name)
	if err != nil {
		return nil, err
	}
	err = ageValidation(user.Age)
	if err != nil {
		return nil, err
	}
	err = documentValidation(user.DocumentNumber)
	if err != nil {
		return nil, err
	}
	err = emailValidation(user.Email)
	if err != nil {
		return nil, err
	}
	err = phoneValidation(user.Phone)
	if err != nil {
		return nil, err
	}
	err = zipCodeValidation(user.Address.ZipCode)
	if err != nil {
		return nil, err
	}
	err = countryValidation(user.Address.Country)
	if err != nil {
		return nil, err
	}
	err = numberValidation(user.Address.Number)
	if err != nil {
		return nil, err
	}

	addressUser, err := s.UserFacade.FindCep(user.Address.ZipCode, user.Address.Number, user.Address.Complement)
	if err != nil {
		return nil, err
	}
	user.Address = *addressUser

	newUser, err := s.UserRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return newUser, nil

}

func (s *service) GetUsers() ([]User, error) {
	users, err := s.UserRepository.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (s *service) GetUserByID(idUser int) (*User, error) {
	users, err := s.UserRepository.GetUserByID(idUser)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *service) GetUserByEmail(emailUser string) (*User, error) {
	users, err := s.UserRepository.GetUserByEmail(emailUser)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *service) UpdateUser(user User, idUser int) (*User, error) {
	addressUser, err := s.UserFacade.FindCep(user.Address.ZipCode, user.Address.Number, user.Address.Complement)
	if err != nil {
		return nil, err
	}
	user.Address = *addressUser
	users, err := s.UserRepository.UpdateUser(user, idUser)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *service) DeleteUser(idUser int) error {
	err := s.UserRepository.DeleteUser(idUser)
	if err != nil {
		return err
	}

	err = s.UserRepository.DeleteALLFollowerConnections(idUser)
	if err != nil {
		return err
	}

	err = s.UserRepository.DeleteALLFollowingConnections(idUser)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) FollowUser(idFollower int, idFollowing int) error {
	if idFollower == idFollowing {
		return errors.New("the id cannot follow itself")
	}
	log.Printf("start")
	// suggestion:
	// - prevent the connection between users that are not listed in the Users table //
	accountFollower, err := s.UserRepository.GetUserByID(idFollower)
	if err != nil {
		return err
	}
	if accountFollower == nil {
		return errors.New("id has no account")
	}
	log.Printf("follower with no account")
	accountFollowing, err := s.UserRepository.GetUserByID(idFollowing)
	if err != nil {
		return err
	}
	if accountFollowing == nil {
		return errors.New("id cannot follow user with no account")
	}
	log.Printf("following with no account")
	followers, err := s.UserRepository.GetFollowingByUserID(idFollower)
	for i := 0; i < len(followers); i++ {
		if idFollowing == followers[i].ID {
			return errors.New("the id cannot follow user more than once")
		}
	}
	time.Sleep(2 * time.Second)
	log.Printf("REPOSITORY")
	err = s.UserRepository.FollowUser(idFollower, idFollowing)
	if err != nil {
		return err
	}
	log.Printf("final")
	return nil
}
func (s *service) DeleteConnection(idFollower int, idFollowing int) error {

	// suggestions:
	// - delete all connections (follower <-> follower) with the user in case his account has been deleted //

	err := s.UserRepository.DeleteConnection(idFollower, idFollowing)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) GetFollowingByUserID(idUser int) ([]User, error) {
	users, err := s.UserRepository.GetFollowingByUserID(idUser)
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (s *service) GetUserFollowers(idUser int) ([]User, error) {
	users, err := s.UserRepository.GetUserFollowers(idUser)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func NewService(userRepository Repository, userFacade Facade) Service {
	return &service{userRepository, userFacade}
}
