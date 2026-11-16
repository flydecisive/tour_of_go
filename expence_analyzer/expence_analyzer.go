package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
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

		if !scanner.Scan() {
			if err:= scanner.Err(); err != nil {
				fmt.Println("Ошибка ввода. Попробуйте еще раз.")
			}
		}

		input := scanner.Text()
		
		fmt.Println("input " + input)
	}
	
}