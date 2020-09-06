package push

type Service interface {
	Check() bool
}
