package loops

func simpleLoop() {
	// Simple loop
	for i := 0; i < 5; i++ {
		println(i)
	}
}

func whileLoop() {
	// While loop
	i := 0
	for i < 5 {
		println(i)
		i++
	}
}

func infiniteLoop() {
	// Infinite loop
	for {
		println("Infinite loop")
	}
}

func rangeLoop() {
	// Range loop
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	println("sum:", sum)
}

func rangeLoopWithItrater() {
	// Range loop with iterator
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		println(k, v)
	}
}

func breakLoop() {
	// Break loop
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		println(i)
	}
}

func continueLoop() {
	// Continue loop
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		println(i)
	}
}

func nestedLoop() {
	// Nested loop
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			println(i, j)
		}
	}
}

func Loop() {
	simpleLoop()
	whileLoop()
	infiniteLoop()
	rangeLoop()
	rangeLoopWithItrater()
	breakLoop()
	continueLoop()
	nestedLoop()
}
