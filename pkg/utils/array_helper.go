package utils

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func ContainsNumber[T Number](dataArr []T, data T) bool {
	for _, val := range dataArr {
		if data == val {
			return true
		}
	}
	return false
}
