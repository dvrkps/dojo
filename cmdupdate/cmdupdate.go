package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

const version = "0.1.0"

func main() {
	up := flag.Bool("update", true, "update command")
	flag.Parse()

	fmt.Printf("version %s\n", version)

	var code int
	if *up {
		code = update()
	}

	os.Exit(code)
}

const srcPath = "github.com/dvrkps/dojo/cmdupdate"

func update() int {
	cmd := exec.Command("go", "get", "-u", srcPath)
	out, err := cmd.CombinedOutput()
	var code int
	if err != nil {
		fmt.Println("ERROR")
		code = 1
	}
	fmt.Println(string(out))
	return code
}
