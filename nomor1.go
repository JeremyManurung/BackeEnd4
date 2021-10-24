package main

import (
	"fmt"
	"math"
)

func Compare(a, b string) string {

	Max := int(math.Max(float64(len(a)), float64(len(b))))
	S := int(math.Min(float64(len(a)), float64(len(b))))

	Pstring := b
	Sstring := a

	if Max == len(a) {
		Pstring = a
		Sstring = b
	}

	var Tstring, Sama string
	var Tampung int

	for i := range Sstring {
		for j := range Pstring {
			Tampung = 0
			Sama = ""

			for int(i+Tampung) < S && int(j+Tampung) < Max && Sstring[i+Tampung] == Pstring[j+Tampung] {
				Sama += string(Pstring[j+Tampung])
				Tampung += 1
			}

			if len(Sama) > len(Tstring) {
				Tstring = Sama
			}
		}
	}

	return Tstring

}

func main() {

	fmt.Println(Compare("AKA", "AKASHI")) // AKA

	fmt.Println(Compare("KANGOORO", "KANG")) // KANG

	fmt.Println(Compare("KI", "KIJANG")) // KI

	fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU

	fmt.Println(Compare("ILALANG", "ILA")) // ILA

	fmt.Println(Compare("SEPATU", "EPA")) // EPA

	fmt.Println(Compare("ABCD", "KFAN")) //

}