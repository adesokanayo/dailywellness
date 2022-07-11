package controller

import (
	"encoding/json"
	"net/http"

	"github.com/adesokanayo/dailywellness/entity"
	"github.com/adesokanayo/dailywellness/errors"
	"github.com/adesokanayo/dailywellness/services"
)

type controller struct{}

//NewPostController creates a controller instance
func NewPostController() TipsManager {
	return &controller{}
}

var (
	postService = services.NewTipService()
)

//TipsManager implements all the operations we need on a post.
type TipsManager interface {
	GetTips(response http.ResponseWriter, request *http.Request)
	AddTips(response http.ResponseWriter, request *http.Request)
	GetTip(response http.ResponseWriter, request *http.Request)
	GetDailyTip(response http.ResponseWriter, request *http.Request)
	GetRandomTip(response http.ResponseWriter, request *http.Request)
}

//GetTips retrieve all posts
func (c *controller) GetTips(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")

	posts, err := postService.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error getting the post"})
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)

}

func (c *controller) GetDailyTip(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")

	posts, err := postService.FindToday()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error getting the post"})
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)

}

func (c *controller) GetTip(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")

	posts, err := postService.FindToday()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error getting the post"})
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)

}

func (c *controller) GetRandomTip(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")

	posts, err := postService.FindOne()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error getting the post"})
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)

}

func (c *controller) AddTips(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	var post entity.Tip
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
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(result)
}
