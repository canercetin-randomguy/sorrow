package main

import (
	"fmt"
	"os"
	"os/user"
	"sorrow/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Welcome to the DIY-Chemical-Recipe interpreter, %s!|\n", user.Username)
	fmt.Println("Dont forget to put ; at the end of each line or interpreter will get stuck!")
	fmt.Println("Feel free to type in commands")
	repl.Start(os.Stdin, os.Stdout)
}
