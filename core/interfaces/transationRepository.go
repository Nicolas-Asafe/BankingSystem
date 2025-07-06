package interfaces

import "github.com/Nicolas-Asafe/BankingSystem/utils"
import "github.com/Nicolas-Asafe/BankingSystem/core/entities"

type TransitionRepository interface{
	Send(entities.Transition) utils.Response
	ReciveTransition(tst entities.Transition) utils.Response
}