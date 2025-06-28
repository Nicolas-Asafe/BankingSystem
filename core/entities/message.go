package entities

type Message struct {
	AuthoId string
	ForId string
	id string
	Title string
	Description *string
	Data *any
}