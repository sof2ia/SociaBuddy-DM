package post

import (
	"errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"time"
)

type mockRepository struct {
	mock.Mock
}

func (m *mockRepository) GetPost() ([]Post, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mockRepository) GetPostByID(idPost int) (*Post, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mockRepository) GetPostByUserID(idUser int) ([]Post, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mockRepository) GetPostByDate(date time.Time) ([]Post, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mockRepository) GetPostByTitle(title string) ([]Post, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mockRepository) EditPost(post Post, idPost int) (*Post, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mockRepository) DeletePost(idPost int) error {
	//TODO implement me
	panic("implement me")
}

func (m *mockRepository) DeleteAllPostsByUserID(idUser int) error {
	//TODO implement me
	panic("implement me")
}

func (m *mockRepository) CreatePost(post Post) (*Post, error) {
	args := m.Called(post)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Post), args.Error(1)
}

var _ = Describe("The Service Test", func() {
	var (
		mockPostRepository *mockRepository
	)
	BeforeEach(func() {
		mockPostRepository = new(mockRepository)
	})
	It("should CreatePost successfully", func() {
		format := "Mon Jan _2 15:04:05 2006"
		customDateStr := time.Now().In(time.Local).Format(format)
		mockPostRepository.On("CreatePost", Post{
			ID:      1,
			IDUser:  2,
			Date:    customDateStr,
			Title:   "title1",
			Content: "content1",
		}).Return(&Post{
			ID:      1,
			IDUser:  2,
			Date:    customDateStr,
			Title:   "title1",
			Content: "content1",
		}, nil)
		newService := NewService(mockPostRepository)
		post, err := newService.CreatePost(Post{
			ID:      1,
			IDUser:  2,
			Date:    customDateStr,
			Title:   "title1",
			Content: "content1",
		})
		Expect(err).ShouldNot(HaveOccurred())
		Expect(post.ID).Should(Equal(1))
		Expect(post.Title).Should(Equal("title1"))
	})
	It("should CreatePost unsuccessfully", func() {
		format := "Mon Jan _2 15:04:05 2006"
		customDateStr := time.Now().In(time.Local).Format(format)
		mockPostRepository.On("CreatePost", Post{
			ID:      1,
			IDUser:  2,
			Date:    customDateStr,
			Title:   "title1",
			Content: "content1",
		}).Return(nil, errors.New("error while CreatePost()"))
		newService := NewService(mockPostRepository)
		post, err := newService.CreatePost(Post{
			ID:      1,
			IDUser:  2,
			Date:    customDateStr,
			Title:   "title1",
			Content: "content1",
		})
		Expect(err).Should(HaveOccurred())
		Expect(post).Should(BeNil())
	})
})
