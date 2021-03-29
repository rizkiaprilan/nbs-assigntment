package main

import (
	"backend-nbs/helpers"
	"backend-nbs/routes"
	"fmt"
)

func main() {
	router := routes.Controller()

	fmt.Println("Start server with port : " + helpers.GetEnv("PORT"))
	router.Start(":" + helpers.GetEnv("PORT"))
}
