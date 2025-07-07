package services

import "github.com/Nicolas-Asafe/BankingSystem/core/entities"
import "github.com/Nicolas-Asafe/BankingSystem/core/interfaces"
import "github.com/Nicolas-Asafe/BankingSystem/utils"

type EconomyService struct{
	Repo interfaces.EconomyRepository
}

func (e EconomyService) FindOneEconomy(Id string,AuthoId string){
	res:=t.Repo.findOne(AuthoId,Id)
	if res.Error != nil{
		utils.Resthis("Error searching for economy",res.Error,nil)
		return
	}
	utils.Resthis("Economy finded successfully",nil,res.Data)
}