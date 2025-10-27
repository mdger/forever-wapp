CREATE TABLE users (
    id UUID PRIMARY KEY,
    user_name VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(100) NOT NULL
);

CREATE TABLE user_actions (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES user_id (id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL
);
