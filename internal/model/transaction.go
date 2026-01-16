package model

import "time"

type TransactionType string

const (
	TypeDeposit    TransactionType = "deposit"    // Пополнение
	TypeWithdrawal TransactionType = "withdrawal" // Списание
	TypeTransfer   TransactionType = "transfer"   // Перевод
)

type TransactionStatus string

const (
	StatusPending   TransactionStatus = "pending"   // В обработке
	StatusCompleted TransactionStatus = "completed" // Завершена
	StatusFailed    TransactionStatus = "failed"    // Ошибка
)

type Currency string

const (
	CurrencyUSD Currency = "USD"
	CurrencyEUR Currency = "EUR"
	CurrencyRUB Currency = "RUB"
)

type Transaction struct {
	ID          string            `json:"id"`                     // Уникальный идентификатор
	Type        TransactionType   `json:"type"`                   // Тип: deposit, withdrawal, transfer
	AccountFrom string            `json:"account_from,omitempty"` // Счёт отправителя (опционально)
	AccountTo   string            `json:"account_to"`             // Счёт получателя (обязательно)
	Amount      float64           `json:"amount"`                 // Сумма транзакции
	Currency    Currency          `json:"currency"`               // Валюта: USD, EUR, RUB
	Description string            `json:"description,omitempty"`  // Описание (опционально)
	Status      TransactionStatus `json:"status"`                 // Статус транзакции
	CreatedAt   time.Time         `json:"created_at"`             // Время создания
}

func (t *Transaction) Validate() error {
	if t.Amount <= 0 {
		return ErrInvalidAmount
	}

	// Проверяем валюту
	if !isValidCurrency(t.Currency) {
		return ErrInvalidCurrency
	}

	// Проверяем тип транзакции
	if !isValidType(t.Type) {
		return ErrInvalidType
	}

	// AccountTo обязателен для всех типов
	if t.AccountTo == "" {
		return ErrMissingAccountTo
	}

	// Для transfer обязателен AccountFrom
	if t.Type == TypeTransfer && t.AccountFrom == "" {
		return ErrMissingAccountFrom
	}

	return nil
}

func isValidCurrency(c Currency) bool {
	switch c {
	case CurrencyUSD, CurrencyEUR, CurrencyRUB:
		return true
	default:
		return false
	}
}

func isValidType(t TransactionType) bool {
	switch t {
	case TypeDeposit, TypeWithdrawal, TypeTransfer:
		return true
	default:
		return false
	}
}
