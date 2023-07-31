package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/hra42/go-search/linearsearch"
	"github.com/hra42/go-search/randomarray"
	"github.com/hra42/go-search/sortarray"
)

func main() {
	// initial setup
	const max = math.MaxInt
	arr := randomarray.Create_random_int_array(100000000, max)
	target := rand.Intn(max)
	fmt.Println(len(arr))

	// linear search
	start_lin := time.Now()
	res_lin := linearsearch.Search(arr, target)
	dur_lin := time.Since(start_lin)
	if res_lin {
		fmt.Println("Target: ", target, " found!")
	} else {
		fmt.Println("Target:", target, "not found!")
	}
	fmt.Println("Linear Search took", dur_lin)

	// binary search
	start_bin := time.Now()
	sortarray.Sort(arr)
	res_bin := linearsearch.Search(arr, target)
	dur_bin := time.Since(start_bin)
	if res_bin {
		fmt.Println("Target: ", target, " found!")
	} else {
		fmt.Println("Target:", target, "not found!")
	}
	fmt.Println("Binary Search took", dur_bin)
}
