-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE "expenses" (
    id text,
    name text,
    icon text,
    amount numeric,
    note text,
    expense_date date,

    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    PRIMARY KEY (id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE "expenses";