package utils

func Of[T any](t T) *T {
	return &t
}

func Deref[T any](t *T) T {
	if t == nil {
		var d T
		return d
	}
	return *t
}
