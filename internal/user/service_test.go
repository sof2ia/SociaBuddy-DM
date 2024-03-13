package user

import (
	"errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockRepository struct {
	mock.Mock
}

type mockFacade struct {
	mock.Mock
}

func (m *mockFacade) FindCep(cepUser string, number string, complement string) (*Address, error) {
	args := m.Called(cepUser, number, complement)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Address), args.Error(1)
}
func (m *mockRepository) CreateUser(user User) (*User, error) {
	args := m.Called(user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*User), args.Error(1)
}
func (m *mockRepository) GetUsers() ([]User, error) {
	args := m.Called()
	return args.Get(0).([]User), args.Error(1)
}
func (m *mockRepository) GetUserByID(idUser int) (*User, error) {
	args := m.Called(idUser)
	return args.Get(0).(*User), args.Error(1)
}
func (m *mockRepository) GetUserByEmail(emailUser string) (*User, error) {
	args := m.Called(emailUser)
	return args.Get(0).(*User), args.Error(1)
}
func (m *mockRepository) UpdateUser(user User, idUser int) (*User, error) {
	args := m.Called(user, idUser)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*User), args.Error(1)
}
func (m *mockRepository) DeleteUser(idUser int) error {
	args := m.Called(idUser)
	return args.Error(0)
}
func (m *mockRepository) DeleteALLFollowerConnections(idFollower int) error {
	args := m.Called(idFollower)
	return args.Error(0)
}
func (m *mockRepository) DeleteALLFollowingConnections(idFollowing int) error {
	args := m.Called(idFollowing)
	return args.Error(0)
}
func (m *mockRepository) FollowUser(idFollower int, idFollowing int) error {
	args := m.Called(idFollower, idFollowing)
	return args.Error(0)
}
func (m *mockRepository) DeleteConnection(idFollower int, idFollowing int) error {
	args := m.Called(idFollower, idFollowing)
	return args.Error(0)
}
func (m *mockRepository) GetFollowingByUserID(idUser int) ([]User, error) {
	args := m.Called(idUser)
	return args.Get(0).([]User), args.Error(1)
}
func (m *mockRepository) GetUserFollowers(idUser int) ([]User, error) {
	args := m.Called(idUser)
	return args.Get(0).([]User), args.Error(1)
}

func TestNewService(t *testing.T) {
	RegisterFailHandler(Fail)
	Describe("The Service Test", func() {
		Describe("Test CreateUser", func() {
			It("should CreateUser successfully", func() {
				mockUserRepository := new(mockRepository)
				mockUserFacade := new(mockFacade)
				mockUserRepository.On("CreateUser", User{
					ID:             1,
					Name:           "Name First",
					Age:            35,
					DocumentNumber: "123.345.567-89",
					Email:          "name.first@gmail.com",
					Phone:          "+55 11 92345 6789",
					Address: Address{
						ZipCode:      "12246-260",
						Country:      "Brasil",
						State:        "SP",
						City:         "São José dos Campos",
						Neighborhood: "Parque Residencial Aquarius",
						Street:       "Avenida Salmão",
						Number:       "456",
						Complement:   "C"},
				}).Return(&User{
					ID:             1,
					Name:           "Name First",
					Age:            35,
					DocumentNumber: "123.345.567-89",
					Email:          "name.first@gmail.com",
					Phone:          "+55 11 92345 6789",
					Address: Address{
						ZipCode:      "12246-260",
						Country:      "Brasil",
						State:        "SP",
						City:         "São José dos Campos",
						Neighborhood: "Parque Residencial Aquarius",
						Street:       "Avenida Salmão",
						Number:       "456",
						Complement:   "C"},
				}, nil)
				mockUserFacade.On("FindCep", "12246-260", "456", "C").Return(&Address{
					ZipCode:      "12246-260",
					Country:      "Brasil",
					State:        "SP",
					City:         "São José dos Campos",
					Neighborhood: "Parque Residencial Aquarius",
					Street:       "Avenida Salmão",
					Number:       "456",
					Complement:   "C",
				}, nil)
				newService := NewService(mockUserRepository, mockUserFacade)
				user, err := newService.CreateUser(User{
					ID:             1,
					Name:           "Name First",
					Age:            35,
					DocumentNumber: "123.345.567-89",
					Email:          "name.first@gmail.com",
					Phone:          "+55 11 92345 6789",
					Address: Address{
						ZipCode:      "12246-260",
						Country:      "Brasil",
						State:        "SP",
						City:         "São José dos Campos",
						Neighborhood: "Parque Residencial Aquarius",
						Street:       "Avenida Salmão",
						Number:       "456",
						Complement:   "C",
					},
				})
				Expect(err).ShouldNot(HaveOccurred())
				Expect(user.ID).Should(Equal(1))
				Expect(user.Name).Should(Equal("Name First"))
			})
			It("should CreateUser unsuccessfully", func() {
				mockUserRepository := new(mockRepository)
				mockUserFacade := new(mockFacade)
				mockUserRepository.On("CreateUser", User{
					ID:             1,
					Name:           "Name First",
					Age:            35,
					DocumentNumber: "123.345.567-89",
					Email:          "name.first@gmail.com",
					Phone:          "+55 11 92345 6789",
					Address: Address{
						ZipCode:      "12246-260",
						Country:      "Brasil",
						State:        "SP",
						City:         "São José dos Campos",
						Neighborhood: "Parque Residencial Aquarius",
						Street:       "Avenida Salmão",
						Number:       "456",
						Complement:   "C",
					},
				}).Return(nil, errors.New("error while CreateUser()"))
				mockUserFacade.On("FindCep", "12246-260", "456", "C").Return(nil, errors.New("error while FindCep()"))
				newService := NewService(mockUserRepository, mockUserFacade)
				user, err := newService.CreateUser(User{
					ID:             1,
					Name:           "Name First",
					Age:            35,
					DocumentNumber: "123.345.567-89",
					Email:          "name.first@gmail.com",
					Phone:          "+55 11 92345 6789",
					Address: Address{
						ZipCode:      "12246-260",
						Country:      "Brasil",
						State:        "SP",
						City:         "São José dos Campos",
						Neighborhood: "Parque Residencial Aquarius",
						Street:       "Avenida Salmão",
						Number:       "456",
						Complement:   "C",
					},
				})
				Expect(err).Should(HaveOccurred())
				Expect(user).Should(BeNil())
			})
		})
		Describe("Test GetUsers", func() {
			It("should GetUsers successfully", func() {
				mockUserRepository := new(mockRepository)
				mockUserRepository.On("GetUsers").Return([]User{
					{ID: 1,
						Name:           "Name First",
						Age:            35,
						DocumentNumber: "123.345.567-89",
						Email:          "name.first@gmail.com",
						Phone:          "+55 11 12345 6789",
						Address: Address{
							ZipCode:      "12246-260",
							Country:      "Brasil",
							State:        "SP",
							City:         "São José dos Campos",
							Neighborhood: "Parque Residencial Aquarius",
							Street:       "Avenida Salmão",
							Number:       "456",
							Complement:   "C"},
					},
				}, nil)
				newService := NewService(mockUserRepository, nil)
				users, err := newService.GetUsers()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(users[0].ID).Should(Equal(1))
				Expect(users[0].Name).Should(Equal("Name First"))
			})
			It("should GetUsers unsuccessfully", func() {
				mockUserRepository := new(mockRepository)
				mockUserRepository.On("GetUsers").Return([]User{}, errors.New("error while GetUsers()"))
				newService := NewService(mockUserRepository, nil)
				users, err := newService.GetUsers()
				Expect(err).Should(HaveOccurred())
				Expect(len(users)).Should(Equal(0))
			})
		})
		Describe("Test GetUserByID", func() {
			It("should GetUserByID successfully", func() {
				mockUserRepository := new(mockRepository)
				mockUserRepository.On("GetUserByID", 1).Return(&User{
					ID:             1,
					Name:           "Name First",
					Age:            35,
					DocumentNumber: "123.345.567-89",
					Email:          "name.first@gmail.com",
					Phone:          "+55 11 12345 6789",
					Address: Address{
						ZipCode:      "12246-260",
						Country:      "Brasil",
						State:        "SP",
						City:         "São José dos Campos",
						Neighborhood: "Parque Residencial Aquarius",
						Street:       "Avenida Salmão",
						Number:       "456",
						Complement:   "C",
					},
				}, nil)
				newService := NewService(mockUserRepository, nil)
				user, err := newService.GetUserByID(1)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(user.ID).Should(Equal(1))
				Expect(user.Name).Should(Equal("Name First"))
			})
			It("should GetUserByID unsuccessfully", func() {
				mockUserRepository := new(mockRepository)
				mockUserRepository.On("GetUserByID", 2).Return(&User{}, errors.New("error while GetUserByID()"))
				newService := NewService(mockUserRepository, nil)
				_, err := newService.GetUserByID(2)
				Expect(err).Should(HaveOccurred())
			})
		})
		Describe("Test GetUserByEmail", func() {
			It("should GetUserByEmail successfully", func() {
				mockUserRepository := new(mockRepository)
				mockUserRepository.On("GetUserByEmail", "name.first@gmail.com").Return(&User{
					ID:             1,
					Name:           "Name First",
					Age:            35,
					DocumentNumber: "123.345.567-89",
					Email:          "name.first@gmail.com",
					Phone:          "+55 11 12345 6789",
					Address: Address{
						ZipCode:      "12246-260",
						Country:      "Brasil",
						State:        "SP",
						City:         "São José dos Campos",
						Neighborhood: "Parque Residencial Aquarius",
						Street:       "Avenida Salmão",
						Number:       "456",
						Complement:   "C",
					},
				}, nil)
				newService := NewService(mockUserRepository, nil)
				user, err := newService.GetUserByEmail("name.first@gmail.com")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(user.ID).Should(Equal(1))
				Expect(user.Name).Should(Equal("Name First"))
			})
			It("should GetUserByEmail unsuccessfully", func() {
				mockUserRepository := new(mockRepository)
				mockUserRepository.On("GetUserByEmail", "name.1@gmail.com").Return(&User{}, errors.New("error while GetUserByEmail()"))
				newService := NewService(mockUserRepository, nil)
				_, err := newService.GetUserByEmail("name.1@gmail.com")
				Expect(err).Should(HaveOccurred())
			})
		})
		Describe("Test UpdateUser", func() {
			It("should UpdateUser successfully", func() {
				mockUserRepository := new(mockRepository)
				mockUserFacade := new(mockFacade)
				mockUserRepository.On("UpdateUser", User{
					ID:             1,
					Name:           "Name First",
					Age:            35,
					DocumentNumber: "123.345.567-89",
					Email:          "name.first@gmail.com",
					Phone:          "+55 11 92345 6789",
					Address: Address{
						ZipCode:      "12246-260",
						Country:      "Brasil",
						State:        "SP",
						City:         "São José dos Campos",
						Neighborhood: "Parque Residencial Aquarius",
						Street:       "Avenida Salmão",
						Number:       "456",
						Complement:   "C"},
				}, 1).Return(&User{
					ID:             1,
					Name:           "Name First",
					Age:            35,
					DocumentNumber: "123.345.567-89",
					Email:          "name.first@gmail.com",
					Phone:          "+55 11 92345 6789",
					Address: Address{
						ZipCode:      "12246-260",
						Country:      "Brasil",
						State:        "SP",
						City:         "São José dos Campos",
						Neighborhood: "Parque Residencial Aquarius",
						Street:       "Avenida Salmão",
						Number:       "456",
						Complement:   "C"},
				}, nil)
				mockUserFacade.On("FindCep", "12246-260", "456", "C").Return(&Address{
					ZipCode:      "12246-260",
					Country:      "Brasil",
					State:        "SP",
					City:         "São José dos Campos",
					Neighborhood: "Parque Residencial Aquarius",
					Street:       "Avenida Salmão",
					Number:       "456",
					Complement:   "C",
				}, nil)
				newService := NewService(mockUserRepository, mockUserFacade)
				user, err := newService.UpdateUser(User{
					ID:             1,
					Name:           "Name First",
					Age:            35,
					DocumentNumber: "123.345.567-89",
					Email:          "name.first@gmail.com",
					Phone:          "+55 11 92345 6789",
					Address: Address{
						ZipCode:      "12246-260",
						Country:      "Brasil",
						State:        "SP",
						City:         "São José dos Campos",
						Neighborhood: "Parque Residencial Aquarius",
						Street:       "Avenida Salmão",
						Number:       "456",
						Complement:   "C",
					},
				}, 1)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(user.ID).Should(Equal(1))
				Expect(user.Name).Should(Equal("Name First"))
			})
			It("should UpdateUser unsuccessfully", func() {
				mockUserRepository := new(mockRepository)
				mockUserFacade := new(mockFacade)
				mockUserRepository.On("UpdateUser", User{
					ID:             1,
					Name:           "Name First",
					Age:            35,
					DocumentNumber: "123.345.567-89",
					Email:          "name.first@gmail.com",
					Phone:          "+55 11 92345 6789",
					Address: Address{
						ZipCode:      "12246-260",
						Country:      "Brasil",
						State:        "SP",
						City:         "São José dos Campos",
						Neighborhood: "Parque Residencial Aquarius",
						Street:       "Avenida Salmão",
						Number:       "456",
						Complement:   "C",
					},
				}, 1).Return(nil, errors.New("error while UpdateUser()"))
				mockUserFacade.On("FindCep", "12246-260", "456", "C").Return(nil, errors.New("error while FindCep()"))
				newService := NewService(mockUserRepository, mockUserFacade)
				user, err := newService.UpdateUser(User{
					ID:             1,
					Name:           "Name First",
					Age:            35,
					DocumentNumber: "123.345.567-89",
					Email:          "name.first@gmail.com",
					Phone:          "+55 11 92345 6789",
					Address: Address{
						ZipCode:      "12246-260",
						Country:      "Brasil",
						State:        "SP",
						City:         "São José dos Campos",
						Neighborhood: "Parque Residencial Aquarius",
						Street:       "Avenida Salmão",
						Number:       "456",
						Complement:   "C",
					},
				}, 1)
				Expect(err).Should(HaveOccurred())
				Expect(user).Should(BeNil())
			})
		})
		Describe("Test DeleteUser", func() {
			It("should DeleteUser successfully", func() {
				mockUserRepository := new(mockRepository)
				mockUserRepository.On("DeleteUser", 1).Return(nil)
				mockUserRepository.On("DeleteALLFollowerConnections", 1).Return(nil)
				mockUserRepository.On("DeleteALLFollowingConnections", 1).Return(nil)
				newService := NewService(mockUserRepository, nil)
				err := newService.DeleteUser(1)
				Expect(err).ShouldNot(HaveOccurred())
			})
			It("should DeleteUser unsuccessfully", func() {
				mockUserRepository := new(mockRepository)
				mockUserRepository.On("DeleteUser", 1).Return(errors.New("error while DeleteUser()"))
				mockUserRepository.On("DeleteALLFollowerConnections", 1).Return(errors.New("error while DeleteALLFollowerConnections()"))
				mockUserRepository.On("DeleteALLFollowingConnections", 1).Return(errors.New("error while DeleteALLFollowingConnections()"))
				newService := NewService(mockUserRepository, nil)
				err := newService.DeleteUser(1)
				Expect(err).Should(HaveOccurred())
			})
		})
		Describe("Test FollowUser", func() {
			It("should FollowUser successfully", func() {
				mockUserRepository := new(mockRepository)
				mockUserRepository.On("GetUserByID", 1).Return(&User{
					ID:             1,
					Name:           "Name First",
					Age:            35,
					DocumentNumber: "123.345.567-89",
					Email:          "name.first@gmail.com",
					Phone:          "+55 11 12345 6789",
					Address: Address{
						ZipCode:      "12246-260",
						Country:      "Brasil",
						State:        "SP",
						City:         "São José dos Campos",
						Neighborhood: "Parque Residencial Aquarius",
						Street:       "Avenida Salmão",
						Number:       "456",
						Complement:   "C",
					},
				}, nil)
				mockUserRepository.On("GetUserByID", 2).Return(&User{
					ID:             1,
					Name:           "Name First",
					Age:            35,
					DocumentNumber: "123.345.567-89",
					Email:          "name.first@gmail.com",
					Phone:          "+55 11 12345 6789",
					Address: Address{
						ZipCode:      "12246-260",
						Country:      "Brasil",
						State:        "SP",
						City:         "São José dos Campos",
						Neighborhood: "Parque Residencial Aquarius",
						Street:       "Avenida Salmão",
						Number:       "456",
						Complement:   "C",
					},
				}, nil)
				mockUserRepository.On("GetFollowingByUserID", 1).Return([]User{
					{ID: 1,
						Name:           "Name First",
						Age:            35,
						DocumentNumber: "123.345.567-89",
						Email:          "name.first@gmail.com",
						Phone:          "+55 11 12345 6789",
						Address: Address{
							ZipCode:      "12246-260",
							Country:      "Brasil",
							State:        "SP",
							City:         "São José dos Campos",
							Neighborhood: "Parque Residencial Aquarius",
							Street:       "Avenida Salmão",
							Number:       "456",
							Complement:   "C"},
					},
				}, nil)
				mockUserRepository.On("FollowUser", 1, 2).Return(nil)
				newService := NewService(mockUserRepository, nil)
				err := newService.FollowUser(1, 2)
				Expect(err).ShouldNot(HaveOccurred())
			})
			It("should FollowUser unsuccessfully", func() {
				mockUserRepository := new(mockRepository)
				mockUserRepository.On("GetUserByID", 1).Return(&User{}, errors.New("error while GetUserByID(follower)"))
				mockUserRepository.On("GetUserByID", 2).Return(&User{}, errors.New("error while GetUserByID(following)"))
				mockUserRepository.On("GetFollowingByUserID", 1).Return([]User{}, errors.New("error while GetFollowingByUserID()"))
				mockUserRepository.On("FollowUser", 1, 2).Return(errors.New("error while FollowUser()"))
				newService := NewService(mockUserRepository, nil)
				err := newService.FollowUser(1, 2)
				Expect(err).Should(HaveOccurred())
			})
		})
		Describe("Test DeleteConnection", func() {
			It("should DeleteConnection successfully", func() {
				mockUserRepository := new(mockRepository)
				mockUserRepository.On("DeleteConnection", 1, 2).Return(nil)
				newService := NewService(mockUserRepository, nil)
				err := newService.DeleteConnection(1, 2)
				Expect(err).ShouldNot(HaveOccurred())
			})
			It("should DeleteConnection unsuccessfully", func() {
				mockUserRepository := new(mockRepository)
				mockUserRepository.On("DeleteConnection", 1, 2).Return(errors.New("error while DeleteConnection()"))
				newService := NewService(mockUserRepository, nil)
				err := newService.DeleteConnection(1, 2)
				Expect(err).Should(HaveOccurred())
			})
		})
		Describe("Test GetFollowingByUserID", func() {
			It("should GetFollowingByUserID successfully", func() {
				mockUserRepository := new(mockRepository)
				mockUserRepository.On("GetFollowingByUserID", 2).Return([]User{
					{ID: 1,
						Name:           "Name First",
						Age:            35,
						DocumentNumber: "123.345.567-89",
						Email:          "name.first@gmail.com",
						Phone:          "+55 11 12345 6789",
						Address: Address{
							ZipCode:      "12246-260",
							Country:      "Brasil",
							State:        "SP",
							City:         "São José dos Campos",
							Neighborhood: "Parque Residencial Aquarius",
							Street:       "Avenida Salmão",
							Number:       "456",
							Complement:   "C"},
					},
				}, nil)
				newService := NewService(mockUserRepository, nil)
				users, err := newService.GetFollowingByUserID(2)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(users[0].ID).Should(Equal(1))
				Expect(users[0].Name).Should(Equal("Name First"))
			})
			It("should GetFollowingByUserID unsuccessfully", func() {
				mockUserRepository := new(mockRepository)
				mockUserRepository.On("GetFollowingByUserID", 2).Return([]User{}, errors.New("error while GetFollowingByUserID()"))
				newService := NewService(mockUserRepository, nil)
				users, err := newService.GetFollowingByUserID(2)
				Expect(err).Should(HaveOccurred())
				Expect(len(users)).Should(Equal(0))
			})
		})
		Describe("Test GetUserFollowers", func() {
			It("should GetUserFollowers successfully", func() {
				mockUserRepository := new(mockRepository)
				mockUserRepository.On("GetUserFollowers", 2).Return([]User{
					{ID: 1,
						Name:           "Name First",
						Age:            35,
						DocumentNumber: "123.345.567-89",
						Email:          "name.first@gmail.com",
						Phone:          "+55 11 12345 6789",
						Address: Address{
							ZipCode:      "12246-260",
							Country:      "Brasil",
							State:        "SP",
							City:         "São José dos Campos",
							Neighborhood: "Parque Residencial Aquarius",
							Street:       "Avenida Salmão",
							Number:       "456",
							Complement:   "C"},
					},
				}, nil)
				newService := NewService(mockUserRepository, nil)
				users, err := newService.GetUserFollowers(2)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(users[0].ID).Should(Equal(1))
				Expect(users[0].Name).Should(Equal("Name First"))
			})
			It("should GetUserFollowers unsuccessfully", func() {
				mockUserRepository := new(mockRepository)
				mockUserRepository.On("GetUserFollowers", 2).Return([]User{}, errors.New("error while GetUserFollowers()"))
				newService := NewService(mockUserRepository, nil)
				users, err := newService.GetUserFollowers(2)
				Expect(err).Should(HaveOccurred())
				Expect(len(users)).Should(Equal(0))
			})
		})
	})
	RunSpecs(t, "The Service Test")
}
func TestService_CreateUser(t *testing.T) {
	RegisterFailHandler(Fail)
	Describe("The Service Test - CreateUser", func() {
		It("should CreateUser successfully", func() {
			mockUserRepository := new(mockRepository)
			mockUserFacade := new(mockFacade)
			mockUserRepository.On("CreateUser", User{
				ID:             1,
				Name:           "Name First",
				Age:            35,
				DocumentNumber: "123.345.567-89",
				Email:          "name.first@gmail.com",
				Phone:          "+55 11 92345 6789",
				Address: Address{
					ZipCode:      "12246-260",
					Country:      "Brasil",
					State:        "SP",
					City:         "São José dos Campos",
					Neighborhood: "Parque Residencial Aquarius",
					Street:       "Avenida Salmão",
					Number:       "456",
					Complement:   "C"},
			}).Return(&User{
				ID:             1,
				Name:           "Name First",
				Age:            35,
				DocumentNumber: "123.345.567-89",
				Email:          "name.first@gmail.com",
				Phone:          "+55 11 92345 6789",
				Address: Address{
					ZipCode:      "12246-260",
					Country:      "Brasil",
					State:        "SP",
					City:         "São José dos Campos",
					Neighborhood: "Parque Residencial Aquarius",
					Street:       "Avenida Salmão",
					Number:       "456",
					Complement:   "C"},
			}, nil)
			mockUserFacade.On("FindCep", "12246-260", "456", "C").Return(&Address{
				ZipCode:      "12246-260",
				Country:      "Brasil",
				State:        "SP",
				City:         "São José dos Campos",
				Neighborhood: "Parque Residencial Aquarius",
				Street:       "Avenida Salmão",
				Number:       "456",
				Complement:   "C",
			}, nil)
			newService := NewService(mockUserRepository, mockUserFacade)
			user, err := newService.CreateUser(User{
				ID:             1,
				Name:           "Name First",
				Age:            35,
				DocumentNumber: "123.345.567-89",
				Email:          "name.first@gmail.com",
				Phone:          "+55 11 92345 6789",
				Address: Address{
					ZipCode:      "12246-260",
					Country:      "Brasil",
					State:        "SP",
					City:         "São José dos Campos",
					Neighborhood: "Parque Residencial Aquarius",
					Street:       "Avenida Salmão",
					Number:       "456",
					Complement:   "C",
				},
			})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(user.ID).Should(Equal(1))
			Expect(user.Name).Should(Equal("Name First"))
		})
		It("should CreateUser unsuccessfully", func() {
			mockUserRepository := new(mockRepository)
			mockUserFacade := new(mockFacade)
			mockUserRepository.On("CreateUser", User{
				ID:             1,
				Name:           "Name First",
				Age:            35,
				DocumentNumber: "123.345.567-89",
				Email:          "name.first@gmail.com",
				Phone:          "+55 11 92345 6789",
				Address: Address{
					ZipCode:      "12246-260",
					Country:      "Brasil",
					State:        "SP",
					City:         "São José dos Campos",
					Neighborhood: "Parque Residencial Aquarius",
					Street:       "Avenida Salmão",
					Number:       "456",
					Complement:   "C",
				},
			}).Return(nil, errors.New("error while CreateUser()"))
			mockUserFacade.On("FindCep", "12246-260", "456", "C").Return(nil, errors.New("error while FindCep()"))
			newService := NewService(mockUserRepository, mockUserFacade)
			user, err := newService.CreateUser(User{
				ID:             1,
				Name:           "Name First",
				Age:            35,
				DocumentNumber: "123.345.567-89",
				Email:          "name.first@gmail.com",
				Phone:          "+55 11 92345 6789",
				Address: Address{
					ZipCode:      "12246-260",
					Country:      "Brasil",
					State:        "SP",
					City:         "São José dos Campos",
					Neighborhood: "Parque Residencial Aquarius",
					Street:       "Avenida Salmão",
					Number:       "456",
					Complement:   "C",
				},
			})
			Expect(err).Should(HaveOccurred())
			Expect(user).Should(BeNil())
		})
	})
	RunSpecs(t, "The Service Test - CreateUser")
}
func TestService_GetUsers(t *testing.T) {
	RegisterFailHandler(Fail)
	Describe("The Service Test - GetUsers", func() {
		It("should GetUsers successfully", func() {
			mockUserRepository := new(mockRepository)
			mockUserRepository.On("GetUsers").Return([]User{
				{ID: 1,
					Name:           "Name First",
					Age:            35,
					DocumentNumber: "123.345.567-89",
					Email:          "name.first@gmail.com",
					Phone:          "+55 11 12345 6789",
					Address: Address{
						ZipCode:      "12246-260",
						Country:      "Brasil",
						State:        "SP",
						City:         "São José dos Campos",
						Neighborhood: "Parque Residencial Aquarius",
						Street:       "Avenida Salmão",
						Number:       "456",
						Complement:   "C"},
				},
			}, nil)
			newService := NewService(mockUserRepository, nil)
			users, err := newService.GetUsers()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(users[0].ID).Should(Equal(1))
			Expect(users[0].Name).Should(Equal("Name First"))
		})
		It("should GetUsers unsuccessfully", func() {
			mockUserRepository := new(mockRepository)
			mockUserRepository.On("GetUsers").Return([]User{}, errors.New("error while GetUsers()"))
			newService := NewService(mockUserRepository, nil)
			users, err := newService.GetUsers()
			Expect(err).Should(HaveOccurred())
			Expect(len(users)).Should(Equal(0))
		})
	})

	RunSpecs(t, "The Service Test - GetUsers")
}
func TestService_GetUserByID(t *testing.T) {
	RegisterFailHandler(Fail)
	Describe("The Service Test - GetUserByID", func() {
		It("should GetUserByID successfully", func() {
			mockUserRepository := new(mockRepository)
			mockUserRepository.On("GetUserByID", 1).Return(&User{
				ID:             1,
				Name:           "Name First",
				Age:            35,
				DocumentNumber: "123.345.567-89",
				Email:          "name.first@gmail.com",
				Phone:          "+55 11 12345 6789",
				Address: Address{
					ZipCode:      "12246-260",
					Country:      "Brasil",
					State:        "SP",
					City:         "São José dos Campos",
					Neighborhood: "Parque Residencial Aquarius",
					Street:       "Avenida Salmão",
					Number:       "456",
					Complement:   "C",
				},
			}, nil)
			newService := NewService(mockUserRepository, nil)
			user, err := newService.GetUserByID(1)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(user.ID).Should(Equal(1))
			Expect(user.Name).Should(Equal("Name First"))
		})
		It("should GetUserByID unsuccessfully", func() {
			mockUserRepository := new(mockRepository)
			mockUserRepository.On("GetUserByID", 2).Return(&User{}, errors.New("error while GetUserByID()"))
			newService := NewService(mockUserRepository, nil)
			_, err := newService.GetUserByID(2)
			Expect(err).Should(HaveOccurred())
		})
	})
	RunSpecs(t, "The Service Test - GetUserByID")
}
func TestService_GetUserByEmail(t *testing.T) {
	RegisterFailHandler(Fail)
	Describe("The Service Test - GetUserByEmail", func() {
		It("should GetUserByEmail successfully", func() {
			mockUserRepository := new(mockRepository)
			mockUserRepository.On("GetUserByEmail", "name.first@gmail.com").Return(&User{
				ID:             1,
				Name:           "Name First",
				Age:            35,
				DocumentNumber: "123.345.567-89",
				Email:          "name.first@gmail.com",
				Phone:          "+55 11 12345 6789",
				Address: Address{
					ZipCode:      "12246-260",
					Country:      "Brasil",
					State:        "SP",
					City:         "São José dos Campos",
					Neighborhood: "Parque Residencial Aquarius",
					Street:       "Avenida Salmão",
					Number:       "456",
					Complement:   "C",
				},
			}, nil)
			newService := NewService(mockUserRepository, nil)
			user, err := newService.GetUserByEmail("name.first@gmail.com")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(user.ID).Should(Equal(1))
			Expect(user.Name).Should(Equal("Name First"))
		})
		It("should GetUserByEmail unsuccessfully", func() {
			mockUserRepository := new(mockRepository)
			mockUserRepository.On("GetUserByEmail", "name.1@gmail.com").Return(&User{}, errors.New("error while GetUserByEmail()"))
			newService := NewService(mockUserRepository, nil)
			_, err := newService.GetUserByEmail("name.1@gmail.com")
			Expect(err).Should(HaveOccurred())
		})
	})
	RunSpecs(t, "The Service Test - GetUserByEmail")
}
func TestService_UpdateUser(t *testing.T) {
	RegisterFailHandler(Fail)
	Describe("Test UpdateUser", func() {
		It("should UpdateUser successfully", func() {
			mockUserRepository := new(mockRepository)
			mockUserFacade := new(mockFacade)
			mockUserRepository.On("UpdateUser", User{
				ID:             1,
				Name:           "Name First",
				Age:            35,
				DocumentNumber: "123.345.567-89",
				Email:          "name.first@gmail.com",
				Phone:          "+55 11 92345 6789",
				Address: Address{
					ZipCode:      "12246-260",
					Country:      "Brasil",
					State:        "SP",
					City:         "São José dos Campos",
					Neighborhood: "Parque Residencial Aquarius",
					Street:       "Avenida Salmão",
					Number:       "456",
					Complement:   "C"},
			}, 1).Return(&User{
				ID:             1,
				Name:           "Name First",
				Age:            35,
				DocumentNumber: "123.345.567-89",
				Email:          "name.first@gmail.com",
				Phone:          "+55 11 92345 6789",
				Address: Address{
					ZipCode:      "12246-260",
					Country:      "Brasil",
					State:        "SP",
					City:         "São José dos Campos",
					Neighborhood: "Parque Residencial Aquarius",
					Street:       "Avenida Salmão",
					Number:       "456",
					Complement:   "C"},
			}, nil)
			mockUserFacade.On("FindCep", "12246-260", "456", "C").Return(&Address{
				ZipCode:      "12246-260",
				Country:      "Brasil",
				State:        "SP",
				City:         "São José dos Campos",
				Neighborhood: "Parque Residencial Aquarius",
				Street:       "Avenida Salmão",
				Number:       "456",
				Complement:   "C",
			}, nil)
			newService := NewService(mockUserRepository, mockUserFacade)
			user, err := newService.UpdateUser(User{
				ID:             1,
				Name:           "Name First",
				Age:            35,
				DocumentNumber: "123.345.567-89",
				Email:          "name.first@gmail.com",
				Phone:          "+55 11 92345 6789",
				Address: Address{
					ZipCode:      "12246-260",
					Country:      "Brasil",
					State:        "SP",
					City:         "São José dos Campos",
					Neighborhood: "Parque Residencial Aquarius",
					Street:       "Avenida Salmão",
					Number:       "456",
					Complement:   "C",
				},
			}, 1)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(user.ID).Should(Equal(1))
			Expect(user.Name).Should(Equal("Name First"))
		})
		It("should UpdateUser unsuccessfully", func() {
			mockUserRepository := new(mockRepository)
			mockUserFacade := new(mockFacade)
			mockUserRepository.On("UpdateUser", User{
				ID:             1,
				Name:           "Name First",
				Age:            35,
				DocumentNumber: "123.345.567-89",
				Email:          "name.first@gmail.com",
				Phone:          "+55 11 92345 6789",
				Address: Address{
					ZipCode:      "12246-260",
					Country:      "Brasil",
					State:        "SP",
					City:         "São José dos Campos",
					Neighborhood: "Parque Residencial Aquarius",
					Street:       "Avenida Salmão",
					Number:       "456",
					Complement:   "C",
				},
			}, 1).Return(nil, errors.New("error while UpdateUser()"))
			mockUserFacade.On("FindCep", "12246-260", "456", "C").Return(nil, errors.New("error while FindCep()"))
			newService := NewService(mockUserRepository, mockUserFacade)
			user, err := newService.UpdateUser(User{
				ID:             1,
				Name:           "Name First",
				Age:            35,
				DocumentNumber: "123.345.567-89",
				Email:          "name.first@gmail.com",
				Phone:          "+55 11 92345 6789",
				Address: Address{
					ZipCode:      "12246-260",
					Country:      "Brasil",
					State:        "SP",
					City:         "São José dos Campos",
					Neighborhood: "Parque Residencial Aquarius",
					Street:       "Avenida Salmão",
					Number:       "456",
					Complement:   "C",
				},
			}, 1)
			Expect(err).Should(HaveOccurred())
			Expect(user).Should(BeNil())
		})
	})
	RunSpecs(t, "The Service Test - UpdateUser")
}
func TestService_DeleteUser(t *testing.T) {
	RegisterFailHandler(Fail)
	Describe("Test DeleteUser", func() {
		It("should DeleteUser successfully", func() {
			mockUserRepository := new(mockRepository)
			mockUserRepository.On("DeleteUser", 1).Return(nil)
			mockUserRepository.On("DeleteALLFollowerConnections", 1).Return(nil)
			mockUserRepository.On("DeleteALLFollowingConnections", 1).Return(nil)
			newService := NewService(mockUserRepository, nil)
			err := newService.DeleteUser(1)
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("should DeleteUser unsuccessfully", func() {
			mockUserRepository := new(mockRepository)
			mockUserRepository.On("DeleteUser", 1).Return(errors.New("error while DeleteUser()"))
			mockUserRepository.On("DeleteALLFollowerConnections", 1).Return(errors.New("error while DeleteALLFollowerConnections()"))
			mockUserRepository.On("DeleteALLFollowingConnections", 1).Return(errors.New("error while DeleteALLFollowingConnections()"))
			newService := NewService(mockUserRepository, nil)
			err := newService.DeleteUser(1)
			Expect(err).Should(HaveOccurred())
		})
	})
	RunSpecs(t, "The Service Test - DeleteUser")
}
func TestService_FollowUser(t *testing.T) {
	RegisterFailHandler(Fail)
	Describe("Test FollowUser", func() {
		It("should FollowUser successfully", func() {
			mockUserRepository := new(mockRepository)
			mockUserRepository.On("GetUserByID", 1).Return(&User{
				ID:             1,
				Name:           "Name First",
				Age:            35,
				DocumentNumber: "123.345.567-89",
				Email:          "name.first@gmail.com",
				Phone:          "+55 11 12345 6789",
				Address: Address{
					ZipCode:      "12246-260",
					Country:      "Brasil",
					State:        "SP",
					City:         "São José dos Campos",
					Neighborhood: "Parque Residencial Aquarius",
					Street:       "Avenida Salmão",
					Number:       "456",
					Complement:   "C",
				},
			}, nil)
			mockUserRepository.On("GetUserByID", 2).Return(&User{
				ID:             1,
				Name:           "Name First",
				Age:            35,
				DocumentNumber: "123.345.567-89",
				Email:          "name.first@gmail.com",
				Phone:          "+55 11 12345 6789",
				Address: Address{
					ZipCode:      "12246-260",
					Country:      "Brasil",
					State:        "SP",
					City:         "São José dos Campos",
					Neighborhood: "Parque Residencial Aquarius",
					Street:       "Avenida Salmão",
					Number:       "456",
					Complement:   "C",
				},
			}, nil)
			mockUserRepository.On("GetFollowingByUserID", 1).Return([]User{
				{ID: 1,
					Name:           "Name First",
					Age:            35,
					DocumentNumber: "123.345.567-89",
					Email:          "name.first@gmail.com",
					Phone:          "+55 11 12345 6789",
					Address: Address{
						ZipCode:      "12246-260",
						Country:      "Brasil",
						State:        "SP",
						City:         "São José dos Campos",
						Neighborhood: "Parque Residencial Aquarius",
						Street:       "Avenida Salmão",
						Number:       "456",
						Complement:   "C"},
				},
			}, nil)
			mockUserRepository.On("FollowUser", 1, 2).Return(nil)
			newService := NewService(mockUserRepository, nil)
			err := newService.FollowUser(1, 2)
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("should FollowUser unsuccessfully", func() {
			mockUserRepository := new(mockRepository)
			mockUserRepository.On("GetUserByID", 1).Return(&User{}, errors.New("error while GetUserByID(follower)"))
			mockUserRepository.On("GetUserByID", 2).Return(&User{}, errors.New("error while GetUserByID(following)"))
			mockUserRepository.On("GetFollowingByUserID", 1).Return([]User{}, errors.New("error while GetFollowingByUserID()"))
			mockUserRepository.On("FollowUser", 1, 2).Return(errors.New("error while FollowUser()"))
			newService := NewService(mockUserRepository, nil)
			err := newService.FollowUser(1, 2)
			Expect(err).Should(HaveOccurred())
		})
	})
	RunSpecs(t, "The Service Test - FollowUser")
}
func TestService_DeleteConnection(t *testing.T) {
	RegisterFailHandler(Fail)
	Describe("Test DeleteConnection", func() {
		It("should DeleteConnection successfully", func() {
			mockUserRepository := new(mockRepository)
			mockUserRepository.On("DeleteConnection", 1, 2).Return(nil)
			newService := NewService(mockUserRepository, nil)
			err := newService.DeleteConnection(1, 2)
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("should DeleteConnection unsuccessfully", func() {
			mockUserRepository := new(mockRepository)
			mockUserRepository.On("DeleteConnection", 1, 2).Return(errors.New("error while DeleteConnection()"))
			newService := NewService(mockUserRepository, nil)
			err := newService.DeleteConnection(1, 2)
			Expect(err).Should(HaveOccurred())
		})
	})
	RunSpecs(t, "The Service Test - DeleteConnection")
}
func TestService_GetFollowingByUserID(t *testing.T) {
	RegisterFailHandler(Fail)
	Describe("Test GetFollowingByUserID", func() {
		It("should GetFollowingByUserID successfully", func() {
			mockUserRepository := new(mockRepository)
			mockUserRepository.On("GetFollowingByUserID", 2).Return([]User{
				{ID: 1,
					Name:           "Name First",
					Age:            35,
					DocumentNumber: "123.345.567-89",
					Email:          "name.first@gmail.com",
					Phone:          "+55 11 12345 6789",
					Address: Address{
						ZipCode:      "12246-260",
						Country:      "Brasil",
						State:        "SP",
						City:         "São José dos Campos",
						Neighborhood: "Parque Residencial Aquarius",
						Street:       "Avenida Salmão",
						Number:       "456",
						Complement:   "C"},
				},
			}, nil)
			newService := NewService(mockUserRepository, nil)
			users, err := newService.GetFollowingByUserID(2)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(users[0].ID).Should(Equal(1))
			Expect(users[0].Name).Should(Equal("Name First"))
		})
		It("should GetFollowingByUserID unsuccessfully", func() {
			mockUserRepository := new(mockRepository)
			mockUserRepository.On("GetFollowingByUserID", 2).Return([]User{}, errors.New("error while GetFollowingByUserID()"))
			newService := NewService(mockUserRepository, nil)
			users, err := newService.GetFollowingByUserID(2)
			Expect(err).Should(HaveOccurred())
			Expect(len(users)).Should(Equal(0))
		})
	})

	RunSpecs(t, "The Service Test - GetFollowingByUserID")
}
func TestService_GetUserFollowers(t *testing.T) {
	RegisterFailHandler(Fail)
	Describe("Test GetUserFollowers", func() {
		It("should GetUserFollowers successfully", func() {
			mockUserRepository := new(mockRepository)
			mockUserRepository.On("GetUserFollowers", 2).Return([]User{
				{ID: 1,
					Name:           "Name First",
					Age:            35,
					DocumentNumber: "123.345.567-89",
					Email:          "name.first@gmail.com",
					Phone:          "+55 11 12345 6789",
					Address: Address{
						ZipCode:      "12246-260",
						Country:      "Brasil",
						State:        "SP",
						City:         "São José dos Campos",
						Neighborhood: "Parque Residencial Aquarius",
						Street:       "Avenida Salmão",
						Number:       "456",
						Complement:   "C"},
				},
			}, nil)
			newService := NewService(mockUserRepository, nil)
			users, err := newService.GetUserFollowers(2)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(users[0].ID).Should(Equal(1))
			Expect(users[0].Name).Should(Equal("Name First"))
		})
		It("should GetUserFollowers unsuccessfully", func() {
			mockUserRepository := new(mockRepository)
			mockUserRepository.On("GetUserFollowers", 2).Return([]User{}, errors.New("error while GetUserFollowers()"))
			newService := NewService(mockUserRepository, nil)
			users, err := newService.GetUserFollowers(2)
			Expect(err).Should(HaveOccurred())
			Expect(len(users)).Should(Equal(0))
		})
	})
	RunSpecs(t, "The Service Test - GetUserFollowers")
}
