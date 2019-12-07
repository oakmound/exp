package main

import (
	"fmt"
	"io"
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
	wd, err := os.Getwd()
	fmt.Println("PWD:", wd, err)

	cmd := exec.Command("go", "test", "--coverprofile=profile.out", "--covermode=atomic", "./screen/...")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("go test err", err)
		os.Exit(1)
	}

	prof, err := os.Open("profile.out")
	if err != nil {
		fmt.Println("Failed to open profile", err)
		os.Exit(1)
	}
	cov, err := os.OpenFile("./coverage/coverage.txt", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to open coverage file", err)
		os.Exit(1)
	}
	written, err := io.Copy(cov, prof)
	if err != nil {
		fmt.Println("Failed to copy data", err)
		os.Exit(1)
	}
	fmt.Println("Wrote", written, "bytes")
}
func windows() {

}

func macos() {

}
