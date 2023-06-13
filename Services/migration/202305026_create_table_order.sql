CREATE SCHEMA IF NOT EXISTS orders;
CREATE TYPE hardware_type_enum AS ENUM ('LAPTOP', 'PERSONAL COMPUTER', 'PRINTER', 'SCANNER');
CREATE TYPE service_type_enum AS ENUM ('CLEANING', 'REPAIRMENT', 'MAINTENANCE');
CREATE TABLE orders.t_orders (
     id                   SERIAL PRIMARY KEY,
     created_by           INTEGER,
     taken_by             INTEGER,
     description          VARCHAR(500),
     location             VARCHAR(500),
     contact_person       VARCHAR(64),
     hardware_type        hardware_type_enum,
     service_type         service_type_enum,
     schedule_date        TIMESTAMP WITHOUT TIME ZONE,
     created_at           TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
     updated_at           TIMESTAMP WITHOUT TIME ZONE
);