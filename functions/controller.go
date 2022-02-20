package functions

import (
	"log"
	"time"
)

func Problem1() {
	start := time.Now()
	res := SumMultiples(1000) // SumMultiples (3 and 5): 233168 | elapsed: 33.29µs
	elapsed := time.Since(start)
	log.Printf("SumMultiples (3 and 5): %d | elapsed: %s", res, elapsed)

	start = time.Now()
	res = SumMultiplesOptimized(1000) // SumMultiplesOptimized (3 and 5): 233168 | elapse: 2.677µs
	elapsed = time.Since(start)
	log.Printf("SumMultiplesOptimized (3 and 5): %d | elapse: %s", res, elapsed)
}
