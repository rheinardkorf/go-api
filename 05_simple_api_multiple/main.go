package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"encoding/xml"
	"gopkg.in/yaml.v2"
	"bytes"
	"encoding/csv"
	"strconv"
)

type User struct {
	Name  string `json:"name" xml:"name" yaml:"name"`
	Email string `json:"email" xml:"email" yaml:"email"`
	ID    int    `json:"id" xml:"id" yaml:"id"`
}

func main() {
	fmt.Println("Starting server...")
	http.HandleFunc("/api", userRouter)
	http.HandleFunc("/api/xml", userRouter)
	http.HandleFunc("/api/yaml", userRouter)
	http.HandleFunc("/api/csv", userRouter)
	http.ListenAndServe(":8080", nil)

}
func userRouter(writer http.ResponseWriter, request *http.Request) {
	theUser := User{
		"Fictional Person",
		"fp@example.com",
		100,
	}

	var output string

	switch request.URL.String() {
	case "/api":
		out, _ := json.Marshal(&theUser)
		output = string(out)
	case "/api/xml":
		out, _ := xml.Marshal(&theUser)
		output = string(out)
	case "/api/yaml":
		out, _ := yaml.Marshal(&theUser)
		output = string(out)
	case "/api/csv":
		b := &bytes.Buffer{}
		w := csv.NewWriter(b)
		w.Write([]string{"name","email","id"})
		w.Write([]string{theUser.Name,theUser.Email,strconv.Itoa(theUser.ID)})
		w.Flush()
		output = b.String()
	}

	fmt.Fprintln(writer, string(output))
}
