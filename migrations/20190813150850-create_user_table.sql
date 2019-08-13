
-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
   id SERIAL PRIMARY KEY,
   name VARCHAR NOT NULL
);
-- +migrate Down
DROP TABLE users;
