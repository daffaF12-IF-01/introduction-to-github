package main

import "fmt"

funt main () {
	var bilangan int
	fmt.scan (&bilangan)
	if bilangan <0 {
		bilangan = -bilangan
	}
	fmt.print (bilangan)
}
