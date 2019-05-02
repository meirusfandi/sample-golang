package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Fullname Name   `json:"fullname"`
}

type Name struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type AllUser []User

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Hello world, %q", html.EscapeString(r.URL.Path))

		file := path.Join("view", "index.html")
		var templates, err = template.ParseFiles(file)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var user AllUser

		user = AllUser{
			User{
				ID:       "1",
				Username: "meirusfandi",
				Email:    "mei@gmail.com",
				Fullname: Name{
					FirstName: "Mei",
					LastName:  "Rusfandi",
				},
			},
			User{
				ID:       "2",
				Username: "mrcondong",
				Email:    "condong@gmail.com",
				Fullname: Name{
					FirstName: "Condong",
					LastName:  "Rusfandi",
				},
			},
		}

		err = templates.Execute(w, user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
