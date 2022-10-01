package main

import (
	"test/internal/app"
	"test/internal/mux"
)

func main() {
	_ = app.Instance("Successful")
	mux.Init()
}
