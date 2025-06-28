package interfaces

import "github.com/Nicolas-Asafe/BankingSystem/utils"

type EconomyRepository interface {
	findOne(AuthoId string, EcoId string) utils.Response
	find(AuthoId string) utils.Response
}
