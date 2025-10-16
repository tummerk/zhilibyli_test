package entity

import "time"

type TransactionType string

const (
	TransactionTypeDeposit    TransactionType = "deposit"    // Пополнение
	TransactionTypeWithdrawal TransactionType = "withdrawal" // Списание
)

type Transaction struct {
	ID         int
	Amount     int
	Type       TransactionType
	OldBalance int
	NewBalance int
	WalletID   int
	CreatedAt  time.Time
}
