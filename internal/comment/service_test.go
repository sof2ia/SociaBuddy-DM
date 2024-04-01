package comment

import (
	"errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"socialBuddy/internal/post"
	"socialBuddy/internal/user"
	"time"
)

type mockRepository struct {
	mock.Mock
}

type mockPostService struct {
	mock.Mock
	post.Service
}

type mockUserService struct {
	mock.Mock
	user.Service
}

func (m *mockRepository) GetCom() ([]Comment, error) {
	args := m.Called()
	return args.Get(0).([]Comment), args.Error(1)
}

func (m *mockRepository) GetComByID(idCom int) (*Comment, error) {
	args := m.Called(idCom)
	return args.Get(0).(*Comment), args.Error(1)
}

func (m *mockRepository) GetComByPostID(idPost int) ([]Comment, error) {
	args := m.Called(idPost)
	return args.Get(0).([]Comment), args.Error(1)
}

func (m *mockRepository) GetComByUserID(idUser int) ([]Comment, error) {
	args := m.Called(idUser)
	return args.Get(0).([]Comment), args.Error(1)
}

func (m *mockRepository) GetComByDate(date time.Time) ([]Comment, error) {
	args := m.Called(date)
	return args.Get(0).([]Comment), args.Error(1)
}

func (m *mockRepository) EditCom(com Comment, idCom int) (*Comment, error) {
	args := m.Called(com, idCom)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Comment), args.Error(1)
}

func (m *mockRepository) DeleteCom(idCom int) error {
	args := m.Called(idCom)
	return args.Error(0)
}

func (m *mockUserService) GetUserByID(idUser int) (*user.User, error) {
	args := m.Called(idUser)
	return args.Get(0).(*user.User), args.Error(1)
}

func (m *mockPostService) GetPostByID(idPost int) (*post.Post, error) {
	args := m.Called(idPost)
	return args.Get(0).(*post.Post), args.Error(1)
}

func (m *mockRepository) CreateCom(com Comment) (*Comment, error) {
	args := m.Called(com)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Comment), args.Error(1)
}

var _ = Describe("The Service Test", func() {
	var (
		mockComRepository *mockRepository
		mockServicePost   *mockPostService
		mockServiceUser   *mockUserService
	)
	BeforeEach(func() {
		mockComRepository = new(mockRepository)
		mockServicePost = new(mockPostService)
		mockServiceUser = new(mockUserService)
	})
	It("should CreateCom successfully", func() {
		customDate := time.Now().In(time.Local)
		mockComRepository.On("CreateCom", mock.AnythingOfType("Comment")).Return(&Comment{
			ID:          1,
			IDPost:      2,
			IDUser:      1,
			DateComment: customDate,
			Content:     "content1",
		}, nil)

		mockServicePost.On("GetPostByID", 2).Return(&post.Post{
			ID:      2,
			IDUser:  1,
			Date:    customDate,
			Title:   "title1",
			Content: "content1",
		}, nil)

		mockServiceUser.On("GetUserByID", 1).Return(&user.User{
			ID:             1,
			Name:           "Name First",
			Age:            35,
			DocumentNumber: "123.345.567-89",
			Email:          "name.first@gmail.com",
			Phone:          "+55 11 92345 6789",
			Address: user.Address{
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

		newService := NewService(mockComRepository, mockServicePost, mockServiceUser)
		comment, err := newService.CreateCom(Comment{
			ID:          1,
			IDPost:      2,
			IDUser:      1,
			DateComment: customDate,
			Content:     "content1",
		})
		Expect(err).ShouldNot(HaveOccurred())
		Expect(comment.ID).Should(Equal(1))
		Expect(comment.IDPost).Should(Equal(2))
	})
	It("should CreateCom unsuccessfully", func() {
		mockComRepository.On("CreateCom", mock.AnythingOfType("Comment")).Return(nil, errors.New("error while CreateCom()"))
		mockServicePost.On("GetPostByID", 2).Return(&post.Post{}, nil)
		mockServiceUser.On("GetUserByID", 1).Return(&user.User{}, nil)
		customDate := time.Now().In(time.Local)
		newService := NewService(mockComRepository, mockServicePost, mockServiceUser)
		comment, err := newService.CreateCom(Comment{
			ID:          1,
			IDPost:      2,
			IDUser:      1,
			DateComment: customDate,
			Content:     "content1",
		})
		Expect(err).Should(HaveOccurred())
		Expect(comment).Should(BeNil())
	})
	It("should GetCom successfully", func() {
		timeNow := time.Date(2023, 11, 13, 0, 0, 0, 0, time.Local)
		mockComRepository.On("GetCom").Return([]Comment{
			{ID: 1,
				IDPost:      2,
				IDUser:      1,
				DateComment: timeNow,
				Content:     "content1",
			},
		}, nil)
		newService := NewService(mockComRepository, nil, nil)
		comments, err := newService.GetCom()
		Expect(err).ShouldNot(HaveOccurred())
		Expect(comments[0].ID).Should(Equal(1))
		Expect(comments[0].IDPost).Should(Equal(2))
	})
	It("should GetCom unsuccessfully", func() {
		mockComRepository.On("GetCom").Return([]Comment{}, errors.New("error while GetCom()"))
		newService := NewService(mockComRepository, nil, nil)
		comments, err := newService.GetCom()
		Expect(err).Should(HaveOccurred())
		Expect(len(comments)).Should(Equal(0))
	})
	It("should GetComByID successfully", func() {
		timeNow := time.Date(2023, 11, 13, 0, 0, 0, 0, time.Local)
		mockComRepository.On("GetComByID", 1).Return(&Comment{
			ID:          1,
			IDPost:      2,
			IDUser:      1,
			DateComment: timeNow,
			Content:     "content1",
		}, nil)
		newService := NewService(mockComRepository, nil, nil)
		comment, err := newService.GetComByID(1)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(comment.ID).Should(Equal(1))
		Expect(comment.IDPost).Should(Equal(2))
	})
	It("should GetComByID unsuccessfully", func() {
		mockComRepository.On("GetComByID", 2).Return(&Comment{}, errors.New("error while GetComByID()"))
		newService := NewService(mockComRepository, nil, nil)
		_, err := newService.GetComByID(2)
		Expect(err).Should(HaveOccurred())
	})
	It("should GetComByPostID successfully", func() {
		timeNow := time.Date(2023, 11, 13, 0, 0, 0, 0, time.Local)
		mockComRepository.On("GetComByPostID", 2).Return([]Comment{
			{ID: 1,
				IDPost:      2,
				IDUser:      1,
				DateComment: timeNow,
				Content:     "content1",
			},
		}, nil)
		newService := NewService(mockComRepository, nil, nil)
		comments, err := newService.GetComByPostID(2)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(comments[0].ID).Should(Equal(1))
		Expect(comments[0].IDUser).Should(Equal(1))
	})
	It("should GetComByPostID unsuccessfully", func() {
		mockComRepository.On("GetComByPostID", 3).Return([]Comment{}, errors.New("error while GetComByPostID()"))
		newService := NewService(mockComRepository, nil, nil)
		_, err := newService.GetComByPostID(3)
		Expect(err).Should(HaveOccurred())
	})
	It("should GetComByUserID successfully", func() {
		timeNow := time.Date(2023, 11, 13, 0, 0, 0, 0, time.Local)
		mockComRepository.On("GetComByUserID", 1).Return([]Comment{
			{ID: 1,
				IDPost:      2,
				IDUser:      1,
				DateComment: timeNow,
				Content:     "content1",
			},
		}, nil)
		newService := NewService(mockComRepository, nil, nil)
		comments, err := newService.GetComByUserID(1)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(comments[0].ID).Should(Equal(1))
		Expect(comments[0].IDPost).Should(Equal(2))
	})
	It("should GetComByUserID unsuccessfully", func() {
		mockComRepository.On("GetComByUserID", 2).Return([]Comment{}, errors.New("error while GetComByUserID()"))
		newService := NewService(mockComRepository, nil, nil)
		_, err := newService.GetComByUserID(2)
		Expect(err).Should(HaveOccurred())
	})
	It("should GetComByDate successfully", func() {
		timeNow := time.Date(2023, 11, 13, 0, 0, 0, 0, time.Local)
		mockComRepository.On("GetComByDate", timeNow).Return([]Comment{
			{ID: 1,
				IDPost:      2,
				IDUser:      1,
				DateComment: timeNow,
				Content:     "content1",
			},
		}, nil)
		newService := NewService(mockComRepository, nil, nil)
		comments, err := newService.GetComByDate(timeNow)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(comments[0].ID).Should(Equal(1))
		Expect(comments[0].IDPost).Should(Equal(2))
	})
	It("should GetComByDate unsuccessfully", func() {
		timeNow := time.Date(2023, 11, 13, 0, 0, 0, 0, time.Local)
		mockComRepository.On("GetComByDate", timeNow).Return([]Comment{}, errors.New("error while GetComByDate()"))
		newService := NewService(mockComRepository, nil, nil)
		_, err := newService.GetComByDate(timeNow)
		Expect(err).Should(HaveOccurred())
	})
	It("should EditCom successfully", func() {
		timeNow := time.Date(2023, 11, 13, 0, 0, 0, 0, time.Local)
		mockComRepository.On("EditCom", mock.AnythingOfType("Comment"), 1).Return(&Comment{
			ID:          1,
			IDPost:      2,
			IDUser:      1,
			DateComment: timeNow,
			Content:     "content1",
		}, nil)
		newService := NewService(mockComRepository, nil, nil)
		comment, err := newService.EditCom(Comment{
			ID:          1,
			IDPost:      2,
			IDUser:      1,
			DateComment: timeNow,
			Content:     "content1",
		}, 1)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(comment.ID).Should(Equal(1))
		Expect(comment.IDPost).Should(Equal(2))
	})
	It("should EditCom unsuccessfully", func() {
		timeNow := time.Date(2023, 11, 13, 0, 0, 0, 0, time.Local)
		mockComRepository.On("EditCom", mock.AnythingOfType("Comment"), 2).Return(&Comment{}, errors.New("error while EditCom()"))
		newService := NewService(mockComRepository, nil, nil)
		comment, err := newService.EditCom(Comment{
			ID:          1,
			IDPost:      2,
			IDUser:      1,
			DateComment: timeNow,
			Content:     "content1",
		}, 2)
		Expect(err).Should(HaveOccurred())
		Expect(comment).Should(BeNil())
	})
	It("should DeleteCom successfully", func() {
		mockComRepository.On("DeleteCom", 1).Return(nil)
		newService := NewService(mockComRepository, nil, nil)
		err := newService.DeleteCom(1)
		Expect(err).ShouldNot(HaveOccurred())
	})
	It("should DeleteCom unsuccessfully", func() {
		mockComRepository.On("DeleteCom", 1).Return(errors.New("error while DeleteCom()"))
		newService := NewService(mockComRepository, nil, nil)
		err := newService.DeleteCom(1)
		Expect(err).Should(HaveOccurred())
	})
})
