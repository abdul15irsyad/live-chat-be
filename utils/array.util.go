package utils

func Includes[T comparable](slice []T, item T) bool {
	for _, value := range slice {
		if value == item {
			return true
		}
	}
	return false
}

func MapSlice[T any, K any](slice []T, mapper func(item T) K) []K {
	result := make([]K, len(slice))
	for i, value := range slice {
		result[i] = mapper(value)
	}
	return result
}

func Values[K comparable, V any](m map[K]V) []V {
	result := make([]V, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}
	return result
}
