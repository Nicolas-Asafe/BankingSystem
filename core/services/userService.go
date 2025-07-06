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
func (u UserService) DeleteUser(Id string) utils.Response{
	res:=u.Repo.RemoveUser(Id)
	if res.Error != nil {
		utils.Resthis("Error for delete user",res.Error,nil)
	}
	return utils.Resthis("User deleted successfully",nil,nil)
}
func (u UserService) EditUserName(Id string,Name string) utils.Response{
	res:=u.Repo.PutUserName(Id,Name)
	if res.Error != nil {
		utils.Resthis("Error for edit name user",res.Error,nil)
	}
	return utils.Resthis("User name edited successfully",nil,Name)
}
func (u UserService) FindOneUser(Id string) utils.Response{
	res:=u.Repo.FindUser(Id)
	if res.Error != nil {
		utils.Resthis("Error for find name user",res.Error,nil)
	}
	return utils.Resthis("User name finded successfully",nil,res.Data)
}
func (u UserService) ListUsers() utils.Response{
	res:=u.Repo.FindUsers()
	if res.Error != nil {
		utils.Resthis("Error for list users",res.Error,nil)
	}
	return utils.Resthis("Users listed successfully",nil,res.Data)
}
func (u UserService) RemoveValueFrom(Id string,value float64) utils.Response{
	res:=u.Repo.RemoveValueFrom(Id,value)
	if res.Error != nil {
		utils.Resthis("Error for remove value",res.Error,nil)
	}
	return utils.Resthis("Value removed successfully",nil,res.Data)
}
func (u UserService) AddValueFor(Id string,value float64) utils.Response{
	res:=u.Repo.AddValueFor(Id,value)
	if res.Error != nil {
		utils.Resthis("Error for add value",res.Error,nil)
	}
	return utils.Resthis("Value added successfully",nil,res.Data)
}
