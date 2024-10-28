package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

var (
	key   = []byte("super-secret-key") // Change this in production
	store = sessions.NewCookieStore(key)
)

// Template cache
var templates = template.Must(template.ParseGlob("templates/*.html"))

// Dummy user for authentication
var username = "admin"
var password = "password"

func loginHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	if r.Method == http.MethodGet {
		templates.ExecuteTemplate(w, "login.html", nil)
		return
	}

	// Handle POST request for login
	r.ParseForm()
	user := r.FormValue("username")
	pass := r.FormValue("password")

	if user == username && pass == password {
		session.Values["authenticated"] = true
		session.Save(r, w)
		http.Redirect(w, r, "/dashboard", http.StatusFound)
	} else {
		templates.ExecuteTemplate(w, "login.html", "Invalid credentials")
	}
}

// Welcome handler (protected page)
func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	templates.ExecuteTemplate(w, "dashboard.html", nil)
}

// Logout handler
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	session.Values["authenticated"] = false
	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusFound)
}

// Order Handler
func OrderHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "Order.html", nil)
}

// Submit Handler
func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	type Medication_Request struct {
		File_Number int
		Nurse_Name  string
		Ward        string
	}
	var DB *sql.DB
	DB, err := sql.Open("sqlite3", "./DB.db")
	if err != nil {
		log.Fatal(err)
	}
	if r.Method == "POST" {
		File_Number := r.FormValue("File_Number")
		Nurse_Name := r.FormValue("Nurse_Name")
		Ward := r.FormValue("Ward")

		_, err := DB.Exec("INSERT INTO Medication_Request (FILE_NUMBER,NURSE_NAME,WARD,BED,MEDICATION,UOM,REQUEST_TIME,NURSE_REMARKS,STATUS,PHARMACY_REMARKS) VALUES (?,?,?,?,?,?,?,?,?,?)", File_Number, Nurse_Name, Ward, nil, nil, nil, nil, nil, nil, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	http.Redirect(w, r, "/Order", http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/dashboard", welcomeHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/Order", OrderHandler)
	http.HandleFunc("/Submit", SubmitHandler)
	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
