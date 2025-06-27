package entities

type User struct{
	Name string
	Password string
	Email string
	Age int64
	Id string
	Balance float64
	HistoryTransitions []Transition
}
