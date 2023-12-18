package main

import (
	"log"
)

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

	log.Printf("%s\n", action.GetName())
	log.Printf("%d\n", action.GetAge())
	log.Printf("%s\n", action.GetEmail())

	action2 := &Action{}

	log.Printf("%s\n", action2.GetName())
	log.Printf("%d\n", action2.GetAge())
	log.Printf("%s\n", action2.GetEmail())
}
