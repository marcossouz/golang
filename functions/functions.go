package functions

/*
 * Problem 1 - Multiples of 3 or 5
 * If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
 * Find the sum of all the multiples of 3 or 5 below 1000.
 */

func Multiples(n int) []int {
	var mult []int
	for i := 1; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			mult = append(mult, i)
		}
	}
	return mult
}

func SumMultiples(N int) int {
	var sum int
	mult := Multiples(N)
	for i := 0; i < len(mult); i++ {
		sum += mult[i]
	}
	return sum
}
