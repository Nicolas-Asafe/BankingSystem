package entities

import "github.com/Nicolas-Asafe/BankingSystem/utils"

type User struct {
	Name               string
	Password           string
	Email              string
	Age                int64
	Id                 string
	Balance            float64
	HistoryTransitions []Transition
	Notifications      []Message
	Log                utils.Logs
}


