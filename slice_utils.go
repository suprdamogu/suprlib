package suprlib

func UIntSliceContains(s []uint, data uint) bool {
	for _, item := range s {
		if item == data {
			return true
		}
	}
	return false
}

func UIntSliceIndexFirst(s []uint, data uint) int {
	for i, item := range s {
		if item == data {
			return i
		}
	}
	return -1
}

func UIntSliceRemoveFirst(s []uint, data uint) []uint {
	i := UIntSliceIndexFirst(s, data)
	if i == -1 {
		return s
	}
	return append(s[:i], s[i+1:]...)
}
