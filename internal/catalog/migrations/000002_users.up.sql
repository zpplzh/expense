CREATE TABLE "users" (
    user_id text NOT NULL,
    email text NOT NULL,
    password VARCHAR NOT NULL,


    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    PRIMARY KEY (user_id),
    CONSTRAINT email_unique UNIQUE (email)
);