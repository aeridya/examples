package main

import (
	"fmt"
	"os"

	"github.com/aeridya/core"
)

func main() {
	e := core.Create("./conf")
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
	them := &ATheme{}
	err := them.Init()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	e = core.Run()
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}
