package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

const usage = `notify <command> - send a desktop notification after command returns`

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}

	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	start := time.Now()

	var exitCode int
	var status string
	err := cmd.Run()
	if err != nil {
		status = err.Error()
		fmt.Println(status)
		switch err := err.(type) {
		case *exec.Error:
			exitCode = 1
		case *exec.ExitError:
			exitCode = err.ProcessState.Sys().(syscall.WaitStatus).ExitStatus()
		}
	} else {
		status = cmd.ProcessState.String()
	}

	notifycmd := fmt.Sprintf(`display notification "%s %s
%s" with title "Command complete"`,
		escape(status),
		time.Since(start),
		escape(strings.Join(os.Args[1:], " ")),
	)
	out, err := exec.Command("osascript", "-e", notifycmd).CombinedOutput()
	if err != nil {
		fmt.Println("notify error:", err, "\noutput:", string(out))
	}

	os.Exit(exitCode)
}

func escape(s string) string {
	return strings.Replace(s, `"`, `\"`, -1)
}
