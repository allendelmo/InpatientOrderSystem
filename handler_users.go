package main

import (
	"ImpatientOrderSystem/internal/auth"
	"ImpatientOrderSystem/internal/database"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	Ward         string    `json:"ward"`
	PermissionId string    `json:"permission_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
}

func (cfg *config) handlerRegisterUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		UserName     string `json:"username"`
		Password     string `json:"password"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		Ward         string `json:"ward"`
		PermissionId string `json:"permission_id"`
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
		respondWithError(w, http.StatusInternalServerError, "cannot hash password", err)
		return
	}

	permissionId, err := strconv.ParseInt(params.PermissionId, 10, 32)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "invalid permission id", err)
		return
	}

	userDB, err := cfg.db.RegisterUser(r.Context(), database.RegisterUserParams{
		Username:       params.UserName,
		HashedPassword: hashedPassword,
		Ward:           params.Ward,
		PermissionID:   int32(permissionId),
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
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				respondWithError(w, http.StatusInternalServerError, "username already taken", err)
				return
			}
		}

		respondWithError(w, http.StatusInternalServerError, "Cannot register user", err)
		return
	}

	respondWithJSON(w, http.StatusOK, response{
		User: User{
			ID:           userDB.ID,
			Username:     userDB.Username,
			Ward:         userDB.Ward,
			PermissionId: strconv.FormatInt(int64(userDB.PermissionID), 10),
			CreatedAt:    userDB.CreatedAt,
			UpdatedAt:    userDB.UpdatedAt,
			FirstName:    userDB.FirstName.String,
			LastName:     userDB.LastName.String,
		},
	})
}
