CREATE TABLE "expenses" (

    id text NOT NULL,
    categoryid text,
    amount integer NOT NULL,
    note text,
    expense_date date NOT NULL,
    user_id text NOT NULL,

    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    PRIMARY KEY (id)
);