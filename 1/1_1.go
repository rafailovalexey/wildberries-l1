package main

import "fmt"

/*
	№ 1 (1-ое решение)

	Дана структура Human (с произвольным набором полей и методов).
	Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/

type Human struct {
	Name  string
	Age   int
	Email string
}

type Action struct {
	Human
}

func (h *Human) GetName() string {
	return h.Name
}

func (h *Human) GetAge() int {
	return h.Age
}

func (h *Human) GetEmail() string {
	return h.Email
}

func main() {
	human := &Human{
		Name:  "Alexey",
		Age:   21,
		Email: "literallymyfault@gmail.com",
	}

	action := &Action{
		*human,
	}

	fmt.Printf("%s\n", action.GetName())
	fmt.Printf("%d\n", action.GetAge())
	fmt.Printf("%s\n", action.GetEmail())

	action2 := &Action{}

	fmt.Printf("%s\n", action2.GetName())
	fmt.Printf("%d\n", action2.GetAge())
	fmt.Printf("%s\n", action2.GetEmail())
}
