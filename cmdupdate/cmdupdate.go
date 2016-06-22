package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

const version = "0.1.3"

func main() {
	up := flag.Bool("update", false, "update command")
	flag.Parse()

	fmt.Printf("version %s\n", version)

	var code int
	if *up {
		if out, err := update(srcPath); err != nil {
			fmt.Printf("Update fail: %v\n", err)
			fmt.Println(string(out))
			os.Exit(1)
		}
		fmt.Println("Update done.")
		os.Exit(0)
	}

	fmt.Println("Command done.")

	os.Exit(code)
}

const srcPath = "github.com/dvrkps/dojo/cmdupdate"

func update(srcPath string) ([]byte, error) {
	cmd := exec.Command("go", "get", "-u", srcPath)
	return cmd.CombinedOutput()
}
