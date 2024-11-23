package main

import (
	"ImpatientOrderSystem/internal/database"
	"database/sql"
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type config struct {
	db *database.Queries
	// platform  string // TODO
	// jwtSecret string // TODO
	dbUrl string
}

func main() {
	const port = "8080"

	// get env variables
	godotenv.Load()
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("Cannot get env variable DB_URL")
	}

	// initialize DB
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)
	defer db.Close()

	// initialize config
	cfg := config{
		db:    dbQueries,
		dbUrl: dbUrl,
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
	mux.HandleFunc("/login", cfg.login)
	mux.HandleFunc("/dispense", cfg.handlerDispense)
	mux.HandleFunc("/logout", logoutHandler)
	mux.HandleFunc("/register", cfg.userRegisterHandler)
	mux.HandleFunc("/Submit", cfg.SubmitHandler)

	mux.HandleFunc("GET /api/medication_orders", cfg.handlerMedicationOrderList)
	mux.HandleFunc("POST /api/medication_orders", cfg.handlerMedicationOrderCreate)
	mux.HandleFunc("POST /api/users", cfg.handlerMedicationOrderCreate)
	//mux.HandleFunc("POST /api/medication_orders", cfg.handlerDispense)
	//http.HandleFunc("/authenticate", authenticate)

	// initialize server config
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	fmt.Println("Server started at :", port)
	log.Fatal(server.ListenAndServe())
}
