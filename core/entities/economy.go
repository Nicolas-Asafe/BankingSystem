package entities

import (
	"errors"
	"github.com/Nicolas-Asafe/BankingSystem/utils"
)

type Economy struct {
	MinValueForAdd      float64
	MinValueForWithdraw float64
	MaxValueForWithdraw float64
	AuthorID            string
	ID                  string
	Description         *string
	Title               string
	Value               float64
	Log                 utils.Logs
}

func (e *Economy) AddValue(amount float64) {
	if amount < e.MinValueForAdd {
		return
	}
	e.Value += amount
}

func (e *Economy) Withdraw(amount float64) (float64, error) {
	if amount > e.MaxValueForWithdraw {
		return 0, errors.New("the maximum value was reached")
	}
	if amount < e.MinValueForWithdraw {
		return 0, errors.New("the minimum value was reached")
	}
	if amount > e.Value {
		return 0, errors.New("you do not have enough balance")
	}

	e.Value -= amount
	return amount, nil
}
