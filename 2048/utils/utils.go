package utils

func RemoveSliceByIndex(s []int, i int) []int {
	if i > 0 && i < len(s) {
		return append(s[:i], s[i+1:]...)
	} else {
		return []int{}
	}
}
