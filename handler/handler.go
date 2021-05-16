package handler

type Handler interface {
	Handle([]string) error
}