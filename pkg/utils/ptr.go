package utils

func Of[T any](t T) *T {
	return &t
}
