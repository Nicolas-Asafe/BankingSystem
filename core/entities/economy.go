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

type EconomyReturn struct{
	Can  bool
	Value float64
	Error error
}

func (e *Economy) AddValue(amount float64) EconomyReturn{
	if amount < e.MinValueForAdd {
		return EconomyReturn{false,0,errors.New("")}
	}
	valueFinal := e.Value+amount
	return EconomyReturn{true,valueFinal,nil}
}

func (e *Economy) Withdraw(amount float64) EconomyReturn {
	if amount > e.MaxValueForWithdraw {
		return EconomyReturn{false,0, errors.New("the maximum value was reached")}
	}
	if amount < e.MinValueForWithdraw {
		return EconomyReturn{false,0, errors.New("the maximum value was reached")}
	}
	if amount > e.Value {
		return EconomyReturn{false,0,errors.New("you do not have enough balance")}
	}

	valueFinal:= e.Value-amount
	return EconomyReturn{true,valueFinal,nil}
}
