package main

import "fmt"

func main() {
	a := make([]string, 4, 5)
	a[0] = "a"
	a[2] = "c"
	a[3] = "d"
	as := a[0:2]
	as[1] = "Z"
	c := append(a, "y")

	c[0] = "q"

	fmt.Println(a, len(a), cap(a))
	fmt.Println(c, len(c), cap(c))
	fmt.Printf("%x %x %x\n", &a[0], &as[0], &c[0])
}
