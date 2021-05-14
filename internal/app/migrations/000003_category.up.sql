CREATE TABLE "category" (

    categoryid text NOT NULL,
    name text NOT NULL,
    icon text NOT NULL,
    user_id text NOT NULL,

    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    PRIMARY KEY (categoryid),
    CONSTRAINT category_unique_user UNIQUE(user_id,name)
);