package contract

type IService interface {
	Handle(ctx Context) error
}
