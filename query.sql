-- name: CreateMedicationOrder :exec
INSERT INTO medication_orders (
        file_number,
        nurse_name,
        ward,
        bed,
        medication,
        uom,
        request_time,
        nurse_remarks,
        status,
        pharmacy_remarks
    )
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
-- name: GetMedicationOrderList :many
SELECT *
FROM medication_orders
WHERE STATUS = 'PENDING';
-- name: GetReadytoCollect :many
SELECT *
FROM medication_orders
WHERE STATUS = 'READY TO COLLECT';