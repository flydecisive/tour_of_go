// сделать примеры логический операцией с функцией для вывода и попробовать функцию вынести в пакет отдельный
package main

import (
	"fmt"
	pt "logical_operations/print_tools"
)

func main() {
	var (
		operation string;
		num1 int;
		num2 int
	)

	fmt.Println("Программа позволяет посчитать значения двух целых чисел при выполнении логический операций.");
	fmt.Println("Доступные логические операции: &, |, &^")
	fmt.Println("Операция exit позволяет выйти из программы.")
	fmt.Print("\n")

	for {
		fmt.Print("Введите операцию: ")
		// _, errOperation := 
		fmt.Scan(&operation)

		// if errOperation != nil {
		// 	fmt.Println("Введена неверная операция. Поробуйте снова.")
		// 	continue
		// } 

		switch operation {
		case "&", "|", "&^", "^":
			// всё ок, идём дальше
		case "exit":
			fmt.Println("Выход из программы.")
			return
		default:
			fmt.Println("Введена неверная операция. Попробуйте снова.")
			continue
		}

		fmt.Print("Введите первый операнд: ")
		_, errNum1 := fmt.Scan(&num1)

		if errNum1 != nil {
			fmt.Println("Допускается вводить только целые числа. Попробуйте снова.")
			continue
		}

		fmt.Print("Введите второй операнд: ")
		_, errNum2 := fmt.Scan(&num2)

		if errNum2 != nil {
			fmt.Println("Допускается вводить только целые числа. Попробуйте снова.")
			continue
		}

		pt.PrintLogicalOperation(operation, num1, num2)
		fmt.Println()
	}
}