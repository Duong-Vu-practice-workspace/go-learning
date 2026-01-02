package main

func getLast[T any](s []T) T {
	// ?
	n := len(s)
	if n == 0 {
		var zero T
		return zero
	}
	return s[n-1]
}
