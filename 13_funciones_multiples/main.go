package main

func main() {
	num1 := 5
	val1, _, val3 := valoresMultiples(num1)
	println(val1, val3)
}

func valoresMultiples(num1 int) (int, int, int) {
	return num1, num1 * 2, num1 * 3
}
