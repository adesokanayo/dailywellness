package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/adesokanayo/innovation/controller"

	router "github.com/adesokanayo/innovation/http"
) //

var (
	postController controller.PostController = controller.NewPostController()
	httpRouter     router.Router             = router.NewMuxRouter()
)

func main() {
	startApp()
}

func startApp() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}
	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Println("Landing Page  loaded ")

	})

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPosts)

	httpRouter.SERVE(":" + port)
}
