package main

import (
	"github.com/hra42/go-search/internal/binarysearch"
	"github.com/hra42/go-search/internal/linearsearch"
	"github.com/hra42/go-search/internal/sortarray"
	"testing"
)

var data = []int{5, 8, 9, 2, 3, 7, 4, 1, 6}
var target = 7

func TestLinearSearch(t *testing.T) {
	result := linearsearch.Search(data, target)
	if !result {
		t.Errorf("Expected to find the target %v, but it was not found", target)
	}
}

func TestBinarySearch(t *testing.T) {
	data = sortarray.Sort(data)
	result := binarysearch.Search(data, target)
	if !result {
		t.Errorf("Expected to find the target %v, but it was not found", target)
	}
}
