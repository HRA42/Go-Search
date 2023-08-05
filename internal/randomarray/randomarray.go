package randomarray

import (
	"math/rand"
	"time"
)

func CreateRandomIntArray(length int, max int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // create local pseudorandom number generator
	res := make([]int, length)

	for i := range res {
		res[i] = r.Intn(max) // uses local prng
	}

	return res
}
