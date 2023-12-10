package main

import "fmt"

var (
	elemanBir   int = 1
	elemanIki   int = 2
	elemanKayit int = 0
	toplam      int = 2
	toplamAnlik int = 0
)

func main() {
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
