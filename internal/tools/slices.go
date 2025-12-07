package tools

type Number interface {
	int8 | int16 | int | int32 | int64 | float32 | float64 | uint8 | uint16 | uint32 | uint64
}

func Pop[T Number](slice []T) ([]T, T) {
	if len(slice) == 0 {
		return slice, 0 // Edge case: пустой срез
	}
	return slice[:len(slice)-1], slice[len(slice)-1]
}
