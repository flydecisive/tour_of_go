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
	)
	fmt.Printf("Игра Угадай число. Необходимо угадать загаданное число. Максимум %v попыток.\n", attempts)
	fmt.Printf("Число загадывается в диапазоне от %v до %v\n", minNum, maxNum)
	guessedNum := minNum + rand.IntN(maxNum)

	for i := 1; i <= attempts; i++ {
		var number int
		fmt.Print("\nВведите число: ")

		if _, err := fmt.Scan(&number); err != nil {
			fmt.Println("Ошибка, можно вводить только целые числа.")
			printRemainingAttempts(attempts, i)
			continue
		}

		if number == guessedNum {
			fmt.Printf("\nПоздравляю! Вы угадали! Загаданное число: %v\n", guessedNum)
			fmt.Printf("Потрачено попыток: %v", i)
			break
		} else if number < guessedNum {
			fmt.Println("Загаданное число больше.")
			printRemainingAttempts(attempts, i)
		} else if number > guessedNum {
			fmt.Println("Загаданное число меньше")
			printRemainingAttempts(attempts, i)
		}

		if i == attempts && number != guessedNum {
			fmt.Println("\nВы проиграли. Не осталось попыток.")
			fmt.Printf("Было загадано число: %v", guessedNum)
		}
	}
}

func printRemainingAttempts(allAttempts, expAttempts int) {
	fmt.Printf("Осталось попыток: %v\n", allAttempts - expAttempts)
}