package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func startEditor(path string) error {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		return errors.New("empty env editor")
	}

	cmd := exec.Command(editor, path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("run: %v", err)
	}

	return nil
}
