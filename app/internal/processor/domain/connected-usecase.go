package domain

type ConnectedUsecase interface {
	Execute(event interface{}, chargerParam string) error
}
