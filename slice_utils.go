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

func Int64SliceContains(s []int64, data int64) bool {
	for _, item := range s {
		if item == data {
			return true
		}
	}
	return false
}

func Int64SliceIndexFirst(s []int64, data int64) int {
	for i, item := range s {
		if item == data {
			return i
		}
	}
	return -1
}

func Int64SliceRemoveFirst(s []int64, data int64) []int64 {
	i := Int64SliceIndexFirst(s, data)
	if i == -1 {
		return s
	}
	return append(s[:i], s[i+1:]...)
}
func Int64SliceRemoveSlice(s []int64, datas []int64) []int64 {
	result := s
	for _, data := range datas {
		result = Int64SliceRemoveFirst(result, data)
	}
	return result
}

func CloneSliceInt64(datas []int64) []int64 {
	s := make([]int64, 0)
	for _, data := range datas {
		s = append(s, data)
	}
	return s
}

func StringSliceContains(datas []string, s string) bool {
	for _, data := range datas {
		if s == data {
			return true
		}
	}
	return false
}
