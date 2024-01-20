package main

import (
	"os"
	"os/exec"
)

func main() {

	cmd := exec.Command("lua", "scan.lua")

	cmd.Stdin = os.Stdin

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		panic(err)
	}

}