-- +goose Up
INSERT INTO STATUSES(
    ID,
    NAME
)
VALUES(
    gen_random_uuid(),
    'PENDING'
);

INSERT INTO STATUSES(
    ID,
    NAME
)
VALUES(
    gen_random_uuid(),
    'READY TO COLLECT'
);

INSERT INTO STATUSES(
    ID,
    NAME
)
VALUES(
    gen_random_uuid(),
    'CANCELLED'
);

INSERT INTO STATUSES(
    ID,
    NAME
)
VALUES(
    gen_random_uuid(),
    'PARTIALLY DISPENSED'
);

-- +goose Down
