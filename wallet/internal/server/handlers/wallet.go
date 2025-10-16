package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"wallet/internal/domain/entity"
	"wallet/internal/server/gen"
)

type WalletService interface {
	Deposit(ctx context.Context, amount int, walletId int) (int, int, error)
	Withdraw(ctx context.Context, amount int, walletId int) (int, int, error)
	GetWallet(ctx context.Context, walletId int) (entity.Wallet, error)
	GetTransactions(ctx context.Context, transactionId int) ([]entity.Transaction, error)
	CreateWallet(ctx context.Context) (int, error)
}

type WalletHandler struct {
	WalletService WalletService
}

func NewWalletHandler(service WalletService) *WalletHandler {
	return &WalletHandler{
		WalletService: service,
	}
}

func (h *WalletHandler) GetWallet(ctx echo.Context, params gen.GetWalletParams) error {
	wallet, err := h.WalletService.GetWallet(ctx.Request().Context(), params.WalletId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, wallet)
}

func (h *WalletHandler) PostWalletDeposit(ctx echo.Context, params gen.PostWalletDepositParams) error {
	var body gen.PostWalletDepositJSONBody
	if err := ctx.Bind(&body); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid_json",
		})
	}
	amount := body.Amount

	transactionID, newBalance, err := h.WalletService.Deposit(ctx.Request().Context(), amount, params.WalletId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, map[string]int{
		"transaction_id": transactionID,
		"newBalance":     newBalance,
	})
}

func (h *WalletHandler) PostWalletWithdraw(ctx echo.Context, params gen.PostWalletWithdrawParams) error {
	var body gen.PostWalletDepositJSONBody
	if err := ctx.Bind(&body); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid_json",
		})
	}
	amount := body.Amount

	transactionID, newBalance, err := h.WalletService.Withdraw(ctx.Request().Context(), amount, params.WalletId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, map[string]int{
		"transaction_id": transactionID,
		"newBalance":     newBalance,
	})
}

func (h *WalletHandler) GetWalletTransactions(ctx echo.Context, params gen.GetWalletTransactionsParams) error {
	transactions, err := h.WalletService.GetTransactions(ctx.Request().Context(), params.WalletId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, transactions)
}

func (h *WalletHandler) PostWallet(ctx echo.Context) error {
	id, err := h.WalletService.CreateWallet(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, id)
}
