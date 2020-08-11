package main

import (
	"fmt"
	"strings"
)

func main() {
	NameParams := make(map[string]Animal)

	for {
		cmd, arg1, arg2 := getReqFromConsole()

		switch cmd {
		case "newanimal":
			var a Animal
			switch arg2 {
			case "cow":
				var cow *Cow
				a = cow
			case "bird":
				var bird *Bird
				a = bird
			case "snake":
				var snake *Snake
				a = snake
			default:
				fmt.Println("Invalid animal type! Accepted types are:\n\t1. cow\n\t2. bird\n\t3. snake")
				continue
			}
			NameParams[arg1] = a
			fmt.Println("Sucess!")

		case "query":
			var isName bool = false
			for name, _ := range NameParams {
				if name == arg1 {
					isName = true
					break
				} 
			}

			if !isName {
				fmt.Println("Invalid name! Try newanimal command")
				continue
			}

			switch arg2 {
			case "eat":
				NameParams[arg1].Eat()
			case "move":
				NameParams[arg1].Move()
			case "speak":
				NameParams[arg1].Speak()
			default:
				fmt.Println("Invalid action! Accepted actions are:\n\t1. eat\n\t2. move\n\t3. speak")
			}

		default:
			fmt.Printf("\"%s %s %s\" wasn't recognized as a valid command\n", cmd, arg1, arg2)
		}
	}
}

func getReqFromConsole() (string, string, string) {
	var cmd, arg1, arg2 string
	fmt.Print(">")
	fmt.Scanf("%s %s %s", &cmd, &arg1, &arg2)

	return strings.ToLower(cmd), strings.ToLower(arg1), strings.ToLower(arg2)
}

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food       string
	locomotion string
	noise      string
}

func (c *Cow) Eat() {
	if c == nil {
		fmt.Println("grass")
	} else {
		fmt.Println(c.food)
	}
}

func (c *Cow) Move() {
	if c == nil {
		fmt.Println("walk")
	} else {
		fmt.Println(c.locomotion)
	}
}

func (c *Cow) Speak() {
	if c == nil {
		fmt.Println("moo")
	} else {
		fmt.Println(c.noise)
	}
}

type Bird struct {
	food       string
	locomotion string
	noise      string
}

func (b *Bird) Eat() {
	if b == nil {
		fmt.Println("worms")
	} else {
		fmt.Println(b.food)
	}
}

func (b *Bird) Move() {
	if b == nil {
		fmt.Println("fly")
	} else {
		fmt.Println(b.locomotion)
	}
}

func (b *Bird) Speak() {
	if b == nil {
		fmt.Println("peep")
	} else {
		fmt.Println(b.noise)
	}
}

type Snake struct {
	food       string
	locomotion string
	noise      string
}

func (s *Snake) Eat() {
	if s == nil {
		fmt.Println("mice")
	} else {
		fmt.Println(s.food)
	}
}

func (s *Snake) Move() {
	if s == nil {
		fmt.Println("slither")
	} else {
		fmt.Println(s.locomotion)
	}
}

func (s *Snake) Speak() {
	if s == nil {
		fmt.Println("hsss")
	} else {
		fmt.Println(s.noise)
	}
}
