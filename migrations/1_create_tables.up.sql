-- +migrate Up
CREATE TABLE users (
    id UUID NOT NULL PRIMARY KEY,
    username  VARCHAR(100) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    email CHAR(100) NOT NULL,

    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- +migrate Up
CREATE TABLE channels (
    id UUID NOT NULL PRIMARY KEY,
    username VARCHAR(100) UNIQUE NOT NULL,
    description TEXT,
    logo TEXT,
    user_id UUID REFERENCES users(id)
);
