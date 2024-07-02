package http

func castToUint(old []int) []uint {
	result := make([]uint, len(old))

	for i, v := range old {
		result[i] = uint(v)
	}

	return result
}
