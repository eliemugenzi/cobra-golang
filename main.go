package main

import (
	"go-cobra/app/cmd"
	"go-cobra/app/welcome"
)

func main() {

	welcome.SayHello()
	cmd.Execute()
}
