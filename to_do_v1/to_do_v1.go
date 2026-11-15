// to-do лист без файлов, работа со структурами

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	id int;
	text string;
	done bool
}

func main() {
	tasks := []Task{}
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

		if input != "" {
			fields := strings.Fields(input)
			command := fields[0]
			
			switch(command) {
				case "help":
					fmt.Print("<--- --->\n" +
					"Список комманд:\n" +
					"help - помощь\n" +
					"add <задача> - добавить задачу\n" + 
					"done <задача> - завершить задачу\n" +
					"delete <задача> - удалить задачу\n" +
					"exit - выход\n" +
					"<--- --->\n")
				case "list":
					fmt.Print("Список задач:")
					fmt.Println(tasks)
				case "exit": 
					fmt.Println("Выход из программы.")
					os.Exit(0)
				case "add":
					task := Task{id: len(tasks), text: strings.Join(fields[1:], " "), done: false}
					tasks = append(tasks, task)
					fmt.Printf("<--- --->\nДобавлена задача: %s\n<--- --->\n", strings.Join(fields[1:], " "))
					fmt.Println(tasks[0].id)
				case "done":
					fmt.Printf("<--- --->\nВыполнена задача: %s\n<--- --->\n", strings.Join(fields[1:], " "))
					// Тут сразу выводить обновленный список задач
					// Выполненная задача отображается перечеркнутой
				case "delete":
					fmt.Printf("<--- --->\nУдалена задача: %s\n<--- --->\n", strings.Join(fields[1:], " "))
				default: 
					fmt.Println("Введена неверная комманда. Попробуйте еще раз.")
			}
		} else {
			fmt.Println("Команда отсутствует. Попробуйте еще раз.")
		}
	}
}

func printTasks(tasks []Task) {
	// Вывод листа с задачами
	
}