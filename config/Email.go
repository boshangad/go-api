package config

type Email struct {
	Default string
	Gateways map[string]map[string]interface{}
}
