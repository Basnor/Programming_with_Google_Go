package main

import "fmt"

func main() {

	//Animal         Food     Locomotion  Sound
	cow   := Animal{ "grass", "walk",     "moo" }
	bird  := Animal{ "worms", "fly",      "peep" }
	snake := Animal{ "mice",  "slither",  "hsss" }

	for {
		name, action := getReqFromConsole()

		switch name {
		case "cow":
			selectAction(action, cow)
		case "snake":
			selectAction(action, snake)
		case "bird":
			selectAction(action, bird)
		default:
			fmt.Println("Animal doesn't exist")
		}
	}
}

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func getReqFromConsole() (string, string) {
	var name string
	var action string
	fmt.Print(">")
	fmt.Scanf("%s %s", &name, &action)

	return name, action
}

func selectAction(action string, animal Animal) {
	switch action {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("Action doesn't exist")
	}
}