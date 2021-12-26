package interfaces

type Gateway interface {
	Name() string
	Send(to PhoneNumber, message Message, config Config) Result
}
