-- +goose Up

INSERT INTO currency(name, ledger, min_volume, max_volume)
VALUES ('EUR', 100, 1000000, 100000000),
       ('USD', 200, 1000000, 100000000),
       ('JPY', 300, 1000000, 100000000),
       ('GBP', 400, 1000000, 100000000), 
       ('AUD', 500, 1000000, 100000000);

INSERT INTO app_entity (external_id, kind)
VALUES ('MASTER', 'MASTER'),
       ('PROVIDER_1', 'PROVIDER');


INSERT INTO fee (currency_pair, percent)    
VALUES 
    ('USD/EUR', 0.02),
    ('EUR/USD', 0.02),
    ('USD/JPY', 0.02),
    ('JPY/USD', 0.02),
    ('USD/GBP', 0.02),
    ('GBP/USD', 0.02),
    ('USD/AUD', 0.02),
    ('AUD/USD', 0.02),
    ('EUR/JPY', 0.02),
    ('JPY/EUR', 0.02),
    ('EUR/GBP', 0.02),
    ('GBP/EUR', 0.02),
    ('EUR/AUD', 0.02),
    ('AUD/EUR', 0.02),
    ('GBP/JPY', 0.02),
    ('JPY/GBP', 0.02),
    ('GBP/AUD', 0.02),
    ('AUD/GBP', 0.02),
    ('AUD/JPY', 0.02),
    ('JPY/AUD', 0.02);

-- +goose Down
