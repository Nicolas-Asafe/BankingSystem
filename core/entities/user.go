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
	Economys           []Economy
}



func (u *User) FindMessage(id string) utils.Response {
	for _, msg := range u.Notifications {
		if msg.Id == id {
			return utils.Resthis("Message found", nil, msg)
		}
	}
	return utils.Resthis("Message not found", errors.New("Message not exists"), nil)
}

func (u *User) FindMessages() utils.Response {
	return utils.Resthis("Messages found", nil, u.Notifications)
}

// Economys separations

func (u *User) NewEconomy(NewEco Economy) utils.Response{
	u.Economys = append(u.Economys, NewEco)
	return utils.Resthis("Economy created successfully",nil,NewEco)
}
func (u *User) RemoveEconomy(EcoId string) utils.Response{
	for i, eco := range u.Economys {
		if eco.ID == EcoId {
			u.Economys = append(u.Economys[:i],u.Economys[:i+1]... )
			return utils.Resthis("Economy deleted successfully",nil,eco)
		}
	}
	return utils.Resthis("Economy not found",errors.New("Economy not exists"),nil)
}