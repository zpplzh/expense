CREATE TABLE "expenses" (
    id text ,
    name text NOT NULL,
    icon text NOT NULL,
    amount integer NOT NULL,
    note text,
    expense_date date NOT NULL,

    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    PRIMARY KEY (id)
);