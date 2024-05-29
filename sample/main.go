package main

import (
	"fmt"
	"net/http"

	"github.com/MateoCaicedoW/file_system/manager"
)

func main() {
	s := http.NewServeMux()

	fs := manager.New(
		manager.WithPrefix("/files/"),
	)

	s.Handle("/", fs)

	fmt.Println("Server running on port 8080")
	err := http.ListenAndServe(":8080", s)
	if err != nil {
		fmt.Println(err)
	}
}