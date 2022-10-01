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
	http.ListenAndServe(":8282", r.router)
}

func (r *mRouter) setHandlers() {
	r.router.HandleFunc("/devices/{id}", handlers.GetDeviceByIdHandler).Methods("GET")
	r.router.HandleFunc("/devices/", handlers.GetDevicesHandler).Methods("GET")
}
