CREATE TABLE "users" (
    user_id text,
    email text,
    password VARCHAR,


    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    PRIMARY KEY (user_id)
);