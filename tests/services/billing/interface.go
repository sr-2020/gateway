package billing

type Service interface {
	Check() bool
}
