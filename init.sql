CREATE TABLE wallets(
    id UUID PRIMARY KEY,
    balance DECIMAL(15, 2) NOT NULL DEFAULT 0 CHECK (balance >= 0)
);

--INSERT для тестов
INSERT INTO wallets (id, balance) VALUES ('ff41c5c7-4f9e-4161-a573-7e3b73d2cd55', 0);
