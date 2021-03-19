-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE "category" (
    name text,
    icon text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    PRIMARY KEY (name)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE "category";