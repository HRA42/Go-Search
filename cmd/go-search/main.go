package main

import (
	"fmt"
	"github.com/hra42/go-search/internal/linearsearch"
	"github.com/hra42/go-search/internal/randomarray"
	"github.com/hra42/go-search/internal/sortarray"
	"math"
	"math/rand"
	"time"
)

func main() {
	// initial setup
	const max = math.MaxInt
	arr := randomarray.CreateRandomIntArray(100000000, max)
	target := rand.Intn(max)
	fmt.Println(len(arr))

	// linear search
	startLin := time.Now()
	resLin := linearsearch.Search(arr, target)
	durLin := time.Since(startLin)
	if resLin {
		fmt.Println("Target: ", target, " found!")
	} else {
		fmt.Println("Target:", target, "not found!")
	}
	fmt.Println("Linear Search took", durLin)

	// binary search
	startBin := time.Now()
	sortarray.Sort(arr)
	resBin := linearsearch.Search(arr, target)
	durBin := time.Since(startBin)
	if resBin {
		fmt.Println("Target: ", target, " found!")
	} else {
		fmt.Println("Target:", target, "not found!")
	}
	fmt.Println("Binary Search took", durBin)
}
