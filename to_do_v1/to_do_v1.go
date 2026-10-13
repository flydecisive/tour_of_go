// to-do лист без файлов, работа со структурами

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("To-DO лист. Для вывода доступных комманд введите help.")

	for {
		fmt.Print("Введите команду: ")

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка ввода. Попробуйте снова.")
			}
		}

		input := scanner.Text()	
		// fmt.Println(input)

		if input != "" {
			fields := strings.Fields(input)
			command := fields[0]
			
			switch(command) {
				case "help":
					fmt.Println("help")
				case "list":
					fmt.Println("Список задач:")
				case "exit": 
					fmt.Println("Выход из программы.")
					os.Exit(0)
				case "add":
					fmt.Println("Задача добавлена.")
				case "done":
					fmt.Println("Задача выполнена.")
				case "delete":
					fmt.Println("Задача удалена.")
				default: 
					fmt.Println("Введена неверная комманда. Попробуйте еще раз.")
			}
		} else {
			fmt.Println("Команда отсутствует. Попробуйте еще раз.")
		}
	}
}