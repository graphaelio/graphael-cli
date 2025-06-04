package main

import (
	"fmt"

	graphael "github.com/graphaelio/graphael-go-client"
)

func main() {
	client := graphael.NewClient(graphael.ClientConfig{
		BaseURL: nil, // Use default base URL
		ID:      "your-client-id",
		Secret:  "your-client-secret",
	})

	if client == nil {
		fmt.Println("Failed to create Graphael client")
		return
	}

	fmt.Println("Hello, world!")
}
