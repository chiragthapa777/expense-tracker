package utils

func If[T any](cond bool, a T, b T) T {
	if cond {
		return a
	} else {
		return b
	}
}

func NilSafeString(data *string) string {
	if data == nil {
		return ""
	}
	return *data
}
