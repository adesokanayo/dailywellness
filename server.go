package main

import (
	"fmt"
	"net/http"

	"github.com/adesokanayo/innovation/controller"

	router "github.com/adesokanayo/innovation/http"
)

const port string = ":8000"

var (
	postController =  controller.NewPostController()
	httpRouter   = router.NewMuxRouter
)

func main() {
 startApp()
}

func startApp(){
	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Println("up and running...")

	})

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPosts)

	httpRouter.SERVE(port)
}
