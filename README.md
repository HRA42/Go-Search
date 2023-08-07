# Go Search
## Overview

This Repo shows different search algorithms implemented in Golang.  
Every search algorithm is packed into its own folder and loaded into main. I also created a benchmark file.

The main package does implement call for every package that in this module:

1. create an Array of 100000 random integers between 0 and 9223372036854775807 (`math.MaxInt`)
2. run linear search and measure the time it takes to search through the array
3. sort the array and run binary search and measure the time it takes to search through array
4. sort the array and run binary search with concurrency and measure the time it takes to sort the array
5. sort the array and run interpolation search and measure the time it takes to search through array
6. sort the array with concurrency and run interpolation search and measure the time it takes to search through array

---
## Benchmark and Unit Tests

### Benchmark
The benchmark file `tests/main_test.go` does the same, but run with the input cases:

> [!WARNING]  
> For some sorting algorithms It's impossible to get a successful run (e.g. Bubble Sort)!

1. 10
2. 100
3. 1000
4. 10000
5. 100000
6. 1000000
7. 10000000
8. 100000000

---

### Unit Tests
Under the intense scrutiny of unit tests, each search algorithm unveils its intricate dance. 
The way they align, correct and reorder data – it's like a performer on a stage,
demonstrating their meticulously rehearsed act. Each test is a critical audience, 
assessing the show and ensuring that the performance is flawless. 
Change the routine? The unit tests won't miss any missteps. 
Adapting this practice transforms your routine code into a dynamic spectacle,
elegantly flowing in the theater of your software. 
In turn, every line of code is now part of a thrilling performance, 
each playing a pivotal role in keeping the audience, which is your software system, captivated.

---

## Codesandbox

You can run this Repo in Codesandbox!  
  
[![Edit in CodeSandbox](https://assets.codesandbox.io/github/button-edit-lime.svg)](https://codesandbox.io/p/github/HRA42/go-search)

---

## Commands to run  

### run main.go

```bash
go run cmd/go-search/main.go
```

### run unit test

```bash
go test -v tests/main_test.go
```

### run benchmark

```bash
go test -bench=. tests/main_test.go
```