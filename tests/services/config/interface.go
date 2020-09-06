package config

type Service interface {
	Check() bool
	Read(string) (string, error)
	Write(string, string) error
}
