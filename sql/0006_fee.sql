-- +goose Up
CREATE TABLE IF NOT EXISTS fee (
     currency_pair                      TEXT NOT NULL UNIQUE,
     percent                            NUMERIC NOT NULL    
);

-- +goose Down
DROP TABLE IF EXISTS fee CASCADE;
