package main

import (
	"fmt"
	"github.com/hra42/go-search/internal/binarysearch"
	"github.com/hra42/go-search/internal/interpolationsearch"
	"github.com/hra42/go-search/internal/linearsearch"
	"github.com/hra42/go-search/internal/randomarray"
	"github.com/hra42/go-search/internal/sortarray"
	"math"
	"testing"
	"time"
)

var testData = []int{5, 8, 9, 2, 3, 7, 4, 1, 6}
var testTarget = 7

var benchmarkCases = []struct {
	input int
}{
	{input: 10},
	{input: 100},
	{input: 1000},
	{input: 10000},
	{input: 100000},
	{input: 1000000},
	{input: 10000000},
	{input: 100000000},
}

func TestLinearSearch(t *testing.T) {
	result := linearsearch.Search(testData, testTarget)
	if !result {
		t.Errorf("Expected to find the target %v, but it was not found", testTarget)
	}
}

func TestBinarySearch(t *testing.T) {
	testData = sortarray.Sort(testData)
	result := binarysearch.Search(testData, testTarget)
	if !result {
		t.Errorf("Expected to find the target %v, but it was not found", testTarget)
	}
}

func TestConcurrentBinarySearch(t *testing.T) {
	testData = sortarray.ConcurrentSort(testData)
	result := binarysearch.ConcurrentSearch(testData, testTarget)
	if !result {
		t.Errorf("Expected to find the target %v, but it was not found", testTarget)
	}
}

func TestInterpolationSearch(t *testing.T) {
	testData = sortarray.Sort(testData)
	result := interpolationsearch.Search(testData, testTarget)
	if !result {
		t.Errorf("Expected to find the target %v, but it was not found", testTarget)
	}
}

func BenchmarkLinearSearch(b *testing.B) {
	for _, v := range benchmarkCases {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			var duration time.Duration
			for i := 0; i < b.N; i++ {
				start := time.Now()
				arr := randomarray.CreateRandomIntArray(v.input, math.MaxInt)
				linearsearch.Search(arr, testTarget)
				duration = time.Since(start)
			}
			fmt.Printf("Für die Eingabe: %v wurde %v benötigt\n", v.input, duration)
		})
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	for _, v := range benchmarkCases {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			var duration time.Duration
			for i := 0; i < b.N; i++ {
				start := time.Now()
				arr := randomarray.CreateRandomIntArray(v.input, math.MaxInt)
				arr = sortarray.Sort(arr)
				binarysearch.Search(arr, testTarget)
				duration = time.Since(start)
			}
			fmt.Printf("Für die Eingabe: %v wurde %v benötigt\n", v.input, duration)
		})
	}
}

// implemented binary search and sorting concurrent
func BenchmarkConcurrentBinarySearch(b *testing.B) {
	for _, v := range benchmarkCases {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			var duration time.Duration
			for i := 0; i < b.N; i++ {
				start := time.Now()
				arr := randomarray.CreateRandomIntArray(v.input, math.MaxInt)
				arr = sortarray.ConcurrentSort(arr)
				binarysearch.ConcurrentSearch(arr, testTarget)
				duration = time.Since(start)
			}
			fmt.Printf("Für die Eingabe: %v wurde %v benötigt\n", v.input, duration)
		})
	}
}

func BenchmarkInterpolationSearch(b *testing.B) {
	for _, v := range benchmarkCases {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			var duration time.Duration
			for i := 0; i < b.N; i++ {
				start := time.Now()
				arr := randomarray.CreateRandomIntArray(v.input, math.MaxInt)
				arr = sortarray.Sort(arr)
				interpolationsearch.Search(arr, testTarget)
				duration = time.Since(start)
			}
			fmt.Printf("Für die Eingabe: %v wurde %v benötigt\n", v.input, duration)
		})
	}
}

// with concurrent sorting
func BenchmarkConcurrentInterpolationSearch(b *testing.B) {
	for _, v := range benchmarkCases {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			var duration time.Duration
			for i := 0; i < b.N; i++ {
				start := time.Now()
				arr := randomarray.CreateRandomIntArray(v.input, math.MaxInt)
				arr = sortarray.ConcurrentSort(arr)
				interpolationsearch.Search(arr, testTarget)
				duration = time.Since(start)
			}
			fmt.Printf("Für die Eingabe: %v wurde %v benötigt\n", v.input, duration)
		})
	}
}
