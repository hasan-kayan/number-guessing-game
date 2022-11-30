package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Aklimdan 0 ile 10 arasinda bir sayi tutacagimve senden tahmin etmeni bekliycem  :)")
	fmt.Println("0 ile 10 arasinda bir sayi tut bakalim :) : ")
	rand.Seed(time.Now().UnixNano())
	var random int = rand.Intn(10)
	var i int
	var n int = 0
	for i != random && n < 4 {
		fmt.Scan(&i)
		if i >= 10 || i <= 0 {
			fmt.Println("0 ile 10 arasinda dedim, 0 ve 10 buna dahil degil :)")
		} else {
			if i > random {
				fmt.Println("Biraz yukarda geziyorsun sanki :) : ")
			} else if i < random {
				fmt.Println("Kucuk dusunmemek lazim :) : ")
			} else {
				fmt.Println("UBN-Jr uyesi tebrik ederim kazandin :) ")
			}
			n++
		}

	}
	if i != random {
		fmt.Println("3 guzel tahmin yaptin, bu sefer degilmis tekrar oynamak isterim :). Tuttugum sayi: ", random)
	}
}
