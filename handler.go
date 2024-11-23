package main

import (
	"ImpatientOrderSystem/internal/database"
	"context"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	Username string
	Password string
}

type Medication_Orders struct {
	Order_Number     int32
	File_Number      int32
	Nurse_Name       string
	Ward             string
	Bed              string
	Medication       string
	UOM              string
	Request_time     time.Time
	Nurse_Remarks    string
	Status           string
	PHARMACY_REMARKS string
}

// Template cache
var templates = template.Must(template.ParseGlob("templates/*.html"))

// TODO: rework auth
var (
	key   = []byte("super-secret-key") // Change this in production
	store = sessions.NewCookieStore(key)
)

// TODO: rework auth
func loginPage(w http.ResponseWriter, r *http.Request) {
	//session, _ := store.Get(r, "session")
	templates.ExecuteTemplate(w, "login.html", nil)
}

// TODO: rework auth
func authenticate(username, password, dbUrl string) (bool, error) {
	var hashedPassword string

	DB, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	// Query database for user
	query := "SELECT password FROM users WHERE username = $1"
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
func (cfg *config) login(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	password := r.FormValue("password")

	// Authenticate user
	authenticated, err := authenticate(username, password, cfg.dbUrl)
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

// Logout handler
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Successfully logged out!")
}

// Order Handler
func OrderHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "Order.html", nil)
}

// Submit Handler
func (cfg *config) SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		fileNumberInt, err := strconv.Atoi(r.FormValue("File_Number"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		createMedicationOrderParams := database.CreateMedicationOrderParams{
			FileNumber:      int32(fileNumberInt),
			NurseName:       sql.NullString{String: r.FormValue("Nurse_Name"), Valid: true},
			Ward:            sql.NullString{String: r.FormValue("Ward"), Valid: true},
			Bed:             sql.NullString{String: r.FormValue("Bed"), Valid: true},
			Medication:      sql.NullString{String: r.FormValue("Medication"), Valid: true},
			Uom:             sql.NullString{String: r.FormValue("UOM"), Valid: true},
			NurseRemarks:    sql.NullString{String: r.FormValue("Nurse_Remarks"), Valid: true},
			PharmacyRemarks: sql.NullString{},
			RequestTime:     time.Now(),
			Status:          "PENDING",
		}

		err = cfg.db.CreateMedicationOrder(r.Context(), createMedicationOrderParams)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func TrackOrderHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "TrackOrder.html", nil)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "register.html", nil)
}

// TODO: refactor db to use SQLC
func (cfg *config) userRegisterHandler(w http.ResponseWriter, r *http.Request) {
	//Hash the password before storing it in the database
	DB, err := sql.Open("postgres", cfg.dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")
		First_Name := r.FormValue("First Name")
		Last_Name := r.FormValue("Last Name")
		Ward := r.FormValue("Ward")
		Permission := r.FormValue("Permission")
		createdAt := time.Now()

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return
		}

		_, err = DB.Exec("INSERT INTO users (id, username, password,ward,PERMISSION,createdat,first_name,last_name) VALUES (gen_random_uuid(),$1,$2,$3,$4,$5,$6,$7)", username, hashedPassword, Ward, Permission, createdAt.Format(time.ANSIC), First_Name, Last_Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}

// FOR DISPLAYING DATA IN DASHBOARD FOR ALL
func (cfg *config) displayhandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/dashboard.html"))

	rows, err := cfg.db.GetMedicationOrderList(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	medicationOrders := []Medication_Orders{}

	// transform from []sqlc.GetMedicationOrderListRow to []Medication_Orders
	for _, row := range rows {
		medicationOrders = append(medicationOrders, Medication_Orders{
			Order_Number:     row.OrderNumber,
			File_Number:      row.FileNumber,
			Nurse_Name:       row.NurseName.String,
			Ward:             row.Ward.String,
			Bed:              row.Bed.String,
			Medication:       row.Medication.String,
			UOM:              row.Uom.String,
			Request_time:     row.RequestTime.Round(2),
			Nurse_Remarks:    row.NurseRemarks.String,
			Status:           row.Status,
			PHARMACY_REMARKS: "",
		})
	}
	if err := tmpl.Execute(w, medicationOrders); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (cfg *config) CollectHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/collect.html"))

	rows, err := cfg.db.GetReadytoCollect(context.Background())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	MEDICATION_ORDER := []Medication_Orders{}

	for _, row := range rows {
		MEDICATION_ORDER = append(MEDICATION_ORDER, Medication_Orders{
			Order_Number:     row.OrderNumber,
			File_Number:      row.FileNumber,
			Nurse_Name:       row.NurseName.String,
			Ward:             row.Ward.String,
			Bed:              row.Bed.String,
			Medication:       row.Medication.String,
			UOM:              row.Uom.String,
			Request_time:     row.RequestTime,
			Nurse_Remarks:    row.NurseRemarks.String,
			Status:           row.Status,
			PHARMACY_REMARKS: "",
		})
	}

	if err := tmpl.Execute(w, MEDICATION_ORDER); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
