package main

import (
	"errors"
	"fmt"
)

func hello(name string) string {
	return fmt.Sprintf("Привет, %s!", name)
}

func printEven(start, end int64) error {
	if start > end {
		return fmt.Errorf("граница диапазона (%d) больше (%d)", start, end)
	}

	for i := start; i <= end; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	return nil
}

func apply(a, b float64, operation string) (result float64, err error) {
	switch operation {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			return 0, errors.New("деление на ноль")
		}
		result = a / b
	default:
		return 0, fmt.Errorf("операция %s не поддерживается", operation)
	}
	return result, nil
}

func main() {
	fmt.Println(hello("мир"))

	if err := printEven(1, 5); err != nil {
		fmt.Println("Ошибка:", err)
	}

	if err := printEven(5, 1); err != nil {
		fmt.Println("Ошибка:", err)
	}

	result, err := apply(3, 5, "+")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", result)
	}

	result, err = apply(7, 10, "*")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", result)
	}

	result, err = apply(3, 5, "#")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", result)
	}
}
