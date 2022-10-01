package main

import (
	"test/internal/app"
	"test/internal/mux"
	"test/internal/rabbitmq"
)

func main() {
	_ = app.Instance("Successful")
	rabbitmq.Init()
	mux.Init()
}
