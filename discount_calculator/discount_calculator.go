package main

import (
	"fmt"
)

func main() {
	var (
		result      float64
		price       float64
		discount    int
		discountSum float64
	)

	fmt.Println("Программа для подсчета цены с учетом скидки.")

	fmt.Print("Введите цену товара: ")
	_, err := fmt.Scan(&price)

	if err != nil {
		fmt.Println("Ошибка: введите цену числом.")
		return
	}

	if price < 0 {
		fmt.Println("Цена не может быть отрицательной.")
		return
	}

	fmt.Print("Введите скидку в процентах: ")
	_, err = fmt.Scan(&discount)

	if err != nil {
		fmt.Println("Ошибка: введите скидку числом.")
		return
	}

	if discount < 0 || discount > 100 {
		fmt.Println("Скидка должна быть от 0 до 100%")
		return
	}

	discountSum = (price * float64(discount)) / 100
	result = price - discountSum

	fmt.Printf("\nСкидка: %.2f руб.\n", discountSum)
	fmt.Printf("К оплате: %.2f руб.", result)
}
