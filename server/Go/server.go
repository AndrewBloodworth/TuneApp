package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
    }
    fmt.Fprintf(w, "POST request successful")
    name := r.FormValue("name")
    address := r.FormValue("address")
    fmt.Fprintf(w, "Name = %s\n", name)
    fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/hello" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
    }


    fmt.Fprintf(w, "Hello!")
		type User struct {
			Id    int     `json:"id"`
			Name  string  `json:"name"`
			Email string  `json:"email"`
			Phone string  `json:"phone"`
 }

		w.Header().Set("Content-Type", "application/json")
      user := User {
                    Id: 1,
                    Name: "John Doe",
                    Email: "johndoe@gmail.com",
                    Phone: "000099999",
               }
     json.NewEncoder(w).Encode(user)
}


func main() {
    fileServer := http.FileServer(http.Dir("./static"))
    http.Handle("/", fileServer)
    http.HandleFunc("/form", formHandler)
    http.HandleFunc("/hello", helloHandler)


    fmt.Printf("Starting server at port 8080\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
