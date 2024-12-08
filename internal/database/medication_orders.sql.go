// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: medication_orders.sql

package database

import (
	"context"
	"database/sql"
	"time"
)

const createMedicationOrder = `-- name: CreateMedicationOrder :exec
INSERT INTO medication_orders (
        order_number,
        file_number,
        nurse_name,
        ward,
        bed,
        medication,
        quantity,
        uom,
        request_time,
        nurse_remarks,
        status_id,
        pharmacy_remarks
    )
VALUES (DEFAULT, $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
`

type CreateMedicationOrderParams struct {
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

func (q *Queries) CreateMedicationOrder(ctx context.Context, arg CreateMedicationOrderParams) error {
	_, err := q.db.ExecContext(ctx, createMedicationOrder,
		arg.FileNumber,
		arg.NurseName,
		arg.Ward,
		arg.Bed,
		arg.Medication,
		arg.Quantity,
		arg.Uom,
		arg.RequestTime,
		arg.NurseRemarks,
		arg.StatusID,
		arg.PharmacyRemarks,
	)
	return err
}

const getMedicationOrderList = `-- name: GetMedicationOrderList :many
SELECT order_number, file_number, nurse_name, ward, bed, medication, quantity, uom, request_time, nurse_remarks, status_id, pharmacy_remarks
FROM medication_orders
WHERE STATUS = 'PENDING'
`

func (q *Queries) GetMedicationOrderList(ctx context.Context) ([]MedicationOrder, error) {
	rows, err := q.db.QueryContext(ctx, getMedicationOrderList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []MedicationOrder
	for rows.Next() {
		var i MedicationOrder
		if err := rows.Scan(
			&i.OrderNumber,
			&i.FileNumber,
			&i.NurseName,
			&i.Ward,
			&i.Bed,
			&i.Medication,
			&i.Quantity,
			&i.Uom,
			&i.RequestTime,
			&i.NurseRemarks,
			&i.StatusID,
			&i.PharmacyRemarks,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getReadytoCollect = `-- name: GetReadytoCollect :many
SELECT order_number, file_number, nurse_name, ward, bed, medication, quantity, uom, request_time, nurse_remarks, status_id, pharmacy_remarks
FROM medication_orders
WHERE STATUS = 'READY TO COLLECT'
`

func (q *Queries) GetReadytoCollect(ctx context.Context) ([]MedicationOrder, error) {
	rows, err := q.db.QueryContext(ctx, getReadytoCollect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []MedicationOrder
	for rows.Next() {
		var i MedicationOrder
		if err := rows.Scan(
			&i.OrderNumber,
			&i.FileNumber,
			&i.NurseName,
			&i.Ward,
			&i.Bed,
			&i.Medication,
			&i.Quantity,
			&i.Uom,
			&i.RequestTime,
			&i.NurseRemarks,
			&i.StatusID,
			&i.PharmacyRemarks,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateMedicationOrder = `-- name: UpdateMedicationOrder :exec
UPDATE medication_orders
SET STATUS = 'READY TO COLLECT'
WHERE ORDER_NUMBER = $1
`

func (q *Queries) UpdateMedicationOrder(ctx context.Context, orderNumber int32) error {
	_, err := q.db.ExecContext(ctx, updateMedicationOrder, orderNumber)
	return err
}