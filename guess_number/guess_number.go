package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	attempts := 10
	var (
		minNum int = 1
		maxNum int = 200
		number int
	)
	fmt.Printf("Игра Угадай число. Необходимо угадать загаданное число. Максимум %v попыток.\n", attempts)
	fmt.Printf("Число загадывается в диапазоне от %v до %v\n", minNum, maxNum)
	guessedNum := minNum + rand.IntN(maxNum-minNum+1)

	for i := 1; i <= attempts; i++ { 
		fmt.Print("\nВведите число: ")

		if _, err := fmt.Scan(&number); err != nil {
			fmt.Println("Ошибка, можно вводить только целые числа.")
			printRemainingAttempts(attempts, i)
			continue
		}

		if number == guessedNum {
			fmt.Printf("\nПоздравляю! Вы угадали! Загаданное число: %v\n", guessedNum)
			fmt.Printf("Потрачено попыток: %v\n", i)
			return
		}

		if number < guessedNum {
			fmt.Println("Загаданное число больше.")
			printRemainingAttempts(attempts, i)
		} else if number > guessedNum {
			fmt.Println("Загаданное число меньше")
			printRemainingAttempts(attempts, i)
		}
	}

	fmt.Println("\nВы проиграли. Не осталось попыток.")
	fmt.Printf("Было загадано число: %v\n", guessedNum)
}

func printRemainingAttempts(allAttempts, expAttempts int) {
	fmt.Printf("Осталось попыток: %v\n", allAttempts - expAttempts)
}