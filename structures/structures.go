// Примеры работы со структурами

// Прмиер структуры
//
//	type person struct {
//		name string
//		age int
//	}
//
// var tom person - переменная структуры person

// Инициализация структур:
// var alice person = person{age: 23, name: "Alice"}
// var tom = person {name: "Tom", age: 24}
// bob := person {name: "Bob", age: 31}

// Анонимная структура:
// var tom struct{
//     name string
//     age int
// }

package main

import "fmt"

type test struct {
	a int
}

type test2 struct {
	a string;
	b int
}

func main() {
	// Простая структура
	var s1 test = test{22}
	fmt.Println("Вывод поля s1: ", s1.a)

	// Присвоение значения полю структуры
	s1.a = 43
	fmt.Println("s1.a = 43: ", s1.a)

	// Анонимная структура + присвоение значения
	var s2 struct{
		b string
	}
	s2.b = "test"
	fmt.Println("s2.b = \"test\": ", s2.b)

	// Анонимная структура с инициализацией
	s3 := struct{c bool}{true}
	fmt.Println("Вывод простой анонимной структуры с инициализацией: ",s3.c)

	// Указатель на поле структуры
	var p_a *int = &s1.a
	fmt.Println("Указатель на поле структруры: ", p_a)

	// Структруа с инициализацией
	s4 := test2{b: 3, a: "test2"}
	var p_s4 *test2 = &s4
	fmt.Println("Указатель на структуру: ", p_s4)

	// Структура с указателем и инициализацией
	var s5 *test2 = &test2{"test3", 8}
	fmt.Println("Вывод структуры с сылкой на неизвестный объект (указатель на структуру): ", s5)
	fmt.Println("Вывод поля этой структуры, обращаясь на прямую: ", s5.a)
	fmt.Println("Вывод поля этой структуры, образаясь через разименование: ", (*p_s4).a)

	// Создание структуры с помощью new
	s6 := new(test)
	fmt.Println("Структура создана с помощью new: ", *s6)

	increment_a(s1)
	fmt.Println("Исходное значение a в структуре после выхова ф-ции: ", s1.a)

	increment_a2(&s1)
	fmt.Println("Исходное значение a в структуре после выхова ф-ции, структура передана через указатель: ", s1.a)
}

// Передача копии структуры, исходная структура не изменится
func increment_a(test test) {
	test.a += 1
	fmt.Println("Увеличивает значение a на 1 в ф-ции, исходная структура не меняется: ", test.a)
}

// Передача структуры через указатель, исходная структура изменится
func increment_a2(test *test) {
	test.a += 1
	fmt.Println("Увеличивает значение a на 1 в ф-ции, исходная структура меняется, структура передана через указатель: ",test.a)
}