// примеры операций сдвига

package main

import "fmt"

func main() {
	const (
		a1  = iota
		a2 
		a3 
	)

	fmt.Println("Без сдвига: ", a1, a2, a3)

	const (
		b1  = 1 << iota
		b2 
		b3 
	)

	fmt.Println("Сдвиг влево на iota: ", b1, b2, b3)

	const (
		c1  = 1 >> iota
		c2 
		c3 
	)

	fmt.Println("Сдвиг вправо на iota: ", c1, c2, c3)

	const (
		d1  = 1 << iota + 2
		d2 
		d3 
	)

	fmt.Println("Сдвиг влево на iota + 2: ", d1, d2, d3)

	var (
		e = 8
		f = 100
	)

	fmt.Println("Сдвиг 8 влево на 3: ", e << 3)
	fmt.Println("Сдвиг 8 вправо на 1: ", e >> 1)
	fmt.Println("Сдвиг 100 влево на 5: ", f << 5)
	fmt.Println("Сдвиг 100 вправо на 4: ", f >> 4)
}