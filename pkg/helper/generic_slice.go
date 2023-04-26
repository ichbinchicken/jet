package helper

type GenericSlice[K any] []K

func (s GenericSlice[K]) FilterOut(excluded func(e K) bool) GenericSlice[K] {
	var result GenericSlice[K]
	for _, e := range s {
		if excluded(e) {
			continue
		}
		result = append(result, e)
	}
	return result
}

func MapSlice[K, V any](s GenericSlice[K], mapper func(e K) V) GenericSlice[V] {
	result := make(GenericSlice[V], len(s))
	for i, e := range s {
		result[i] = mapper(e)
	}
	return result
}
