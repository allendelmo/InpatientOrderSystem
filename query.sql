-- name: CreateMedicationOrder :exec
INSERT INTO
    medication_orders (
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
VALUES
    (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetMedicationOrderList :many
SELECT
    file_number,
    nurse_name,
    ward,
    bed,
    request_time,
    status
FROM
    medication_orders
WHERE
    STATUS = 'PENDING'