package contract

type IRouter interface {
	Boot()
	Load()
	Register()
	Start() error
	Close()
}
