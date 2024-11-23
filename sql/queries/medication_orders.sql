-- name: CreateMedicationOrder :exec
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
VALUES (DEFAULT, $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);
-- name: GetMedicationOrderList :many
SELECT *
FROM medication_orders
WHERE STATUS = 'PENDING';
-- name: GetReadytoCollect :many
SELECT *
FROM medication_orders
WHERE STATUS = 'READY TO COLLECT';
-- name: UpdateMedicationOrder :exec
UPDATE medication_orders
SET STATUS = 'READY TO COLLECT'
WHERE ORDER_NUMBER = $1;