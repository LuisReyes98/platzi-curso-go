package main

func main() {
	i := 0
	for i < 5 {
		println(i)
		i++
	}

	for j := 0; j < 5; j++ {
		println("J:", j)
	}

	for rango := range 3 {
		println("Rango:", rango)
	}

	for {
		println("Loop infinito")
		break
	}

	for rango2, value := range []int{4, 5, 6} {
		println("Rango2:", rango2)
		println("Value:", value)
	}

	for i := range 6 {
		if i%2 == 0 {
			continue
		}

		println("i impar:", i)
	}
}
