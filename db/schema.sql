CREATE SCHEMA IF NOT EXISTS billing_service;

CREATE TABLE billing_service.accounts  (
    id INT PRIMARY KEY,
    user_id serial,
    real_account money DEFAULT 0 CHECK (real_account >= 0 :: money),
    reserving_account money DEFAULT 0 CHECK (
        (reserving_account >= 0 :: money)
        AND (reserving_account <= real_account)
    )
);

CREATE TABLE billing_service.operations  (
    id INT PRIMARY KEY,
    cost money,
    account_id INT REFERENCES billing_service.accounts(id),
    service_id serial,
    -- 0 is pending, 1 is approved, 2 is failed
    health_status int DEFAULT 0,
    creation_date timestamp DEFAULT CURRENT_TIMESTAMP, 
    update_date timestamp DEFAULT CURRENT_TIMESTAMP
);