package controller

import (
	"encoding/json"
	"net/http"

	"github.com/adesokanayo/innovation/entity"
	"github.com/adesokanayo/innovation/errors"
	"github.com/adesokanayo/innovation/services"
)

type controller struct{}

//NewPostController creates a controller instance
func NewPostController() PostController {
	return &controller{}
}

var (
	postService services.PostService = services.NewPostService()
)

//PostController is
type PostController interface {
	GetPosts(response http.ResponseWriter, request *http.Request)
	AddPosts(response http.ResponseWriter, request *http.Request)
}

//GetPosts retrieve all posts
func (*controller) GetPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")

	posts, err := postService.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{"Error getting the post"})
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)

}

func (*controller) AddPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error adding the posts"})
		return
	}

	err1 := postService.Validate(&post)
	if err1 != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}
	result, err2 := postService.Create(&post)
	if err2 != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: err2.Error()})
		return
	}
	postService.Create(&post)
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(result)
}
