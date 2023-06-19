CREATE SCHEMA IF NOT EXISTS articles;
CREATE TABLE articles.t_articles (
     id                   SERIAL PRIMARY KEY,
     title                VARCHAR(255),
     description          TEXT,
     image_url            TEXT,
     created_at           TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
     updated_at           TIMESTAMP WITHOUT TIME ZONE
);