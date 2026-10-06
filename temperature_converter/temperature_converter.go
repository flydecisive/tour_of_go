// конвертер температур

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	
	for {
		fmt.Print("Введите значение или exit для выхода: ")
		_, err := fmt.Scan(&input)
		
		if err != nil {
			fmt.Println("Ошибка ввода")
		} else {
			if input == "exit" {
				break
			} else {
				temp, err := strconv.ParseFloat(input, 64)
				if err != nil {
					fmt.Println("Ошибка конвертации")
					continue
				} else {
					convertTemp(temp)
				}
			}
		}
	}
}

func convertTemp(temp float64) {
	fahrenheit := temp * 1.8 + 32
	fmt.Printf("Температура %.2f°C равна %.2f°F\n", temp, fahrenheit)
}