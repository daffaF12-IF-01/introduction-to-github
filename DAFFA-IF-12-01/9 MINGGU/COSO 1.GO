package main

import "fmt"

func main () {
	var a int
	var teks string 
	fmt.scan (&a)
	teks = "negatif"
	if a>0 {
		teks= "positif"
	}
	fmt.println(teks)
}
