package gmap

func Contains[T any](m map[string]T, key string) bool {
	_, ok := m[key]
	return ok
}
