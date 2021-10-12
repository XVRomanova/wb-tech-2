package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		input, err := reader.ReadString('\n')
		input = input[0 : len(input)-1]
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err = shell(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

}

func shell(command string) error {
	args := strings.Split(command, " ")
	switch args[0] {
	case "cd":
		err := os.Chdir(args[1])
		if err != nil {
			return err
		}
		dir, err := os.Getwd()
		if err != nil {
			return err
		}
		output := "Current working directory: " + dir

		fmt.Println(output)

	case "exit":
		os.Exit(0)
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}