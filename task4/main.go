package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var numbers [10]int
	for i := range numbers {
		numbers[i] = rand.Intn(100) + 1
	}

	sum := 0
	minimum := numbers[0]
	maximum := numbers[0]

	for _, number := range numbers {
		sum += number
		if number < minimum {
			minimum = number
		}
		if number > maximum {
			maximum = number
		}
	}

	average := float64(sum) / float64(len(numbers))

	fmt.Println("Массив:", numbers)
	fmt.Println("Сумма:", sum)
	fmt.Println("Минимум:", minimum)
	fmt.Println("Максимум:", maximum)
	fmt.Printf("Среднее значение: %.2f\n", average)
}
