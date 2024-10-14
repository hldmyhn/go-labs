package main

import (
	"errors"
	"fmt"
	"math"
)

func formatIP(ip [4]byte) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func listEven(start, end int) ([]int, error) {
	if start > end {
		return nil, errors.New("левая граница больше правой")
	}

	var evens []int
	for i := start; i <= end; i++ {
		if i%2 == 0 {
			evens = append(evens, i)
		}
	}

	return evens, nil
}

func countChars(s string) map[rune]int {
	charCount := make(map[rune]int)
	for _, char := range s {
		charCount[char]++
	}
	for char, count := range charCount {
		fmt.Printf("'%c': %d\n", char, count)
	}

	return charCount
}

type Point struct {
	X, Y float64
}

type Triangle struct {
	A, B, C Point
}

type Circle struct {
	Center Point
	Radius float64
}

func (t Triangle) Area() float64 {
	return math.Abs((t.A.X*(t.B.Y-t.C.Y) + t.B.X*(t.C.Y-t.A.Y) + t.C.X*(t.A.Y-t.B.Y)) / 2.0)
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Shape interface {
	Area() float64
}

func printArea(s Shape) {
	fmt.Printf("Площадь фигуры: %.2f\n", s.Area())
}

func Map(slice []float64, fn func(float64) float64) []float64 {
	result := make([]float64, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

func main() {
	fmt.Println("IP:", formatIP([4]byte{127, 0, 0, 1}))

	evens, err := listEven(1, 10)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Чётные числа:", evens)
	}

	countChars("abc aaa b")

	printArea(Triangle{Point{0, 0}, Point{5, 0}, Point{0, 5}})
	printArea(Circle{Point{0, 0}, 10})

	slice := []float64{1.0, 2.0, 3.0, 4.0}

	square := func(x float64) float64 {
		return x * x
	}

	fmt.Println([]float64{1.0, 2.0, 3.0, 4.0}, "-->", Map(slice, square))

}
