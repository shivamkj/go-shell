package main

import (
	"fmt"
	"strings"

	"github.com/shivamkj/go-shell"
)

func main() {

	// Running a simple command
	user, _ := shell.Sh("echo $USER")
	fmt.Println("Current user:", user)

	// Running a python script and providing it input
	sum, _ := shell.ShI("python3 ./example/test_sum.py", "2\n5\n")
	fmt.Println("2 + 5 =", sum)

	// Use standard Input and Output
	shell.Sh("read -p 'Your Name': name && echo Your name is: $name", shell.UseStdin, shell.UseStdOut)

	// Run command and process output
	if filesOutput, err := shell.Sh("ls"); err == nil {
		files := strings.Split(filesOutput, "\n")
		fmt.Println("=== Files in the current Directory: ===")
		for i, file := range files {
			fmt.Printf("File %d: %s \n", i, file)
		}
	}

	fmt.Println("\n === Script executed successfully ===")
}
