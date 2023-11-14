package main

import (
	"fmt"
	"math/big"
)

func main() {
	value1 := new(big.Int).Exp(big.NewInt(2), big.NewInt(64), nil)
	value2 := new(big.Int).Exp(big.NewInt(2), big.NewInt(32), nil)

	m := multiply(value1, value2)
	fmt.Printf("Умножение: %s\n", m.String())

	d, err := divide(value1, value2)
	if err != nil {
		fmt.Println("Ошибка при делении:", err)
	}
	fmt.Printf("Деление: %s\n", d.String())

	a := add(value1, value2)
	fmt.Printf("Сложение: %s\n", a.String())

	s := subtract(value1, value2)
	fmt.Printf("Вычитание: %s\n", s.String())
}

func multiply(a *big.Int, b *big.Int) *big.Int {
	result := new(big.Int).Mul(a, b)

	return result
}

func divide(a *big.Int, b *big.Int) (*big.Int, error) {
	if b.Cmp(big.NewInt(0)) == 0 {
		return nil, fmt.Errorf("деление на ноль")
	}

	result := new(big.Int).Div(a, b)

	return result, nil
}

func add(a *big.Int, b *big.Int) *big.Int {
	result := new(big.Int).Add(a, b)

	return result
}

func subtract(a *big.Int, b *big.Int) *big.Int {
	result := new(big.Int).Sub(a, b)

	return result
}
