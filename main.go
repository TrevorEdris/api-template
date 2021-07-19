package main

import (
	"github.com/inconshreveable/log15"
)

var (
	log = log15.New("module", "RepoNameHere")
)

func fibonacci(n int) (int, error) {
	if n < 0 {
		return -1, InvalidStartingNumberError
	}

	f := make([]int, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n], nil
}

func main() {
	log.Info("Hello from the repo template!")
	fibNumToGet := 10
	output, err := fibonacci(fibNumToGet)
	if err != nil {
		log.Error("Encountered error with calculating fibonacci number", "error", err, "input", fibNumToGet)
	} else {
		log.Info("Calculated fibonacci value", "input", fibNumToGet, "output", output)
	}
}
