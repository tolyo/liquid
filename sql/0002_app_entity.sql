
-- +goose Up
CREATE TABLE IF NOT EXISTS app_entity
(
    id                  BIGSERIAL PRIMARY KEY,
    pub_id              TEXT default uuid_generate_v4() UNIQUE NOT NULL,
    kind                TEXT NOT NULL,   -- CLIENT, PROVIDER, sMASTER
    external_id         TEXT NOT NULL UNIQUE,
    updated_at          TIMESTAMP default current_timestamp NOT NULL,
    created_at          TIMESTAMP default current_timestamp NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS app_entity;