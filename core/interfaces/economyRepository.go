package interfaces

import "github.com/Nicolas-Asafe/BankingSystem/core/entities"
import "github.com/Nicolas-Asafe/BankingSystem/utils"

type EconomyRepository interface {
	Save(economy entities.Economy) utils.Response
	Remove(id string) utils.Response
	Put(title *string, description *string) utils.Response

}
