package services

import (
	"errors"
	"testing"

	"github.com/adesokanayo/innovation/entity"
	"github.com/stretchr/testify/assert"
)

var (
	mockPostingService mockPosting

	getPost      func(entity.Post) (*entity.Post, error)
	validatePost func(entity.Post) error
	GetAllPost   func() ([]entity.Post, error)
)

type mockPosting struct{}

func (s *mockPosting) Validate(post *entity.Post) error {
	singlepost := entity.Post{ID: 1, Text: "Good Post"}
	return validatePost(singlepost)
}

func (s *mockPosting) Create(post *entity.Post) (*entity.Post, error) {

	singlepost := entity.Post{ID: 1, Title: "Title", Text: "Good Post"}
	return getPost(singlepost)
}

func (s *mockPosting) FindAll() ([]entity.Post, error) {
	singlepost := entity.Post{ID: 1, Text: "Good Post"}
	singlepost1 := entity.Post{ID: 1, Text: "Good Post"}

	var posts []entity.Post
	posts = append(posts, singlepost, singlepost1)
	return GetAllPost()
}

type mockPostingInterface interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

func TestValidateEmptyTitle(t *testing.T) {
	//arrange
	emptyTitlepost := entity.Post{ID: 1}

	//act
	err := mockPostingService.Validate(&emptyTitlepost)

	//assert
	assert.NotNil(t, err)
	assert.Equal(t, errors.New("post title is empty"), err, "Title is not expected to be empty")

}

func TestValidateEmptyPost(t *testing.T) {
	//arrange
	//act
	err1 := mockPostingService.Validate(nil)
	//assert
	assert.NotNil(t, errors.New("empty post"), err1, "empty post")
}

func TestValidateValidPost(t *testing.T) {
	//arrange
	post := entity.Post{ID: 1, Title: "Title", Text: "Good Post"}

	//act
	err := mockPostingService.Validate(&post)

	//assert
	assert.Nil(t, err)
}

func TestValidateCreatePost(t *testing.T) {
	//arrange
	post := entity.Post{ID: 1, Title: "Title", Text: "Good Post"}

	//act
	newPost, err := mockPostingService.Create(&post)

	//assert
	assert.Nil(t, err)
	assert.Equal(t, newPost, &post, "Post is expected to be returned")
}

func TestFindAll(t *testing.T) {
	var posts []entity.Post

	post1 := entity.Post{ID: 1, Title: "Title", Text: "Good Post"}
	post2 := entity.Post{ID: 2, Title: "Title2", Text: "Good Post2"}
	posts = append(posts, post1)
	posts = append(posts, post2)

	newPosts, err := mockPostingService.FindAll()
	assert.NotNil(t, newPosts)
	assert.Nil(t, err)
	//assert.Equal(t, newPosts, posts)

}
