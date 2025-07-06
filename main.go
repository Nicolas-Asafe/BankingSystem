package main

import (
	"github.com/Nicolas-Asafe/BankingSystem/core/entities"
	"github.com/Nicolas-Asafe/BankingSystem/core/services"
	"github.com/Nicolas-Asafe/BankingSystem/utils"
)

func main() {
	log:=utils.Logs{
		WithTime: true,
	}
	userService:= services.UserService{}
	tstService:= services.TransitionService{}
	user1:=entities.User{
		Name: "Nicolas",
		Password: "1234",
		Email: "asdas@asd.com",
		Age: 12,
		Id: "23",
		Balance: 100.23,
		HistoryTransitions: []entities.Transition{},
		Notifications: []entities.Message{},
		Log: log,
		Economys: []entities.Economy{},
	}
	user2:=entities.User{
		Name: "Nicole",
		Password: "1234",
		Email: "asdas@asd.com",
		Age: 16,
		Id: "12",
		Balance: 10.23,
		HistoryTransitions: []entities.Transition{},
		Notifications: []entities.Message{},
		Log: log,
		Economys: []entities.Economy{},
	}
	tst:= entities.Transition{
		Value: 100,
		FromId: user1.Id,
		ForId: user2.Id,
		TransitionId: "asasd",
		Log: log,	
	}
	services.TransitionService.SendTransition(tstService,tst)
	services.TransitionService.ReciveTransition(tstService,tst)
	services.UserService.NewUser(userService,user1)
}