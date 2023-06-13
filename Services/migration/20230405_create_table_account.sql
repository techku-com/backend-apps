CREATE SCHEMA IF NOT EXISTS accounts;
CREATE TYPE account_role AS ENUM ('TECHNICIAN', 'CUSTOMER');
CREATE TABLE accounts.t_user_accounts (
     id                   SERIAL PRIMARY KEY,
     username             VARCHAR(50) UNIQUE,
     email                VARCHAR(50) UNIQUE,
     password             VARCHAR(255),
     role                 account_role DEFAULT 'USER',
     is_admin             BOOLEAN,
     phone_number         VARCHAR(50),
     created_at           TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
     updated_at           TIMESTAMP WITHOUT TIME ZONE
);