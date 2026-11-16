package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var commands = 
	`
	1. Добавить расход
	2. Показать все расходы
	3. Показать статистику
	0. Выйти
	`

	for {
		fmt.Println(commands)
		fmt.Print("\nВыберите действие: ")

		input := scan()

		switch(input) {
			case "1":
				fmt.Print("\nВведите сумму: ")
				amountInput := scan()
				amount, err := strconv.Atoi(amountInput)
				if err != nil {
					fmt.Println("Введено не число. Попробуйте еще раз.")
					continue
				}
				fmt.Print("Введите категорию: ")
				category := scan()
				// Тут добавлять расход в объект
				fmt.Println("Расход добавлен.")
				fmt.Printf("Расход %v, категория %v\n", amount, category)
			case "2":
				fmt.Println("Показать все расходы")
			case "3":
				fmt.Println("Показать статистику")
			case "0": 
				fmt.Println("Выход.")
				return
			default:
				fmt.Println("Введена не верная команда, попробуйте еще раз.")
				continue
		}
	}
	
}

func scan() string {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		if err:= scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка ввода. Попробуйте еще раз.")
		}
	}

	input := scanner.Text()

	return input
}