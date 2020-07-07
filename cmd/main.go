package main

import (
	demoapi "github.com/lromais/rest-api-golang"
)

func main() {
	a := demoapi.App{}
	a.Initialize()

	a.Run(":8080")
}
