package interfaces

import (
	"github.com/Nicolas-Asafe/BankingSystem/core/entities"
	"github.com/Nicolas-Asafe/BankingSystem/utils"
)

type UserRepository interface{
	NewUser(user entities.User) utils.Response
	RemoveUser(Id string) utils.Response
	FindUser(Id string) utils.Response
	FindUsers() utils.Response
	PutUserName(Id string,Name string) utils.Response
}