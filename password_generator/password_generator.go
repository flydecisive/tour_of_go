// Генератор паролей
// С правильной обработкой ввода
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var symbols := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%"
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Генератор паролей.")
	
	// for scanner.Scan()
	for {
		fmt.Print("Укажите длинну пароля: ")
		// scanner.Scan()
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка ввода.")
			}
		}
		input := scanner.Text()
		
		// if ()
		passwordLen, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Введено не число. Попробуйте еще раз.")
			continue
		}

		if passwordLen > 0 {
			fmt.Println("Ваш пароль: ")
			break
		} else {
			fmt.Println("Число должно быть больше 0. Попробуйте снова.")
			continue
		}
	}
}