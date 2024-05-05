package main

import (
	"fmt"
	"link-shortener/adapter"
	"link-shortener/usecases"
)

func main() {
	// Implement the main function to start the API server
	urlMappingRepo := adapter.NewUrlMappingRepo(nil)
	urlMappingCore := usecases.NewUrlMappingCore(urlMappingRepo)

	fmt.Println(urlMappingCore.CreateUrlMapping("https://www.LinkSpiral.com"))

}
