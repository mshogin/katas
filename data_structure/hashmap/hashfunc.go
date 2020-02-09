package hashmap

import "math"

type hashFunk func(key string) int

var defaultHashFunc = func(key string) int {
	return int(
		math.Mod(float64(key[0]), 10),
	)
}
