package main

import (
	"fmt"
)

func Caesar(offset int, input string) string {
	// your code here
	var hasil string
	str := []uint8(input)
	huruf := 26
	z := 122
	if offset > huruf {
		offset = offset % huruf
	}
	for i := 0; i < len(str); i++ {
		if str[i]+uint8(offset) > uint8(z) {
			str[i] = str[i] - uint8(huruf)
		}
		hasil += string(str[i] + uint8(offset))
	}
	return hasil

}

func main() {
	fmt.Println(Caesar(3, "abc"))                         //def
	fmt.Println(Caesar(2, "alta"))                        // cnvc
	fmt.Println(Caesar(10, "alterraacademy"))             //kvdobbkkmknowi
	fmt.Println(Caesar(1, "abcdefghijklmnopqrstuvwxyz"))  //bcdefghijklmnopqrstuvwxyza
	fmt.Print(Caesar(1000, "abcdefghijklmnopqrstuvwxyz")) //mnopqrstuvwxyzabcdefghijkl
}
