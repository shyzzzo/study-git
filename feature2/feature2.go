package feature2

import "fmt"

func add(a, b int) int {
	sum := a + b
	fmt.Println(sum)

	return sum
}

func Feature2() {
	fmt.Println("Я рома стрелецкий лучший программист")
	add(1, 5)
}
