package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	//"github.com/birddevelper/gomologin"

	"github.com/gorilla/sessions"
	_ "github.com/mattn/go-sqlite3" // SQLite driver
	"golang.org/x/crypto/bcrypt"
)

var (
	key   = []byte("super-secret-key") // Change this in production
	store = sessions.NewCookieStore(key)
)

// Template cache
var templates = template.Must(template.ParseGlob("templates/*.html"))
var DB *sql.DB

//Dummy user for authentication
//var username = "admin"
//var password = "password"

type Users struct {
	Username string
	Password string
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	//session, _ := store.Get(r, "session")
	if r.Method == http.MethodGet {
		templates.ExecuteTemplate(w, "login.html", nil)
		return
	}

}
func authenticate(username, password string) (bool, error) {
	var hashedPassword string

	DB, err := sql.Open("sqlite3", "./DB.db")
	if err != nil {
		log.Fatal(err)
	}
	// Query database for user
	query := "SELECT password FROM users WHERE username = ?"
	err = DB.QueryRow(query, username).Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No user found with username: %s", username)
			return false, nil // User does not exist
		}
		log.Printf("Error querying user: %v", err)
		return false, err // Query error
	}

	log.Printf("Retrieved hashed password from DB: %s", hashedPassword)

	// Compare the hashed password with the provided password
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Println("Password mismatch")
		log.Println("username:", username)
		log.Println("hashed:", hashedPassword)
		log.Println("password from database:", password)
		return false, nil // Password is incorrect
	} else {
		log.Println("Password matched")

		return true, nil // Successful authentication
	}

}

// Handle login POST request
func login(w http.ResponseWriter, r *http.Request) {

	// if err := r.ParseForm(); err != nil {
	// 	http.Error(w, "Unable to process form", http.StatusBadRequest)
	// 	return
	// }

	username := r.FormValue("username")
	password := r.FormValue("password")

	// Authenticate user
	authenticated, err := authenticate(username, password)
	if err != nil {
		http.Error(w, "Server error, unable to log in", http.StatusInternalServerError)
		log.Printf("Login error: %v", err)
		return
	}

	if authenticated {
		//fmt.Fprintf(w, "Login successful!")
		http.Redirect(w, r, "/dashboard", http.StatusFound)
	} else {
		fmt.Fprintf(w, "Invalid username or password.")

	}
}

// Dashboard handler
func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	// session, _ := store.Get(r, "session")

	// if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
	// 	http.Error(w, "Forbidden", http.StatusForbidden)
	// 	return
	// }

	templates.ExecuteTemplate(w, "dashboard.html", nil)
}

// Logout handler
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	session.Values["authenticated"] = false
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusFound)
}

// Order Handler
func OrderHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "Order.html", nil)
}

// Submit Handler
func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	type Medication_Orders struct {
		File_Number int
		Nurse_Name  string
		Ward        string
	}
	//var DB *sql.DB
	DB, err := sql.Open("sqlite3", "./DB.db")
	if err != nil {
		log.Fatal(err)
	}
	if r.Method == "POST" {
		File_Number := r.FormValue("File_Number")
		Nurse_Name := r.FormValue("Nurse_Name")
		Ward := r.FormValue("Ward")
		Bed := r.FormValue("Bed")
		Medication := r.FormValue("Medication")
		Status := "PENDING"
		UOM := r.FormValue("UOM")
		Request_time := time.Now()
		Nurse_Remarks := r.FormValue("Nurse_Remarks")

		_, err := DB.Exec("INSERT INTO Medication_Orders (FILE_NUMBER,NURSE_NAME,WARD,BED,MEDICATION,UOM,REQUEST_TIME,NURSE_REMARKS,STATUS,PHARMACY_REMARKS) VALUES (?,?,?,?,?,?,?,?,?,?)", File_Number, Nurse_Name, Ward, Bed, Medication, UOM, Request_time.Format(time.ANSIC), Nurse_Remarks, Status, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	defer DB.Close()
	templates.ExecuteTemplate(w, "dashboard.html", nil)
	//http.Redirect(w, r, "/Order", http.StatusSeeOther)
}

func TrackOrderHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "TrackOrder.html", nil)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "register.html", nil)
}

func userRegisterHandler(w http.ResponseWriter, r *http.Request) {
	//Hash the password before storing it in the database
	DB, err := sql.Open("sqlite3", "./DB.db")
	if err != nil {
		log.Fatal(err)
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	_, err = DB.Exec("INSERT INTO users (username, password) VALUES (?,?)", username, hashedPassword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	const port = "8080"
	mux := http.NewServeMux()

	mux.HandleFunc("/", loginPage)  //.Methods("GET")
	mux.HandleFunc("/login", login) //.Methods("POST")
	mux.HandleFunc("/dashboard", dashboardHandler)
	mux.HandleFunc("/logout", logoutHandler)
	mux.HandleFunc("/Order", OrderHandler)
	mux.HandleFunc("/Submit", SubmitHandler)
	mux.HandleFunc("/TrackOrder", TrackOrderHandler)
	mux.HandleFunc("/register", userRegisterHandler)
	mux.HandleFunc("/reg", RegisterHandler)
	//http.HandleFunc("/authenticate", authenticate)

	// server config
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	fmt.Println("Server started at :", port)
	log.Fatal(server.ListenAndServe())
}
