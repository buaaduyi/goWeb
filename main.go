package main

import "goweb/controller"

func main() {

	ctrl := new(controller.Controller)
	ctrl.HTTPAddr = "localhost:8080"
	ctrl.RedisAddr = "localhost:6379"
	ctrl.Init()
}
