DROP TABLE IF EXISTS "category";

CREATE TABLE "category" (

    categoryid serial NOT NULL,
    name text NOT NULL,
    icon text NOT NULL,
    user_id text NOT NULL,

    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    PRIMARY KEY (categoryid)
);
