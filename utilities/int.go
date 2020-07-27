package utilities

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func IsNegativeInt(x int) bool {
	if x < 0 {
		return true
	}

	return false
}
