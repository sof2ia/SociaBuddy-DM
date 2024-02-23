package user

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
}

func (s *service) CreateUser(user User) (*User, error) {
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
	return nil
}

func NewService(userRepository Repository, userFacade Facade) Service {
	return &service{userRepository, userFacade}
}
