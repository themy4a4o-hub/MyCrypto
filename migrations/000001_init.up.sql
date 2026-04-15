CREATE SCHEMA cryptoapp;

CREATE TABLE cryptoapp.rates (
    id             SERIAL         PRIMARY KEY,
    cryptocurrency VARCHAR(10)    NOT NULL,
    price          NUMERIC(20,2)  NOT NULL,
    timestamp      TIMESTAMP      NOT NULL
);

CREATE INDEX idx_rates_crypto_time on cryptoapp.rates (cryptocurrency, timestamp);

CREATE TABLE cryptoapp.subscriptions (
    id                SERIAL PRIMARY KEY,
    chat_id           BIGINT                NOT NULL,
    interval_minutes  INT                   NOT NULL,
    is_active         BOOLEAN               NOT NULL    DEFAULT TRUE,
    last_sent_at      TIMESTAMP
);
CREATE INDEX idx_subs_active ON cryptoapp.subscriptions (chat_id) WHERE is_active;
