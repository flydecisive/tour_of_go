// Примеры указателей

package main

import "fmt"

func main() {
	fmt.Print("Базовый пример указателей:\n")
	var x int = 4 // создаем переменную
	var p *int = &x // указатель на перемеременную x
	fmt.Println("Address: ", p)
	fmt.Println("Value: ", *p)
	*p = 25 // меняем значение x через указатель
	fmt.Println("Значение x после *p=25: ", x)

	// Функция new
	fmt.Print("\nПример с функцией new():\n")
	newPointer := new(int)
	fmt.Println("Дефолтное значение new(int): ", *newPointer)
	*newPointer = 20
	fmt.Println("Новое значение: ", *newPointer)

	// Указатели на указатели
	fmt.Print("\nПример указателей на указатели:\n")
	var a int = 9

	var p_a *int = &a //указатель на a
	fmt.Println("Значение a: ", *p_a)
	fmt.Println("Адресс a: ", p_a)

	var p_p_a **int = &p_a // Указатель на указатель. Указывается 2 звездочки перед int

	fmt.Println("Значение p_a: ", *p_p_a)
	fmt.Println("Адресс p_a: ", p_p_a)
	fmt.Println("Значение a, обращаясь к p_p_a: ", **p_p_a)

	// Массив указателей
	fmt.Print("\nПример массива указателей:\n")
	var f, g, k int = 1, 2, 4
	var p_nums[4] *int
	p_nums[0] = &f
	p_nums[1] = &g
	p_nums[3] = &k

	fmt.Println("Массив указателей: ", p_nums)
	fmt.Println("Значение массива p_nums[1]: ", *p_nums[1])

	// Возврат указателя из функции
	fmt.Print("\nПример возврата указателя из функции:\n")
	p1 := createPointer(7)
	fmt.Println("p1: ", *p1, ";p1 адресс: ", p1)
	p2 := createPointer(10)
	fmt.Println("p2: ", *p2, ";p2 адресс: ", p2)

}

// Функция возвращающая указатель
func createPointer(x int) *int {
	p := new(int)
	*p = x
	return p
}