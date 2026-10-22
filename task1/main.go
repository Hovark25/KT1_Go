package main

import "fmt"

func main() {
	var first, second float64
	var operation string

	fmt.Print("Первое число: ")
	if _, err := fmt.Scan(&first); err != nil {
		fmt.Println("Ошибка ввода числа")
		return
	}

	fmt.Print("Второе число: ")
	if _, err := fmt.Scan(&second); err != nil {
		fmt.Println("Ошибка ввода числа")
		return
	}

	fmt.Print("Операция (+, -, /, *): ")
	if _, err := fmt.Scan(&operation); err != nil {
		fmt.Println("Ошибка ввода операции")
		return
	}

	var result float64

	switch operation {
	case "+":
		result = first + second
	case "-":
		result = first - second
	case "*":
		result = first * second
	case "/":
		if second == 0 {
			fmt.Println("деление на ноль")
			return
		}
		result = first / second
	default:
		fmt.Println("Неизвестная операция")
		return
	}

	fmt.Printf("Результат: %g\n", result)
}
