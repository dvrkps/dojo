package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

const version = "0.1.2"

func main() {
	up := flag.Bool("update", false, "update command")
	flag.Parse()

	fmt.Printf("version %s\n", version)

	var code int
	if *up {
		code = update()
	}

	fmt.Println("Command done.")

	os.Exit(code)
}

const srcPath = "github.com/dvrkps/dojo/cmdupdate"

func update() int {
	fmt.Println("Update start.")
	cmd := exec.Command("go", "get", "-u", srcPath)
	out, err := cmd.CombinedOutput()
	var code int
	if err != nil {
		fmt.Println("ERROR")
		code = 1
	}
	fmt.Println(string(out))
	if code < 1 {
		fmt.Println("Update done.")
	} else {
		fmt.Println("Update fail.")
	}
	return code
}
