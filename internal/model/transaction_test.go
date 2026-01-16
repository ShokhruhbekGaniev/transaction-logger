package model

import (
	"testing"
	"time"
)

func TestTransaction_Validate(t *testing.T) {
	tests := []struct {
		name    string      // Имя теста (для понимания что проверяем)
		tx      Transaction // Входные данные
		wantErr error       // Ожидаемая ошибка (nil если валидно)
	}{
		{
			name: "valid deposit transaction",
			tx: Transaction{
				Type:      TypeDeposit,
				AccountTo: "ACC001",
				Amount:    100.50,
				Currency:  CurrencyUSD,
				Status:    StatusPending,
				CreatedAt: time.Now(),
			},
			wantErr: nil,
		},
		{
			name: "valid transfer transaction",
			tx: Transaction{
				Type:        TypeTransfer,
				AccountFrom: "ACC001",
				AccountTo:   "ACC002",
				Amount:      50.00,
				Currency:    CurrencyEUR,
				Status:      StatusCompleted,
				CreatedAt:   time.Now(),
			},
			wantErr: nil,
		},
		{
			name: "invalid amount - zero",
			tx: Transaction{
				Type:      TypeDeposit,
				AccountTo: "ACC001",
				Amount:    0,
				Currency:  CurrencyUSD,
			},
			wantErr: ErrInvalidAmount,
		},
		{
			name: "invalid amount - negative",
			tx: Transaction{
				Type:      TypeDeposit,
				AccountTo: "ACC001",
				Amount:    -100,
				Currency:  CurrencyUSD,
			},
			wantErr: ErrInvalidAmount,
		},
		{
			name: "invalid currency",
			tx: Transaction{
				Type:      TypeDeposit,
				AccountTo: "ACC001",
				Amount:    100,
				Currency:  "GBP", // Не поддерживается
			},
			wantErr: ErrInvalidCurrency,
		},
		{
			name: "invalid type",
			tx: Transaction{
				Type:      "refund", // Не поддерживается
				AccountTo: "ACC001",
				Amount:    100,
				Currency:  CurrencyUSD,
			},
			wantErr: ErrInvalidType,
		},
		{
			name: "missing account_to",
			tx: Transaction{
				Type:     TypeDeposit,
				Amount:   100,
				Currency: CurrencyUSD,
			},
			wantErr: ErrMissingAccountTo,
		},
		{
			name: "transfer without account_from",
			tx: Transaction{
				Type:      TypeTransfer,
				AccountTo: "ACC002",
				Amount:    100,
				Currency:  CurrencyUSD,
			},
			wantErr: ErrMissingAccountFrom,
		},
		{
			name: "withdrawal without account_from is valid",
			tx: Transaction{
				Type:      TypeWithdrawal,
				AccountTo: "ACC001",
				Amount:    100,
				Currency:  CurrencyRUB,
			},
			wantErr: nil, // Для withdrawal AccountFrom не обязателен
		},
	}

	// Запускаем каждый тест
	for _, tt := range tests {
		// t.Run создаёт подтест с именем tt.name
		t.Run(tt.name, func(t *testing.T) {
			err := tt.tx.Validate()

			// Проверяем результат
			if err != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}
