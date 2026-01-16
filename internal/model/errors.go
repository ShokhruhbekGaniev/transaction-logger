package model

import "errors"

var (
	ErrInvalidAmount = errors.New("amount must be greater than 0")

	ErrInvalidCurrency = errors.New("currency must be one of: USD, EUR, RUB")

	ErrInvalidType = errors.New("type must be one of: deposit, withdrawal, transfer")

	ErrMissingAccountTo = errors.New("account_to is required")

	ErrMissingAccountFrom = errors.New("account_from is required for transfer")
)
