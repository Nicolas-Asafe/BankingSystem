package interfaces

import (
	"github.com/Nicolas-Asafe/BankingSystem/core/entities"
	"github.com/Nicolas-Asafe/BankingSystem/utils"
)

type UserRepository interface{
	NewUser(user entities.User) utils.Response
	RemoveUser(Id int) utils.Response
	FindUser(Id int) utils.Response
	FindUsers() utils.Response
	PutUserName(Id int,Name string) utils.Response
}