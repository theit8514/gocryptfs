package main

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"os/exec"
	"strings"
	"github.com/mgutz/str"
)

func hasAskPass() bool {
	return len(getAskPassEnv()) > 0
}

func getAskPassEnv() string {
	return os.Getenv("GOCRYPTFS_ASKPASS")
}

func getAskPassParts() []string {
	askpass := getAskPassEnv()
	return str.ToArgv(askpass)
}

func readAskPass() string {
	fmt.Printf("Executing askpass program.\n")
	parts := getAskPassParts()
	exeFile := parts[0]
	parts = parts[1:len(parts)]
	cmd := exec.Command(exeFile, parts...)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return strings.Trim(string(out), "\r\n")
}

func readPasswordTwice() string {
	fmt.Printf("Password: ")
	p1 := readPassword()
	fmt.Printf("\nRepeat: ")
	p2 := readPassword()
	fmt.Printf("\n")
	if p1 != p2 {
		fmt.Printf("Passwords do not match\n")
		os.Exit(ERREXIT_PASSWORD)
	}
	return p1
}

// Get password from terminal
func readPassword() string {
	fd := int(os.Stdin.Fd())
	p, err := terminal.ReadPassword(fd)
	if err != nil {
		fmt.Printf("Error: Could not read password: %v\n", err)
		os.Exit(ERREXIT_PASSWORD)
	}
	return string(p)
}
