package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Expense struct {
	Amount float64
	Category string
}

func main() {
	var expenses []Expense

	for {
		// fmt.Println(commands)
		
		
		input := getUserCommand()

		switch(input) {
			case "1":
				fmt.Print("\nВведите сумму: ")
				amountInput := scan()
				amount, err := strconv.ParseFloat(amountInput, 64)
				if err != nil {
					fmt.Println("Введено не число. Попробуйте еще раз.")
					continue
				}
				if amount < 1 {
					fmt.Print("Расход должен быть больше 0.\n")
					continue
				}
				fmt.Print("Введите категорию: ")
				category := scan()
				expenses = append(expenses, Expense{Amount: amount, Category: category})
				fmt.Println("Расход добавлен.")
			case "2":
				printAmounts(expenses)
			case "3":
				printStatistic()
			case "0": 
				fmt.Println("Выход.")
				return
			default:
				fmt.Println("Введена не верная команда, попробуйте еще раз.")
				continue
		}
	}
	
}

func getUserCommand() string {
	commands := []string{"Выйти", "Добавить расход", "Показать все расходы", "Показать статистику"}

	fmt.Println("\nСписок действий:")
	for i, command := range commands {
		fmt.Printf("\t%v. %v\n", i, command)
	}

	fmt.Print("\nВыберите действие: ")
	input := scan()
	return input
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

func printAmounts(expenses []Expense) {
	if len(expenses) == 0 {
		fmt.Println("Расходов нет.")
		return
	}
	printDivider()
	fmt.Println("Список расходов:")
	for i, expense := range expenses {
		fmt.Printf("\t%v. %v - %.2f руб.\n", i + 1, expense.Category, expense.Amount)
	}
	printDivider()
}

func printStatistic() {
	fmt.Println("Статистика расходов")
}

func printDivider() {
	fmt.Println("\n\t=================")
}

// 2. Просмотр расходов

// Покажи все записи в порядке добавления:
// 1. еда — 450.00 руб.
// 2. транспорт — 100.00 руб.
// 3. еда — 250.00 руб.
// 4. развлечения — 1200.00 руб.
// Если записей нет, выведи: Расходов пока нет.

// 3. Статистика

// Для примера выше:
// Количество расходов: 4
// Общая сумма: 2000.00 руб.
// Средний расход: 500.00 руб.

// По категориям:
// еда: 700.00 руб.
// транспорт: 100.00 руб.
// развлечения: 1200.00 руб.
// Средний расход здесь — общая сумма, делённая на количество записей, а не на количество категорий. При отсутствии записей делить на ноль нельзя — достаточно сообщения об отсутствии расходов.