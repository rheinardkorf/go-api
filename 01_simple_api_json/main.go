package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	ID    int    `json:"id"`
}

func main() {
	fmt.Println("Starting JSON server...")
	http.HandleFunc("/api", userRouter)
	http.ListenAndServe(":8080", nil)

}
func userRouter(writer http.ResponseWriter, request *http.Request) {
	theUser := User{
		"Fictional Person",
		"fp@example.com",
		100,
	}

	output, _ := json.Marshal(&theUser)
	fmt.Fprintln(writer,string(output))

}
