package main

import (
	"fmt"
	"math"
)

func firstarea(radius float64) float64 {
	return math.Pi * math.Pow(radius, 2)
}
func main() {
	var radius float64
	fmt.Print("Enter radius: ")

	_, err := fmt.Scan(&radius)
	if err != nil {
		fmt.Println("wrong,please input a number")
		return
	}
	area := firstarea(radius)
	fmt.Println("radius is%.2f,area is%.2f\n", radius, area)

}
