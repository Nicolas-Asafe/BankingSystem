package entities

import (
	"errors"

	"github.com/Nicolas-Asafe/BankingSystem/utils"
)

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


func (u *User) FindMessage(id string) utils.Response{
	for _,msg := range u.Notifications{
		if msg.Id == id{
			return utils.Resthis("Message found",nil,msg)
		}
	}
	return utils.Resthis("Message not found",errors.New("Message not exists"),nil)
}

func (u *User) FindMessages()utils.Response{
	return utils.Resthis("Messages found",nil,u.Notifications)
}