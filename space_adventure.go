package main
import (
	"fmt"
	"bufio"
)

fmt.Println("Welcome to the Solar System!")
fmt.Println("There are 9 planets to explore")
fmt.Println("What is your name?")

func main() {

    //reading a string
    reader := bufio.newReader(os.Stdin)
    var name string
    fmt.Println("What is your name?")
    name, _ := reader.readString("\n")

	fmt.Println("Nice to Meet you, ", name, " My name is Eliza, I'm an old friend of Alexa.")
	fmt.Println("Let's go on an adventure!")
}

