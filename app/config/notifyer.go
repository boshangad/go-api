package config

type Notifyer interface {
	Callback(*Viper)
}
