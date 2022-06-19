package main

import (
	"fmt"
)

var (
	f1 float64 = 5.99
	f2 float64 = 7
	f3 float64 = 23
)

// Одне яблуко(f1) коштує 5.99 грн. Ціна однієї груші(f2) - 7 грн.
// Ми маємо(f3) 23 грн.

func main() {

	fmt.Println("1.Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?")
	fmt.Println("Ціна яблука:", f1, "грн;", "Ціна груші:", f2, "грн")
	mult := 9*f1 + 8*f2
	fmt.Println("Вартість 9 яблук та 8 груш:", mult, "грн")

	fmt.Println("2.Скільки груш ми можемо купити?")
	fmt.Println("Ціна груші:", f2, "грн")
	fmt.Println("Купили груш:", int(float64(f3)/float64(f2)), "шт")

	fmt.Println("3.Скільки яблук ми можемо купити?")
	fmt.Println("Ціна яблука:", f1, "грн")
	fmt.Println("Купили яблук:", int(float64(f3)/float64(f1)), "шт")

	fmt.Println("4.Чи ми можемо купити 2 груші та 2 яблука?")
	fmt.Println("Ціна яблука:", f1, "грн;", "Ціна груші:", f2, "грн")
	canBuyBothThings := (f1*2 + float64(f2)*2) < float64(f3)
	fmt.Println("Чи ми можемо купити 2 груші та 2 яблука?", canBuyBothThings)

}
