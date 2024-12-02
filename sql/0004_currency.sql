-- +goose Up
CREATE TABLE IF NOT EXISTS currency (
   name                               VARCHAR(3) PRIMARY KEY,
   precision                          INT default 2 NOT NULL,
   ledger                             INT NOT NULL UNIQUE,
   min_volume                         NUMERIC NOT NULL,
   max_volume                         NUMERIC NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS currency;