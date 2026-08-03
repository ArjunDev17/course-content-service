CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE courses
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    title VARCHAR(255) NOT NULL,

    description TEXT,

    category VARCHAR(100),

    instructor VARCHAR(255),

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);