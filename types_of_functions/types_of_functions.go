// примеры типов функций.
// присвоение функции переменной
// передача функции в другую функцию
// возврат функции из функции

package main

import "fmt"

func main() {
	var f func(int,int) int = add
	fmt.Println(f(3, 4))

	m := test(5, 6, add)
	fmt.Println(m)

	// Если при присвоении передать параметр, то он будет работать для ф-ции test2
	c := test2()
	// Тут при вызове передаются параметры уже для возвращаемой функции
	fmt.Println(c(8, 9))
}

func add (x int, y int) int {
	return x + y
}

func test (x, y int, add func(int,int) int) int{
	fmt.Print("Передача функции в качестве параметра в другую функцию: ")
	return add(x, y)
}

func test2 () (func(int,int) int) {
	fmt.Print("Возврат функции из другой функции: ")
	return func(a, b int) int {return a * b}
}