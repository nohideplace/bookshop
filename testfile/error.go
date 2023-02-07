package main

import "fmt"

type NameError struct {
}

func (e *NameError) Error() string {
	return "妹有名字哇"
}

func test(a string) (string, error) {
	if a == "" {
		return "", &NameError{}
	}
	return a, nil
}

func main() {
	a := "666"
	b := ""
	data1, err1 := test(a)
	data2, err2 := test(b)
	fmt.Println(data1, err1)
	fmt.Println(data2, err2)

}
