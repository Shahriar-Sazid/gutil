package gutil

func Filter[T any](slice []T, f func(T) bool) (res []T) {
	for _, v := range slice {
		if f(v) {
			res = append(res, v)
		}
	}
	return
}

func FilterWithRef[T any](slice []T, f func(*T) bool) (res []*T) {
	for i := 0; i < len(slice); i++ {
		el := &slice[i]
		if f(el) {
			res = append(res, el)
		}
	}

	return
}

func Map[T any, V any](slice []T, f func(T) V) []V {
	res := make([]V, len(slice))
	for _, v := range slice {
		res = append(res, f(v))
	}
	return res
}

func Reduce[T any, V any](slice []T, f func(V, T, int) V, init V) (res V) {
	res = init
	for idx, v := range slice {
		res = f(res, v, idx)
	}
	return
}

func Find[T any](slice []T, f func(*T) bool) (res *T) {
	for i := 0; i < len(slice); i++ {
		el := &slice[i]
		if f(el) {
			return el
		}
	}
	return
}

func FindValue[T any](slice []T, f func(T) bool) (res T, found bool) {
	for _, v := range slice {
		if f(v) {
			return v, true
		}
	}
	return res, false
}

func IsZero[T comparable](v T) bool {
	var zero T
	return v == zero
}

func Coalesce[T comparable](a T, b ...T) (res T) {
	if !IsZero(a) {
		return a
	} else {
		for _, v := range b {
			if !IsZero(v) {
				return v
			}
		}
	}
	return
}

func Default[T comparable](a T, b T) T {
	return Coalesce(a, b)
}
