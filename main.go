package main

import (
	"yo-dian-service/routers"
)

func main() {
	r := routers.Router()
	r.Run()
}
