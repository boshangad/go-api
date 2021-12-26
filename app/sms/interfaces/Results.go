package interfaces

type Result interface {
	error
	IsSuccess() bool
	Code() string
}
