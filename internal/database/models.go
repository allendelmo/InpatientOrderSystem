// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type MedicationOrder struct {
	OrderNumber     int32
	FileNumber      int32
	NurseName       sql.NullString
	Ward            sql.NullString
	Bed             sql.NullString
	Medication      sql.NullString
	Quantity        sql.NullInt32
	Uom             sql.NullString
	RequestTime     time.Time
	NurseRemarks    sql.NullString
	StatusID        int32
	PharmacyRemarks sql.NullString
}

type User struct {
	ID             uuid.UUID
	Username       string
	HashedPassword string
	Ward           string
	PermissionID   int32
	CreatedAt      time.Time
	UpdatedAt      time.Time
	FirstName      sql.NullString
	LastName       sql.NullString
}
