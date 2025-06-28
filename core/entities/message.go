package entities

import "github.com/Nicolas-Asafe/BankingSystem/utils"

type Message struct {
	AuthoId     string
	ForId       string
	Id          string
	Title       string
	Description string
	Data        string
	Log         utils.Logs
}

type MessageReceived struct {
	AuthoId     string
	Id          string
	Title       string
	Description string
	Data        string
	Log         utils.Logs
}

func (m *Message) View() utils.Response {
	dataForView := MessageReceived{
		AuthoId:     m.AuthoId,
		Id:          m.Id,
		Title:       m.Title,
		Description: m.Description,
		Data:        m.Data,
	}
	return utils.Resthis("Message was seen", nil, dataForView)
}
