package post

import (
	"errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"socialBuddy/internal/user"
	"time"
)

type mockRepository struct {
	mock.Mock
}
type mockUserService struct {
	mock.Mock
	user.Service
}

func (m *mockUserService) GetUserByID(idUser int) (*user.User, error) {
	args := m.Called(idUser)
	return args.Get(0).(*user.User), args.Error(1)
}

func (m *mockRepository) CreatePost(post Post) (*Post, error) {
	args := m.Called(post)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Post), args.Error(1)
}

func (m *mockRepository) GetPosts() ([]Post, error) {
	args := m.Called()
	return args.Get(0).([]Post), args.Error(1)
}

func (m *mockRepository) GetPostByID(idPost int) (*Post, error) {
	args := m.Called(idPost)
	return args.Get(0).(*Post), args.Error(1)
}

func (m *mockRepository) GetPostByUserID(idUser int) ([]Post, error) {
	args := m.Called(idUser)
	return args.Get(0).([]Post), args.Error(1)
}

func (m *mockRepository) GetPostByDate(date time.Time) ([]Post, error) {
	args := m.Called(date)
	return args.Get(0).([]Post), args.Error(1)
}

func (m *mockRepository) GetPostByTitle(title string) ([]Post, error) {
	args := m.Called(title)
	return args.Get(0).([]Post), args.Error(1)
}

func (m *mockRepository) EditPost(post Post, idPost int) (*Post, error) {
	args := m.Called(post, idPost)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Post), args.Error(1)
}

func (m *mockRepository) DeletePost(idPost int) error {
	args := m.Called(idPost)
	return args.Error(0)
}

var _ = Describe("The Service Test", func() {
	var (
		mockPostRepository *mockRepository
		mockService        *mockUserService
	)
	BeforeEach(func() {
		mockPostRepository = new(mockRepository)
		mockService = new(mockUserService)
	})
	It("should CreatePost successfully", func() {
		customDate := time.Now().In(time.Local)
		mockPostRepository.On("CreatePost", mock.AnythingOfType("Post")).Return(&Post{
			ID:      1,
			IDUser:  2,
			Date:    customDate,
			Title:   "title1",
			Content: "content1",
		}, nil)

		mockService.On("GetUserByID", 2).Return(&user.User{
			ID:             2,
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
		newService := NewService(mockPostRepository, mockService)
		post, err := newService.CreatePost(Post{
			ID:     1,
			IDUser: 2,
			//Date:    customDate,
			Title:   "title1",
			Content: "content1",
		})
		Expect(err).ShouldNot(HaveOccurred())
		Expect(post.ID).Should(Equal(1))
		Expect(post.Title).Should(Equal("title1"))
	})
	It("should CreatePost unsuccessfully", func() {
		//customDate := time.Now().In(time.Local)
		mockPostRepository.On("CreatePost", mock.AnythingOfType("Post")).Return(nil, errors.New("error while CreatePost()"))
		mockService.On("GetUserByID", 2).Return(&user.User{}, nil)
		newService := NewService(mockPostRepository, mockService)
		post, err := newService.CreatePost(Post{
			ID:     1,
			IDUser: 2,
			//Date:    customDate,
			Title:   "title1",
			Content: "content1",
		})
		Expect(err).Should(HaveOccurred())
		Expect(post).Should(BeNil())
	})
	It("should GetPosts successfully", func() {
		timeNow := time.Date(2023, 11, 13, 0, 0, 0, 0, time.Local)
		mockPostRepository.On("GetPosts").Return([]Post{
			{ID: 1,
				IDUser:  2,
				Date:    timeNow,
				Title:   "title1",
				Content: "content1",
			},
		}, nil)
		newService := NewService(mockPostRepository, nil)
		posts, err := newService.GetPosts()
		Expect(err).ShouldNot(HaveOccurred())
		Expect(posts[0].ID).Should(Equal(1))
		Expect(posts[0].Title).Should(Equal("title1"))
	})
	It("should GetPosts unsuccessfully", func() {
		mockPostRepository.On("GetPosts").Return([]Post{}, errors.New("error while GetPosts()"))
		newService := NewService(mockPostRepository, nil)
		posts, err := newService.GetPosts()
		Expect(err).Should(HaveOccurred())
		Expect(len(posts)).Should(Equal(0))
	})
	It("should GetPostByID successfully", func() {
		timeNow := time.Date(2023, 11, 13, 0, 0, 0, 0, time.Local)
		mockPostRepository.On("GetPostByID", 1).Return(&Post{
			ID:      1,
			IDUser:  2,
			Date:    timeNow,
			Title:   "title1",
			Content: "content1",
		}, nil)
		newService := NewService(mockPostRepository, nil)
		post, err := newService.GetPostByID(1)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(post.ID).Should(Equal(1))
		Expect(post.Title).Should(Equal("title1"))
	})
	It("should GetPostByID unsuccessfully", func() {
		mockPostRepository.On("GetPostByID", 2).Return(&Post{}, errors.New("error while GetPostByID()"))
		newService := NewService(mockPostRepository, nil)
		_, err := newService.GetPostByID(2)
		Expect(err).Should(HaveOccurred())
	})
	It("should GetPostByUserID successfully", func() {
		timeNow := time.Date(2023, 11, 13, 0, 0, 0, 0, time.Local)
		mockPostRepository.On("GetPostByUserID", 2).Return([]Post{
			{ID: 1,
				IDUser:  2,
				Date:    timeNow,
				Title:   "title1",
				Content: "content1",
			},
		}, nil)
		newService := NewService(mockPostRepository, nil)
		posts, err := newService.GetPostByUserID(2)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(posts[0].ID).Should(Equal(1))
		Expect(posts[0].Title).Should(Equal("title1"))
	})
	It("should GetPostByUserID unsuccessfully", func() {
		mockPostRepository.On("GetPostByUserID", 1).Return([]Post{}, errors.New("error while GetPostByUserID()"))
		newService := NewService(mockPostRepository, nil)
		posts, err := newService.GetPostByUserID(1)
		Expect(err).Should(HaveOccurred())
		Expect(len(posts)).Should(Equal(0))
	})
	It("should GetPostByDate successfully", func() {
		timeNow := time.Date(2023, 11, 13, 0, 0, 0, 0, time.Local)
		mockPostRepository.On("GetPostByDate", timeNow).Return([]Post{
			{ID: 1,
				IDUser:  2,
				Date:    timeNow,
				Title:   "title1",
				Content: "content1",
			},
		}, nil)
		newService := NewService(mockPostRepository, nil)
		posts, err := newService.GetPostByDate(timeNow)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(posts[0].ID).Should(Equal(1))
		Expect(posts[0].Title).Should(Equal("title1"))
	})
	It("should GetPostByDate unsuccessfully", func() {
		timeNow := time.Date(2023, 11, 13, 0, 0, 0, 0, time.Local)
		mockPostRepository.On("GetPostByDate", timeNow).Return([]Post{}, errors.New("error while GetPostByDate()"))
		newService := NewService(mockPostRepository, nil)
		posts, err := newService.GetPostByDate(timeNow)
		Expect(err).Should(HaveOccurred())
		Expect(len(posts)).Should(Equal(0))
	})
	It("should GetPostByTitle successfully", func() {
		timeNow := time.Date(2023, 11, 13, 0, 0, 0, 0, time.Local)
		mockPostRepository.On("GetPostByTitle", "title1").Return([]Post{
			{ID: 1,
				IDUser:  2,
				Date:    timeNow,
				Title:   "title1",
				Content: "content1",
			},
		}, nil)
		newService := NewService(mockPostRepository, nil)
		posts, err := newService.GetPostByTitle("title1")
		Expect(err).ShouldNot(HaveOccurred())
		Expect(posts[0].ID).Should(Equal(1))
		Expect(posts[0].Title).Should(Equal("title1"))
	})
	It("should GetPostByTitle unsuccessfully", func() {
		mockPostRepository.On("GetPostByTitle", "title1").Return([]Post{}, errors.New("error while GetPostByTitle()"))
		newService := NewService(mockPostRepository, nil)
		posts, err := newService.GetPostByTitle("title1")
		Expect(err).Should(HaveOccurred())
		Expect(len(posts)).Should(Equal(0))
	})
	It("should EditPost successfully", func() {
		customDate := time.Now().In(time.Local)
		mockPostRepository.On("EditPost", mock.AnythingOfType("Post"), 1).Return(&Post{
			ID:      1,
			IDUser:  2,
			Date:    customDate,
			Title:   "title1",
			Content: "content1",
		}, nil)
		newService := NewService(mockPostRepository, nil)
		post, err := newService.EditPost(Post{
			ID:     1,
			IDUser: 2,
			//Date:    customDate,
			Title:   "title1",
			Content: "content1",
		}, 1)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(post.ID).Should(Equal(1))
		Expect(post.Title).Should(Equal("title1"))
	})
	It("should EditPost unsuccessfully", func() {
		mockPostRepository.On("EditPost", mock.AnythingOfType("Post"), 2).Return(nil, errors.New("error while EditPost()"))
		newService := NewService(mockPostRepository, nil)
		post, err := newService.EditPost(Post{
			ID:     1,
			IDUser: 2,
			//Date:    customDate,
			Title:   "title1",
			Content: "content1",
		}, 2)
		Expect(err).Should(HaveOccurred())
		Expect(post).Should(BeNil())
	})
	It("should DeletePost successfully", func() {
		mockPostRepository.On("DeletePost", 1).Return(nil)
		newService := NewService(mockPostRepository, nil)
		err := newService.DeletePost(1)
		Expect(err).ShouldNot(HaveOccurred())
	})
	It("should DeletePost unsuccessfully", func() {
		mockPostRepository.On("DeletePost", 1).Return(errors.New("error while DeletePost()"))
		newService := NewService(mockPostRepository, nil)
		err := newService.DeletePost(1)
		Expect(err).Should(HaveOccurred())
	})
})
