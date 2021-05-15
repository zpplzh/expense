CREATE TABLE "session" (
    sessionid text NOT NULL,
    expiry timestamp with time zone NOT NULL,
    user_id text NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    PRIMARY KEY (sessionid)
);
