package main

import (
	"fmt"
	"net/http"
	"encoding/csv"
	"bytes"
	"strconv"
)

type User struct {
	Name  string `csv:"name"`
	Email string `csv:"email"`
	ID    int    `csv:"id"`
}

func main() {
	fmt.Println("Starting CSV server...")
	http.HandleFunc("/api", userRouter)
	http.ListenAndServe(":8080", nil)

}
func userRouter(writer http.ResponseWriter, request *http.Request) {
	theUser := User{
		"Fictional Person",
		"fp@example.com",
		100,
	}

	b := &bytes.Buffer{}
	w := csv.NewWriter(b)
	w.Write([]string{"name","email","id"})
	w.Write([]string{theUser.Name,theUser.Email,strconv.Itoa(theUser.ID)})
	w.Flush()
	output := string(b.String())
	fmt.Fprintln(writer, string(output))

}
