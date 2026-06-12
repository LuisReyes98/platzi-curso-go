package main

import (
	"time"
)

func main() {

	i := 2

	println(i)
	switch i {
	case 1:
		println("i es 1")
	case 2:
		println("i es 2")
	case 3:
		println("i es 3")
	default:
		println("i no es ni 1, ni 2, ni 3")

	}

	today := time.Now().Weekday()

	switch today {
	case time.Monday:
		println("Hoy es lunes")
	case time.Tuesday:
		println("Hoy es martes")
	case time.Wednesday:
		println("Hoy es miércoles")
	case time.Thursday:
		println("Hoy es jueves")
	case time.Friday:
		println("Hoy es viernes")
	default:
		println("Hoy es fin de semana")
	}

	j := 3

	switch {
	case j%2 == 0:
		println("j es par")
	case j < 5:
		println("j es menor que 5")
	default:
		println("j es impar")
	}

}
