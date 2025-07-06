package services

import (
	"github.com/Nicolas-Asafe/BankingSystem/core/entities"
	"github.com/Nicolas-Asafe/BankingSystem/core/interfaces"
	"github.com/Nicolas-Asafe/BankingSystem/utils"
)

type TransitionService struct {
	Repo interfaces.TransitionRepository
}

func (t TransitionService) SendTransition(tst entities.Transition){
	res:=t.Repo.Send(tst)
	if res.Error != nil{
		utils.Resthis("Error performing transition",res.Error,nil)
		return
	}
	utils.Resthis("Transaction completed successfully",nil,tst)
}

func (t TransitionService) ReciveTransition(id string,tst entities.Transition){
	res:=t.Repo.ReciveTransition(id,tst)
	if res.Error != nil{
		utils.Resthis("Error for recive transition",res.Error,nil)
		return
	}
	utils.Resthis("Transition recived successfully",nil,tst)
}