package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Coverage code to allow for cross platform tests

func main() {
	if len(os.Args) < 2 {
		panic("Panics are bad and this should be changed but YEESH put an arg in")
	}
	platform := os.Args[1]
	fmt.Println("Starting coverage tester for ", platform)

	switch platform {
	case "ubuntu-latest":
		ubuntu()
	case "windows-latest":
		windows()
	case "macos-latest":
		macos()
	default:
		panic("VM being run-on is unknown")
	}

}

func ubuntu() {
	// cmd := exec.Command("cd", "screen/")

	cmd := exec.Command("go", "test", "-coverprofile=profile.out", "-covermode=atomic", "./screen/...")
	cmd2 := exec.Command("cat", "profile.out", ">>", "coverage.txt")
	cmd3 := exec.Command("rm", "profile.out")
	// "cat profile.out >> coverage.txt",
	// "rm profile.out")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Erring", err)
	}
	err2 := cmd2.Run()
	if err2 != nil {
		fmt.Println("2", err)
	}
	err3 := cmd3.Run()
	if err3 != nil {
		fmt.Println("3", err)
	}

}
func windows() {

}

func macos() {

}
