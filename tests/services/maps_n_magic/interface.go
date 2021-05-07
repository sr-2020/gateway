package maps_n_magic

type Service interface {
	Check() bool
	FileList() bool
}
