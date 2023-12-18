package main

import (
	"log"
)

/*
	№ 8 (1-ое решение)

	Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

func SetBit(number int64, index int, value int) int64 {
	mask := int64(1) << index

	if value == 1 {
		number |= mask
	}

	if value == 0 {
		number &= ^mask
	}

	return number
}

func main() {
	number := int64(42)
	position := 1
	value := 1

	result := SetBit(number, position, value)

	log.Printf("original number %d\n", number)
	log.Printf("result after setting the %d bit to position %d: %d\n", value, position, result)
}
