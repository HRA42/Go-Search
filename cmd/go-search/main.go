package main

import (
	"fmt"
	"github.com/hra42/go-search/internal/binarysearch"
	"github.com/hra42/go-search/internal/interpolationsearch"
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
	arr = sortarray.Sort(arr)
	resBin := binarysearch.Search(arr, target)
	durBin := time.Since(startBin)
	if resBin {
		fmt.Println("Target: ", target, " found!")
	} else {
		fmt.Println("Target:", target, "not found!")
	}
	fmt.Println("Binary Search took", durBin)

	// concurrent binary search
	startConBin := time.Now()
	arr = sortarray.ConcurrentSort(arr)
	resConBin := binarysearch.ConcurrentSearch(arr, target)
	durConBin := time.Since(startConBin)
	if resConBin {
		fmt.Println("Target: ", target, " found!")
	} else {
		fmt.Println("Target:", target, "not found!")
	}
	fmt.Println("Concurrent Binary Search took", durConBin)

	// interpolation search
	startInt := time.Now()
	sortarray.Sort(arr)
	resInt := interpolationsearch.Search(arr, target)
	durInt := time.Since(startInt)
	if resInt {
		fmt.Println("Target: ", target, " found!")
	} else {
		fmt.Println("Target:", target, "not found!")
	}
	fmt.Println("Interpolation Search took", durInt)

	// interpolation search with concurrent sorting
	startConInt := time.Now()
	arr = sortarray.ConcurrentSort(arr)
	resConInt := interpolationsearch.Search(arr, target)
	durConInt := time.Since(startConInt)
	if resConInt {
		fmt.Println("Target: ", target, " found!")
	} else {
		fmt.Println("Target:", target, "not found!")
	}
	fmt.Println("Concurrent Interpolation Search took", durConInt)
}
