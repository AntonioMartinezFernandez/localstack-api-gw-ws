package domain

type DisconnectedUsecase interface {
	Execute(event interface{}) error
}
