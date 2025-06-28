package interfaces

import (
	"github.com/Nicolas-Asafe/BankingSystem/core/entities"
	"github.com/Nicolas-Asafe/BankingSystem/utils"
)

type MessageRepository interface {
	Send(AuthoId string,ForAuthoId string,message entities.Message) utils.Response
	FindOne(id string) utils.Response
	Find() utils.Response
}