package service

import (
	"context"
	"errors"
	"log/slog"
	"wallet/internal/domain/entity"
)

type WalletRepository interface {
	ChangeBalance(ctx context.Context, walletID int, amount int) error
	GetBalance(ctx context.Context, walletID int) (int, error)
	CreateWallet(ctx context.Context) (int, error)
	GetTransactions(ctx context.Context, walletID int) ([]entity.Transaction, error)
	CreateTransaction(ctx context.Context, transaction entity.Transaction) (int, error)
}

type WalletService struct {
	repo WalletRepository
}

func NewWalletService(repo WalletRepository) *WalletService {
	return &WalletService{
		repo: repo,
	}
}

func (w *WalletService) Deposit(ctx context.Context, amount int, walletId int) (int, int, error) {
	if amount <= 0 {
		return 0, 0, errors.New("amount must be greater than zero")
	}
	return w.changeBalance(ctx, walletId, amount)
}

func (w *WalletService) Withdraw(ctx context.Context, amount int, walletId int) (int, int, error) {
	if amount <= 0 {
		return 0, 0, errors.New("amount must be greater than zero")
	}
	return w.changeBalance(ctx, walletId, -amount)
}

func (w *WalletService) GetTransactions(ctx context.Context, walletID int) ([]entity.Transaction, error) {
	return w.repo.GetTransactions(ctx, walletID)
}

func (w *WalletService) GetWallet(ctx context.Context, walletID int) (entity.Wallet, error) {
	balance, err := w.repo.GetBalance(ctx, walletID)
	if err != nil {
		slog.Error("error getting balance", slog.String("err", err.Error()))
		return entity.Wallet{}, err
	}
	return entity.Wallet{
		ID:      walletID,
		Balance: balance,
	}, nil
}

func (w *WalletService) CreateWallet(ctx context.Context) (int, error) {
	return w.repo.CreateWallet(ctx)
}

func (w *WalletService) changeBalance(ctx context.Context, walletID int, amount int) (int, int, error) {
	err := w.repo.ChangeBalance(ctx, walletID, amount)
	var newBalance int
	if err != nil {
		slog.Error("error change balance", slog.String("err", err.Error()))
		return 0, 0, err
	}
	newBalance, err = w.repo.GetBalance(ctx, walletID)
	if err != nil {
		slog.Error("error getting balance", slog.String("err", err.Error()))
		return 0, 0, err
	}

	var transaction entity.Transaction
	transaction.Type = entity.TransactionTypeWithdrawal
	if amount > 0 {
		transaction.Type = entity.TransactionTypeDeposit
	}

	transaction.Amount = amount
	transaction.WalletID = walletID
	transaction.NewBalance = newBalance
	transaction.OldBalance = newBalance - amount

	transactionID, err := w.repo.CreateTransaction(ctx, transaction)
	if err != nil {
		slog.Error("error creating transaction", slog.String("err", err.Error()))
	}
	return transactionID, newBalance, nil
}
