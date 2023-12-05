package main

import (
	"htmx/model"
	"htmx/routes"
)

func main() {
	// ...
	model.Setup()
	routes.SetupAndRun()
	// ...
}
