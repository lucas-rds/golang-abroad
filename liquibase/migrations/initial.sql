--liquibase formatted sql

--changeset lucas.santos:1
CREATE TABLE IF NOT EXISTS agency (
    id serial PRIMARY KEY,
    name VARCHAR(255)
);
--rollback DROP TABLE IF EXISTS agency;