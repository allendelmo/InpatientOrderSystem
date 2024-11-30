package main

import (
	"ImpatientOrderSystem/internal/auth"
	"ImpatientOrderSystem/internal/database"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `json:"id"`
	Username   string    `json:"username"`
	Ward       string    `json:"ward"`
	Permission string    `json:"permission"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
}

func (cfg *config) handlerRegisterUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		UserName   string `json:"username"`
		Password   string `json:"password"`
		FirstName  string `json:"first_name"`
		LastName   string `json:"last_name"`
		Ward       string `json:"ward"`
		Permission string `json:"permission"`
	}
	type response struct {
		User
	}
	var err error

	// username := r.FormValue("username")
	// password := r.FormValue("password")
	// firstName := r.FormValue("First Name")
	// lastName := r.FormValue("Last Name")
	// ward := r.FormValue("Ward")
	// permission := r.FormValue("Permission")
	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&params)

	hashedPassword, err := auth.HashPassword(params.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Cannot hash password", err)
		return
	}

	userDB, err := cfg.db.RegisterUser(r.Context(), database.RegisterUserParams{
		Username:       params.UserName,
		HashedPassword: hashedPassword,
		Ward:           params.Ward,
		Permission:     params.Permission,
		FirstName: sql.NullString{
			String: params.FirstName,
			Valid:  true,
		},
		LastName: sql.NullString{
			String: params.LastName,
			Valid:  true,
		},
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Cannot register user", err)
		return
	}

	respondWithJSON(w, http.StatusOK, response{
		User: User{
			ID:         userDB.ID,
			Username:   userDB.Username,
			Ward:       userDB.Ward,
			Permission: userDB.Permission,
			CreatedAt:  userDB.CreatedAt,
			UpdatedAt:  userDB.UpdatedAt,
			FirstName:  userDB.FirstName.String,
			LastName:   userDB.LastName.String,
		},
	})
}
