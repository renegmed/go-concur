//
// Pipeline Design Pattern
//  Create a concurrent structure
//  Exploit the parallelism of multicore machines

package pipeline

// LaunchPipeline generates ouptput
func LaunchPipeline(amount int) int {
	// firstCh := generator(amount)
	// secondCh := power(firstCh)
	// thirdCh := sum(secondCh)

	// result := <-thirdCh

	// return result

	// can also be written like this:
	return <-sum(power(generator(amount)))

	// or can be writter like a future example
	// .generator(amount).power().sum() HOW?
}

func generator(max int) <-chan int {
	outChInt := make(chan int, 100)

	go func() {
		for i := 1; i <= max; i++ {
			outChInt <- i
		}
		close(outChInt)
	}()

	return outChInt
}

func power(in <-chan int) <-chan int {
	out := make(chan int, 100)

	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()

	return out
}

func sum(in <-chan int) <-chan int {
	out := make(chan int, 100)

	go func() {
		var sum int

		for v := range in {
			sum += v
		}

		out <- sum
		close(out)
	}()

	return out
}
