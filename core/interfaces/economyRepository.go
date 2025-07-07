package interfaces

import "github.com/Nicolas-Asafe/BankingSystem/utils"
import "github.com/Nicolas-Asafe/BankingSystem/core/entities"

type EconomyRepository interface {
	findOne(AuthoId string, EcoId string) utils.Response
	find(AuthoId string) utils.Response
	newEconomy(Economy entities.Economy) utils.Response
	removeEconomy(Id string) utils.Response
	putName(newName string) utils.Response
	AddValue(value float64,Id string) utils.Response
	Withdraw(value float64,Id string) utils.Response
}
