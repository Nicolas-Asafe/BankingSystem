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