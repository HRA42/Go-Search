package randomarray

import (
	"math/rand"
	"unsafe"
)

func Create_random_int_array(length int, max int) (res []int) {
	res = make([]int, length)
	bts := unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(res))), len(res)*int(unsafe.Sizeof(int(0))))
	rand.Read(bts)
	for i, v := range res {
		res[i] = v % max
	}
	return
}
