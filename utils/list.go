package utils

// Unique - remove duplicates
func Unique(arr []int) []int {
	keysMap := map[int]struct{}{}
	for _, v := range arr {
		keysMap[v] = struct{}{}
	}
	keys := make([]int, len(keysMap))
	i := 0
	for k := range keysMap {
		keys[i] = k
		i++
	}
	return keys
}
