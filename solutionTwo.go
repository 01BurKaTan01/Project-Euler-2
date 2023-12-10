package main

import "fmt"

func main() {
	var elemanBir int = 1
	var elemanIki int = 2
	var elemanKayit int = 0
	var toplam int = 2
	toplamAnlik := 0

	for elemanIki < 4000000 {

		toplamAnlik = elemanBir + elemanIki

		if toplamAnlik%2 == 0 {
			toplam += toplamAnlik
		}

		elemanKayit = elemanIki
		elemanIki += elemanBir
		elemanBir = elemanKayit

	}

	fmt.Println(toplam)
}
