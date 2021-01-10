package services

import (
	"errors"
	"testing"

	"github.com/adesokanayo/dailywellness/entity"
	"github.com/stretchr/testify/assert"
)

var (
	mockPostingService mockPosting

	getPost      func(entity.Post) (*entity.Post, error)
	validatePost func(entity.Post) error
	findAllPost   func() ([]entity.Post, error)
	createPost  func(*entity.Post)(*entity.Post, error)
)

type mockPosting struct{}

func (s *mockPosting) Validate(post *entity.Post) error {
	singlepost := entity.Post{ID: 1, Text: "Good Post"}
	err := validatePost(singlepost)
	return err
}

func (s *mockPosting) Create(post *entity.Post) (*entity.Post, error) {

	return createPost(post)
}

func (s *mockPosting) FindAll() ([]entity.Post, error) {
	return findAllPost()
}

type mockPostingInterface interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll([]*entity.Post) ([]entity.Post, error)
}

func TestValidateEmptyTitle(t *testing.T) {
	//arrange
	emptyTitlepost := entity.Post{ID: 1}

	validatePost =  func(entity.Post) error{
		return errors.New("empty title")
	}
	//act
	err := mockPostingService.Validate(&emptyTitlepost)

	//assert
	assert.NotNil(t, err)
	assert.Equal(t, errors.New("empty title"), err, "Title is not expected to be empty")

}

func TestValidateEmptyPost(t *testing.T) {
	//arrange
	validatePost =  func(entity.Post) error{
		return errors.New("empty post")
	}
	//act
	err1 := mockPostingService.Validate(nil)

	//assert
	assert.NotNil(t, errors.New("empty post"), err1, "empty post")
}

func TestValidate_GoodPost(t *testing.T) {
	//arrange
	post := entity.Post{ID: 1, Title: "Title", Text: "Good Post"}
	validatePost =  func(entity.Post) error{
		return nil
	}
	//act
	err := mockPostingService.Validate(&post)

	//assert
	assert.Nil(t, err)
}

func TestCreatePost(t *testing.T) {
	//arrange
	post := entity.Post{ID: 1, Title: "Title", Text: "Good Post"}

	createPost = func (post *entity.Post) (*entity.Post, error){
		return post, nil

	} 
	//act
	newPost, err := mockPostingService.Create(&post)

	//assert
	assert.Nil(t, err)
	assert.Equal(t, newPost, &post, "Post is expected to be returned")
}

func TestFindAll(t *testing.T) {
	var posts  []entity.Post

	post1 := entity.Post{ID: 1, Title: "Title", Text: "Good Post"}
	post2 := entity.Post{ID: 2, Title: "Title2", Text: "Good Post2"}
	posts = append(posts, post1, post2)

	
	findAllPost =   func() ([]entity.Post, error){
		return posts, nil
	}
	newPosts, err := mockPostingService.FindAll()
	assert.NotNil(t, newPosts)
	assert.Nil(t, err)
	assert.Equal(t, newPosts, posts)

}
