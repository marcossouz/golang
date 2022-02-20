package main

import (
	"example.com/m/functions"
	"log"
	"time"
)

func main() {
	start := time.Now()
	res := functions.SumMultiples(1000) // SumMultiples (3 and 5): 233168 | elapsed: 49.672µs
	elapsed := time.Since(start)
	log.Printf("SumMultiples (3 and 5): %d | elapsed: %s", res, elapsed)

	start = time.Now()
	res = functions.SumMultiplesOptimized(1000) // SumMultiplesOptimized (3 and 5): 233168 | elapse: 3.573µs
	elapsed = time.Since(start)
	log.Printf("SumMultiplesOptimized (3 and 5): %d | elapse: %s", res, elapsed)
}
