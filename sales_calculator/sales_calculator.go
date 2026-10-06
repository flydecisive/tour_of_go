// калькулятор скидок

package main

import "fmt"

func main() {
    var itemCost float64
	var discount int
	fmt.Print("Введите стоимость товара: ")
	_, err := fmt.Scan(&itemCost)

	fmt.Print("Введите скидку в процентах: ")
	_, err2 := fmt.Scan(&discount)

	if err != nil || err2 != nil {
		fmt.Println("Ошибка ввода: ", err)
		return
	}

	if discount > 100 {
		fmt.Print("Скидка не может быть больше 100%. Введите правильную скидку: ")
		_, err = fmt.Scan(&discount)

		if err != nil {
			fmt.Println("Ты ввел неправильную скидку")
			return
		}
	}
	
	if discount >= 0 && discount <= 100 {
		fmt.Printf("Итоговая стоимость: %.2f руб.", itemCost - (itemCost * float64(discount) / 100))
	}
}
