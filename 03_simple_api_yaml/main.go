package main

import (
	"fmt"
	"net/http"
	"gopkg.in/yaml.v2"
)

type User struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email"`
	ID    int    `yaml:"id"`
}

func main() {
	fmt.Println("Starting YAML server...")
	http.HandleFunc("/api", userRouter)
	http.ListenAndServe(":8080", nil)

}
func userRouter(writer http.ResponseWriter, request *http.Request) {
	theUser := User{
		"Fictional Person",
		"fp@example.com",
		100,
	}

	output, _ := yaml.Marshal(&theUser)
	fmt.Fprintln(writer, string(output))

}
