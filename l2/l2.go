package main

import (
	"fmt"
	"os"
)

type Command struct {
	command  string
	shortcut string
	callback func() bool
}

type Commands = []Command

func main() {
	arguments := os.Args

	commands := &Commands{
		{
			command:  "--test",
			shortcut: "-t",
			callback: func() bool {
				fmt.Println("test")

				return false
			},
		},
		{
			command:  "--help",
			shortcut: "-h",
			callback: func() bool {
				fmt.Println("help")

				return true
			},
		},
	}

	dictionary := make(map[string]Command, 10)

	for _, command := range *commands {
		if _, isExist := dictionary[command.command]; !isExist {
			dictionary[command.command] = command
		}

		if _, isExist := dictionary[command.shortcut]; !isExist {
			dictionary[command.shortcut] = command
		}
	}

	for _, argument := range arguments {
		if command, isExist := dictionary[argument]; isExist {
			exit := command.callback()

			if exit {
				break
			}
		}
	}
}
