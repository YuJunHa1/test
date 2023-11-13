package main

import "fmt"

func main() {
	s := []int{0, 0, 0, 0, 0}

	s[4] = 100
	s[0] = 7
	s[1] = 91

	for _, v := range s {
		fmt.Println(v)
	}

	copyS := s[1:4]
	for _, v := range copyS {
		fmt.Println(v)
	}

	test := [3]string{"inha", "go", "student"}
	testS := test[:2]
	testS2 := test[1:]

	//testS2[0] = "python"
	//test[1] = "python"
	testS[1] = "python"

	fmt.Println(test, len(test))
	fmt.Println(testS, len(testS))
	fmt.Println(testS2, len(testS2))
}
