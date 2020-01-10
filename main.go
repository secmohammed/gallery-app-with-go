package main

import (
	"fmt"

	"lenslocked.com/routes"
)

func main() {
	routes.RegisterRoutes()
	fmt.Println("Hello, It's already served")
}
