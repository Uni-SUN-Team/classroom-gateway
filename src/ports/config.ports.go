package ports

type Config interface {
	ConfigENV() error
}
