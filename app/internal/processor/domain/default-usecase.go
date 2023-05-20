package domain

type DefaultUsecase interface {
	Execute(event interface{}) error
}
