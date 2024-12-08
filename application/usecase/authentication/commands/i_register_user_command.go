package commands

type IRegisterUserCommand interface {
	Register() (string, error)
}
