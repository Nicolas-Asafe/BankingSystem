package services

import "github.com/Nicolas-Asafe/BankingSystem/core/interfaces"
import "github.com/Nicolas-Asafe/BankingSystem/core/entities"
import "github.com/Nicolas-Asafe/BankingSystem/utils"

type MessageService struct{
	Repo interfaces.MessageRepository
}

func (m MessageRepository) SendMessage(msg entities.Message){
	res := m.Repo.Send(msg)
	if res.Error != nil {
		utils.Resthis("Error sending message",res.Error,nil)
		return
	}
	utils.Resthis("Message sent successfully",nil,)
}