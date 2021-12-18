package interfaces

type Result interface {
	error
	Code() int
}
