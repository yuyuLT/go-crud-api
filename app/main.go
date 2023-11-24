package main

import (
	"api/routes"
)

func main() {
	router := routes.NewRoutes()
	router.Run()
}
