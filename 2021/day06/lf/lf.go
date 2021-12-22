package lf

func OneDay(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for lf := range in {
			if lf == 0 {
				lf = 6
				out <- 8
			} else {
				lf -= 1
			}
			out <- lf
		}
	}()
	return out
}

func NDays(n int, in <-chan int) <-chan int {
	for n > 0 {
		in = OneDay(in)
		n -= 1
	}
	return in
}

func NDaysStatic(n int, in <-chan int) int {
	buckets := make([]int, 9)
	for f := range in {
		buckets[f] += 1
	}
	for n > 0 {
		newBuckets := make([]int, 9)
		copy(newBuckets, buckets[1:])
		newBuckets[6] += buckets[0]
		newBuckets[8] = buckets[0]
		buckets = newBuckets
		n -= 1
	}
	count := 0
	for _, b := range buckets {
		count += b
	}
	return count
}
