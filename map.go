package gutil

func Keys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}

	return r
}

func Values[K comparable, V any](m map[K]V) []V {
	r := make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}

	return r
}

func ToMap[T any, V comparable](slice []T, key func(*T) V) map[V]*T {
	res := make(map[V]*T, len(slice))

	for i := 0; i < len(slice); i++ {
		el := &slice[i]
		res[key(el)] = el
	}

	return res
}

func ToMapWithValue[T any, V comparable](slice []T, key func(T) V) map[V]T {
	res := make(map[V]T, len(slice))

	for _, v := range slice {
		res[key(v)] = v
	}

	return res
}
