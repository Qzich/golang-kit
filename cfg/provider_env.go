package cfg

type Provider interface {
	Provide() map[string]string
}