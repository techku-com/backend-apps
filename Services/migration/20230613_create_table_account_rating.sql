CREATE SCHEMA IF NOT EXISTS accounts;
CREATE TABLE accounts.t_user_accounts_rating (
     id                   SERIAL PRIMARY KEY,
     created_by           INTEGER,
     created_to           INTEGER,
     order_id             INTEGER,
     description          VARCHAR(255),
     rating               INTEGER,
     created_at           TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW()
);