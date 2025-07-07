package services

import "github.com/Nicolas-Asafe/BankingSystem/core/entities"
import "github.com/Nicolas-Asafe/BankingSystem/core/interfaces"
import "github.com/Nicolas-Asafe/BankingSystem/utils"

type EconomyService struct{
	Repo interfaces.EconomyRepository
}

func (e EconomyService) FindOneEconomy(Id string,AuthoId string){
	res:=e.Repo.findOne(AuthoId,Id)
	if res.Error != nil{
		utils.Resthis("Error searching for economy",res.Error,nil)
		return
	}
	utils.Resthis("Economy finded successfully",nil,res.Data)
}
func (e EconomyService) FindEconomys(){
	res:=e.Repo.find()
	if res.Error != nil{
		utils.Resthis("Error searching for economys",res.Error,nil)
		return
	}
	utils.Resthis("Economys listed successfully",nil,res.Data)
}
func (e EconomyService) CreateEconomy(eco entities.Economy){
	res:=e.Repo.newEconomy(eco)
	if res.Error != nil{
		utils.Resthis("Error for create a new economy",res.Error,nil)
		return
	}
	utils.Resthis("Economy created successfully",nil,res.Data)
}
func (e EconomyService) DeleteEconomy(Id string){
	res:=e.Repo.removeEconomy(Id)
	if res.Error != nil {
			utils.Resthis("Error for delete this economy",res.Error,nil)
			return
	}
	utils.Resthis("Economy deleted successfully",nil,res.Data)
}