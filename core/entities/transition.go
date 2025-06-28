package entities

import "github.com/Nicolas-Asafe/BankingSystem/utils"

type Transition struct {
	Value            float64
	FromId           string
	ForId            string
	Description      *string
	DateOfTransition string
	TransitionId     string
	Log              utils.Logs
}
