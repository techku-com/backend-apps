CREATE SCHEMA IF NOT EXISTS orders;
CREATE TABLE orders.t_orders (
     id                   SERIAL PRIMARY KEY,
     created_by           INTEGER,
     taken_by             INTEGER,
     status               INTEGER,
     issues               VARCHAR(500),
     address              VARCHAR(500),
     created_at           TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
     updated_at           TIMESTAMP WITHOUT TIME ZONE
);

CREATE TABLE orders.t_orders_history (
     id                     SERIAL PRIMARY KEY,
     order_id               INTEGER,
     price                  INTEGER,
     description            VARCHAR(500),
     created_at             TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW()
);