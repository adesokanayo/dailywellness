package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/adesokanayo/dailywellness/controller"

	router "github.com/adesokanayo/dailywellness/http"
) //

var (
	postController controller.TipsManager = controller.NewPostController()
	httpRouter     router.Router          = router.NewMuxRouter()
)

func main() {
	startApp()
}

func startApp() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}
	httpRouter.GET("/health", func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte("Health check "))
		resp.WriteHeader(http.StatusOK)
		resp.Write([]byte("Health check "))
		fmt.Println("Landing Page  loaded ")

	})

	httpRouter.GET("/readiness", func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte("readiness check "))
		resp.WriteHeader(http.StatusOK)
		resp.Write([]byte("Readiness check "))
		fmt.Println("Landing Page  loaded ")

	})
	httpRouter.GET("/tips", postController.GetTips)
	httpRouter.GET("/dailytip", postController.GetDailyTip)
	httpRouter.POST("/tips", postController.AddTips)
	httpRouter.GET("/randomtip", postController.GetRandomTip)

	httpRouter.SERVE(":" + port)
}
