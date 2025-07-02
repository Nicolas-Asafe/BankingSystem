package services

import (
	"github.com/Nicolas-Asafe/BankingSystem/core/entities"
	"github.com/Nicolas-Asafe/BankingSystem/core/interfaces"
	"github.com/Nicolas-Asafe/BankingSystem/utils"
)

type UserService struct{
	Repo interfaces.UserRepository
}


func (u UserService) NewUser(User entities.User) utils.Response{
	res:=u.Repo.NewUser(User)
	if res.Error != nil {
		utils.Resthis("Error for create user",res.Error,nil)
	}
	return utils.Resthis("User created successfully",nil,User)
}