package repository

import (
	"context"
	"database/sql"
	"fmt"
	"wallet/internal/domain/entity"
)

type WalletRepository struct {
	db *sql.DB
}

func NewWalletRepository(db *sql.DB) WalletRepository {
	return WalletRepository{
		db,
	}
}

func (wr *WalletRepository) CreateTransaction(ctx context.Context, transaction entity.Transaction) (int, error) {
	query := `
        INSERT INTO transactions (wallet_id, type, amount,old_balance,new_balance) 
        VALUES ($1, $2, $3,$4, $5) 
        RETURNING id, created_at
    `

	err := wr.db.QueryRowContext(ctx,
		query,
		transaction.WalletID,
		transaction.Type,
		transaction.Amount,
		transaction.OldBalance,
		transaction.NewBalance,
	).Scan(&transaction.ID, &transaction.CreatedAt)

	return transaction.ID, err
}

func (wr *WalletRepository) GetBalance(ctx context.Context, walletId int) (int, error) {
	query := `SELECT balance FROM wallets WHERE id = $1`
	var balance int
	err := wr.db.QueryRow(query, walletId).Scan(&balance)
	return balance, err
}

func (wr *WalletRepository) CreateWallet(ctx context.Context) (int, error) {
	var walletID int
	query := `INSERT INTO wallets (balance) VALUES (0) RETURNING id`

	err := wr.db.QueryRowContext(ctx, query).Scan(&walletID)
	if err != nil {
		return 0, err
	}

	return walletID, nil
}

func (wr *WalletRepository) ChangeBalance(ctx context.Context, walletID int, amount int) error {
	const op = "WalletRepository.ChangeBalance"
	query := `UPDATE wallets 
		SET balance = balance + $1 
		WHERE id = $2 AND balance + $1 >= 0
		RETURNING balance`
	res, err := wr.db.ExecContext(ctx, query, amount, walletID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: failed to get rows affected: %w", op, err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("%s: wallet not found or insufficient funds", op)
	}
	return nil
}

func (wr *WalletRepository) GetTransactions(ctx context.Context, walletID int) ([]entity.Transaction, error) {
	const op = "GetTransactions"
	query := `SELECT id, wallet_id, type, amount,old_balance,new_balance ,created_at 
				FROM transactions 
				WHERE wallet_id = $1
				ORDER BY id ASC `

	rows, err := wr.db.QueryContext(ctx, query, walletID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var transactions []entity.Transaction

	for rows.Next() {
		var t entity.Transaction
		err = rows.Scan(
			&t.ID,
			&t.WalletID,
			&t.Type,
			&t.Amount,
			&t.OldBalance,
			&t.NewBalance,
			&t.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: failed to scan row: %w", op, err)
		}
		transactions = append(transactions, t)
	}

	return transactions, rows.Err()
}
