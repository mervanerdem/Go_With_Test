package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}
func Div(a, b int) (res float64, err error) {
	if b == 0 {
		fmt.Errorf("Zero can not be divided")
		return 0, err
	}
	res = float64(a) / float64(b)
	return res, nil
}

func main() {

}
