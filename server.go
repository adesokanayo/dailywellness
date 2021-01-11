package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/adesokanayo/dailywellness/controller"

	router "github.com/adesokanayo/dailywellness/http"
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

	httpRouter.GET("/tips", postController.GetPosts)
	httpRouter.GET("/dailytip", postController.GetDailyPost)
	httpRouter.POST("/tips", postController.AddPosts)

	httpRouter.SERVE(":" + port)
}
