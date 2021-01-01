package main

import (
	"fmt"
	"strings"
)

const upDog = "     __  \n" +
	"(___()'`;\n" +
	"/,    /` \n" +
	`\\"--\\  `

func dogSay(say string) {
	indent := strings.Repeat(" ", len(strings.Split(upDog, "\n")[0]))
	topBorder := strings.Repeat("-", len(say)+2)
	fmt.Println(indent + "┌" + topBorder + "┐")
	fmt.Println(indent + "| " + say + " |")
	fmt.Println(indent + "v" + topBorder + "┘")
	fmt.Println(upDog)
}

func main() {
	dogSay("Hi! I am Go and I love you!")
}
