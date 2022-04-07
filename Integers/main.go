package main

import "fmt"

func Add(a, b float64) float64 {
	return a + b
}

func Sub(a, b float64) float64 {
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
