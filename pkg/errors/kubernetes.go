package errors

type KubernetesError interface {
	Error() string
	GetStatus() int
	GetMessage() string
}
