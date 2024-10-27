package main

import (
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

func login(w http.ResponseWriter, r *http.Request) {
	var FileName = "login.html"
	t, err := template.ParseFiles(FileName)

	if err != nil {
		fmt.Println("Error when parsing file", err)
		return
	}

	err = t.ExecuteTemplate(w, FileName, nil)

	if err != nil {
		fmt.Println("Error when executing template")
	}
}

var userDB = map[string]string{
	"allen": "goodPassword",
}

func loginSubmit(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if userDB[username] == password {
		w.WriteHeader(http.StatusOK)
		//fmt.Fprintf(w, "Login Successful")
		tmpl, err := template.New("name").Parse(`

		`)
		// Error checking elided
		err = tmpl.Execute(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
		//fmt.Fprintf(w, "Error")
	}

}

func dashboard(w http.ResponseWriter, r *http.Request) {

}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/login": //Handles login page
		if r.Method == http.MethodGet {
			login(w, r)
		}

	case "/login-submit": //Handles after login
		loginSubmit(w, r)

	case "/dashboard":
		dashboard(w, r)

	default:
		fmt.Fprintf(w, "Test")
	}
}

func main() {
	//fmt.Println("Server is running at http://localhost/login")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
