-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tigerbeetle_id_counter
(
     current_id BIGINT NOT NULL
);

INSERT INTO tigerbeetle_id_counter (current_id) VALUES (1);
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS tigerbeetle_id_counter;