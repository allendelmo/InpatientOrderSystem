-- name: RegisterUser :one
INSERT INTO users (
        id,
        created_at,
        updated_at,
        username,
        hashed_password,
        ward,
        permission,
        first_name,
        last_name
    )
VALUES (gen_random_uuid(), NOW(), NOW(), $1, $2, $3, $4, $5, $6)
RETURNING *;