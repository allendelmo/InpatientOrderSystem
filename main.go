package main

import (
	"ImpatientOrderSystem/internal/database"
	"database/sql"
	_ "embed"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

type config struct {
	db *database.Queries
	// platform  string // TODO
	// jwtSecret string // TODO
}

func main() {
	const port = "8080"

	// initialize DB
	db, err := sql.Open("sqlite3", "./DB.db") // Open a connection to the SQlite database file named Todos.db
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)
	defer db.Close()

	// initialize config
	cfg := config{
		db: dbQueries,
	}

	// create server mux and register handlers
	mux := http.NewServeMux()
	mux.HandleFunc("/", loginPage)
	mux.HandleFunc("/dashboard", cfg.displayhandler)
	mux.HandleFunc("/reg", RegisterHandler)
	mux.HandleFunc("/collect", cfg.CollectHandler)
	mux.HandleFunc("/TrackOrder", TrackOrderHandler)
	mux.HandleFunc("/Order", OrderHandler)

	// TODO: fix api handlers (use Methods: GET, POST)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logoutHandler)
	mux.HandleFunc("/register", cfg.userRegisterHandler)
	mux.HandleFunc("/Submit", cfg.SubmitHandler)

	mux.HandleFunc("GET /api/medication_orders", cfg.handlerMedicationOrderList)
	mux.HandleFunc("POST /api/medication_orders", cfg.handlerMedicationOrderCreate)
	mux.HandleFunc("POST /api/users", cfg.handlerMedicationOrderCreate)

	//http.HandleFunc("/authenticate", authenticate)

	// initialize server config
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	fmt.Println("Server started at :", port)
	log.Fatal(server.ListenAndServe())
}
