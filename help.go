package main

import (
	"fmt"
)

func Usage() {
	fmt.Println("usage: ccgit")
	fmt.Println("")
	fmt.Println("These are the supported commands")
	fmt.Println("")
	fmt.Println("start working area (see also: ccgit help tutorial)")
	fmt.Println("   init       Create an empty Git repository or reinitialize an existing one")
	fmt.Println("")
	fmt.Println("work on the current change (see also: git help everyday)")
	fmt.Println("   add       Add file contents to the index")
}

func CommandNotFound(command string) {
	fmt.Println("ccgit: ", command, " is not a ccgit command. See 'ccgit help'.")
}
