package main

import "fmt"

/*
	№ 1 (1 решение)

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
		Age:   20,
		Email: "literallymyfault@gmail.com",
	}

	action := &Action{
		*human,
	}

	fmt.Println(action.GetName())
	fmt.Println(action.GetAge())
	fmt.Println(action.GetEmail())

	action2 := &Action{}

	fmt.Println(action2.GetName())
	fmt.Println(action2.GetAge())
	fmt.Println(action2.GetEmail())
}
