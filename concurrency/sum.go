package concurrency

const concurrency = 32

// Sum calculates the sum of all the numbers provided
func Sum(numbers []int) int {
	sums := make(chan int, concurrency)

	for i := 0; i < concurrency; i++ {
		go func(initial int) {
			sum := 0

			for j := initial; j < len(numbers); j += concurrency {
				sum += numbers[j]
			}

			sums <- sum
		}(i)
	}

	finalSum := 0
	for i := 0; i < concurrency; i++ {
		finalSum += <-sums
	}

	return finalSum
}
