-- name: CreateMedicationOrder :exec
INSERT INTO medication_orders (
        file_number,
        nurse_name,
        ward,
        bed,
        quantity,
        medication,
        uom,
        request_time,
        nurse_remarks,
        status,
        pharmacy_remarks
    )
VALUES (?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
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
WHERE ORDER_NUMBER = ?;