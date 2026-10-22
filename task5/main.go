package main

import (
	"fmt"
	"math/rand"
)

func main() {
	secret := rand.Intn(100) + 1
	fmt.Println("Загадано число от 1 до 100")

	for {
		var guess int
		fmt.Print("Ваше число: ")
		if _, err := fmt.Scan(&guess); err != nil {
			fmt.Println("Ошибка ввода числа")
			return
		}

		switch {
		case guess < secret:
			fmt.Println("Больше")
		case guess > secret:
			fmt.Println("Меньше")
		default:
			fmt.Println("число угадано")
			return
		}
	}
}
