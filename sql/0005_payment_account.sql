-- +goose Up
CREATE TABLE IF NOT EXISTS payment_account (
    id                      BIGSERIAL PRIMARY KEY,
    tigerbeetle_id          BIGINT default generate_tiger_id() UNIQUE NOT NULL,
    app_entity_id           BIGINT REFERENCES app_entity(id) NOT NULL,
    currency_name           TEXT REFERENCES currency(name) NOT NULL,
    updated_at              TIMESTAMP default current_timestamp NOT NULL,
    created_at              TIMESTAMP default current_timestamp NOT NULL,
    -- enforce one currency account per application entity
    UNIQUE (app_entity_id, currency_name)
);

-- +goose Down
DROP TABLE IF EXISTS payment_account CASCADE;
