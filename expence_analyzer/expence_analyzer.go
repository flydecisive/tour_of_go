package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Expense struct {
	Amount float64
	Category string
}

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	var expenses []Expense
	
	for {
		input, ok := getUserCommand()
		if !ok {
    		return
		}

		switch(input) {
			case "1":
				fmt.Print("\nВведите сумму: ")
				amountInput, ok := scan(scanner)
				if !ok {
					return
				}
				amount, err := strconv.ParseFloat(amountInput, 64)
				if err != nil {
					fmt.Println("Введено не число. Попробуйте еще раз.")
					continue
				}
				if math.IsNaN(amount) || math.IsInf(amount, 0) || amount <= 0 {
					fmt.Print("Расход должен быть больше 0.\n")
					continue
				}
				fmt.Print("Введите категорию: ")
				category, ok := scan(scanner)
				if !ok {
					return
				}
				category = strings.TrimSpace(category)
				if category == "" {
					fmt.Println("Категория не должна быть пустой.")
					continue
				}
				expenses = append(expenses, Expense{Amount: amount, Category: category})
				fmt.Println("Расход добавлен.")
			case "2":
				printAmounts(expenses)
			case "3":
				printStatistic(expenses)
			case "0": 
				fmt.Println("Выход.")
				return
			default:
				fmt.Println("Введена не верная команда, попробуйте еще раз.")
				continue
		}
	}
	
}

func getUserCommand() (string, bool) {
	commands := []string{"Выйти", "Добавить расход", "Показать все расходы", "Показать статистику"}

	fmt.Println("\nСписок действий:")
	for i, command := range commands {
		fmt.Printf("\t%v. %v\n", i, command)
	}

	fmt.Print("\nВыберите действие: ")

	return scan(scanner)
}

func scan(scanner *bufio.Scanner) (string, bool) {
	if !scanner.Scan() {
		if err:= scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка ввода. Попробуйте еще раз.")
		}

		return "", false
	}

	return scanner.Text(), true
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

func printStatistic(expenses []Expense) {
	var expensesCount = float64(len(expenses))
	var expensesTotalSum float64 = 0
	var expensesMid float64
	var categories []string

	if len(expenses) == 0 {
		fmt.Println("Расходов нет.")
		return
	}

	for _, expense := range expenses {
		expensesTotalSum += expense.Amount
		if !slices.Contains(categories, expense.Category) {
			categories = append(categories, expense.Category)
		}
	}
	expensesMid = expensesTotalSum / expensesCount
	printDivider()
	fmt.Println("Статистика расходов:")
	fmt.Printf("\tКоличество расходов: %v\n", expensesCount)
	fmt.Printf("\tОбщая сумма: %.2f руб.\n", expensesTotalSum)
	fmt.Printf("\tСредний расход: %.2f руб.\n", expensesMid)

	fmt.Println("\nПо категориям:")
	for _, category := range categories {
		sumForCategory := 0.0
		for _, expense := range expenses {
			if expense.Category == category {
				sumForCategory += expense.Amount
			}
		}
		fmt.Printf("\t%v: %.2f руб.\n", category, sumForCategory)
	}
	printDivider()
}

func printDivider() {
	fmt.Println("\n\t=================")
}