// Генератор паролей
// С правильной обработкой ввода
package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Генератор паролей.")
	
	for {
		fmt.Print("Укажите длинну пароля: ")
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка ввода.")
			}
		}
		input := scanner.Text()
		
		passwordLen, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Введено не число. Попробуйте еще раз.")
			continue
		}

		if passwordLen > 0 {
			fmt.Printf("Ваш пароль: %s", generatePassword(int32(passwordLen)))
			break
		} else {
			fmt.Println("Число должно быть больше 0. Попробуйте снова.")
			continue
		}
	}
}

func generatePassword(limit int32) string{
	var stringBuilder strings.Builder
	var symbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%"
	for i := 0; i < int(limit); i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(symbols))))
		stringBuilder.WriteString(string(rune(symbols[int(n.Int64())])))
	}
	return stringBuilder.String()
}