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
		port = "3000"
	}
	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Println("Landing Page  loaded ")

	})

	httpRouter.GET("/tips", postController.GetTips)
	httpRouter.GET("/dailytip", postController.GetDailyTip)
	httpRouter.POST("/tips", postController.AddTips)
	httpRouter.GET("/randomtip", postController.GetRandomTip)

	httpRouter.SERVE(":" + port)
}
