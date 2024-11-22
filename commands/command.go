package commands

type Executor interface {
	Output() ([]byte, error)
}
