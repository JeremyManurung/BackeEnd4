package main

import (
	"fmt"
)

func getMinMax(numbers ...*int) (min int, max int) {
	min = *numbers[0]
	max = *numbers[0]
	for _, n := range numbers {
		if *n < min {
			min = *n
		}
		if *n > max {
			max = *n
		}
	}
	return
}

func main() {

	var a1, a2, a3, a4, a5, a6, min, max int

	fmt.Print("Masukkan angka 1 : ")
	fmt.Scan(&a1)
	fmt.Print("Masukkan angka 2 : ")
	fmt.Scan(&a2)
	fmt.Print("Masukkan angka 3 : ")
	fmt.Scan(&a3)
	fmt.Print("Masukkan angka 4 : ")
	fmt.Scan(&a4)
	fmt.Print("Masukkan angka 5 : ")
	fmt.Scan(&a5)
	fmt.Print("Masukkan angka 6 : ")
	fmt.Scan(&a6)

	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println(max, "adalah nilai maksimalnya")
	fmt.Println(min, "adalah nilai minimlanya")

	

}