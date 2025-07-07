package interfaces

import (
	"github.com/Nicolas-Asafe/BankingSystem/core/entities"
	"github.com/Nicolas-Asafe/BankingSystem/utils"
)

type MessageRepository interface {
	Send(message entities.Message) utils.Response
	ReciveMessage(message entities.Message) utils.Response
}
