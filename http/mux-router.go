package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct{}

//NewMuxRouter creates a newmux router
func NewMuxRouter() Router {
	return &muxRouter{}
}

var (
	muxDispatcher = mux.NewRouter()
)

func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {

	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {

	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}
func (*muxRouter) SERVE(port string) {

	fmt.Printf("mux http server running on port %v", port)
	http.ListenAndServe(port, muxDispatcher)
}
