package utils

type Response struct{
	Message string
	Error error
	Data *struct{}
}

func Resthis(Message string,Error error,Data *struct{})Response{
	return Response{Message,Error,Data,}
}