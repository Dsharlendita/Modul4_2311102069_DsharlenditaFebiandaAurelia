package main

import "fmt"

func faktorial(n int) int {
	if n == 0 {
		return 1
	}
	fakt := 1
	for i := 1; i <= n; i++ {
		fakt *= i
	}
	return fakt
}

func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	var a, b, c, d int

	fmt.Scan(&a, &b, &c, &d)

	Pac := permutasi(a, c)
	Cac := kombinasi(a, c)

	Pbd := permutasi(b, d)
	Cbd := kombinasi(b, d)

	fmt.Println(Pac, Cac)
	fmt.Println(Pbd, Cbd)
}
