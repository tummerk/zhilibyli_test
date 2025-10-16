CREATE TABLE transactions (
                             id SERIAL PRIMARY KEY,
                             wallet_id INTEGER NOT NULL,
                             amount    INTEGER NOT NULL ,
                             type VARCHAR(50) NOT NULL,
                             old_balance INTEGER not null,
                             new_balance INTEGER NOT NULL,
                             created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                             FOREIGN KEY (wallet_id) REFERENCES wallets(id)
);