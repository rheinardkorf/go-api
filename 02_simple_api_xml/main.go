package main

import (
	"fmt"
	"net/http"
	"encoding/xml"
)

type User struct {
	Name  string `xml:"name"`
	Email string `xml:"email"`
	ID    int    `xml:"id"`
}

func main() {
	fmt.Println("Starting XML server...")
	http.HandleFunc("/api", userRouter)
	http.ListenAndServe(":8080", nil)

}
func userRouter(writer http.ResponseWriter, request *http.Request) {
	theUser := User{
		"Fictional Person",
		"fp@example.com",
		100,
	}

	output, _ := xml.Marshal(&theUser)
	fmt.Fprintln(writer, string(output))

}
