package mux

import (
	"github.com/gorilla/mux"
	"net/http"
	"test/internal/handlers"
)

type mRouter struct {
	router *mux.Router
}

func router() *mRouter {
	gR := new(mRouter)
	gR.router = mux.NewRouter()
	gR.setHandlers()
	return gR
}

func Init() {
	r := router()
	http.ListenAndServe(":8181", r.router)
}

func (r *mRouter) setHandlers() {
	r.router.HandleFunc("/update", handlers.UpdatesHandler)
}
