package helper

type GenericSlice[T any] []T

func (s GenericSlice[T]) FilterOut(excluded func(e T) bool) GenericSlice[T] {
	var result GenericSlice[T]
	for _, e := range s {
		if excluded(e) {
			continue
		}
		result = append(result, e)
	}
	return result
}
