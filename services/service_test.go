package services

import (
	"errors"
	"testing"

	"github.com/adesokanayo/dailywellness/entity"
	"github.com/stretchr/testify/assert"
)

var (
	mockPostingService mockPosting

	getPost      func(entity.Tip) (*entity.Tip, error)
	validatePost func(entity.Tip) error
	findAllPost   func() ([]entity.Tip, error)
	createPost  func(*entity.Tip)(*entity.Tip, error)
)

type mockPosting struct{}

func (s *mockPosting) Validate(post *entity.Tip) error {
	singlepost := entity.Tip{ID: 1, Text: "Good Post"}
	err := validatePost(singlepost)
	return err
}

func (s *mockPosting) Create(post *entity.Tip) (*entity.Tip, error) {

	return createPost(post)
}

func (s *mockPosting) FindAll() ([]entity.Tip, error) {
	return findAllPost()
}

type mockPostingInterface interface {
	Validate(post *entity.Tip) error
	Create(post *entity.Tip) (*entity.Tip, error)
	FindAll([]*entity.Tip) ([]entity.Tip, error)
}

func TestValidateEmptyTitle(t *testing.T) {
	//arrange
	emptyTitlepost := entity.Tip{ID: 1}

	validatePost =  func(entity.Tip) error{
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
	validatePost =  func(entity.Tip) error{
		return errors.New("empty post")
	}
	//act
	err1 := mockPostingService.Validate(nil)

	//assert
	assert.NotNil(t, errors.New("empty post"), err1, "empty post")
}

func TestValidate_GoodPost(t *testing.T) {
	//arrange
	post := entity.Tip{ID: 1, Title: "Title", Text: "Good Post"}
	validatePost =  func(entity.Tip) error{
		return nil
	}
	//act
	err := mockPostingService.Validate(&post)

	//assert
	assert.Nil(t, err)
}

func TestCreatePost(t *testing.T) {
	//arrange
	post := entity.Tip{ID: 1, Title: "Title", Text: "Good Post"}

	createPost = func (post *entity.Tip) (*entity.Tip, error){
		return post, nil

	} 
	//act
	newPost, err := mockPostingService.Create(&post)

	//assert
	assert.Nil(t, err)
	assert.Equal(t, newPost, &post, "Post is expected to be returned")
}

func TestFindAll(t *testing.T) {
	var posts  []entity.Tip

	post1 := entity.Tip{ID: 1, Title: "Title", Text: "Good Post"}
	post2 := entity.Tip{ID: 2, Title: "Title2", Text: "Good Post2"}
	posts = append(posts, post1, post2)

	
	findAllPost =   func() ([]entity.Tip, error){
		return posts, nil
	}
	newPosts, err := mockPostingService.FindAll()
	assert.NotNil(t, newPosts)
	assert.Nil(t, err)
	assert.Equal(t, newPosts, posts)

}
