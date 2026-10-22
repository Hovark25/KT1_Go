package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите строку: ")
	if !scanner.Scan() {
		fmt.Println("Ошибка ввода строки")
		return
	}

	// Руны
	characters := []rune(scanner.Text())
	isPalindrome := true

	for i := 0; i < len(characters)/2; i++ {
		if characters[i] != characters[len(characters)-1-i] {
			isPalindrome = false
			break
		}
	}

	if isPalindrome {
		fmt.Println("Строка является палиндромом")
	} else {
		fmt.Println("Строка не является палиндромом")
	}
}
