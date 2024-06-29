CREATE TABLE IF NOT EXISTS accounts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    account_number VARCHAR(255) NOT NULL,
    balance DECIMAL(10,2) DEFAULT 0.00
);

INSERT INTO accounts (account_number, balance)
VALUES ('123456',1000.00), ('789012', 500.00);
