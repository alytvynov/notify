package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
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

	err := cmd.Run()
	status := cmd.ProcessState.String()
	if err != nil {
		status = err.Error()
	}

	notifycmd := fmt.Sprintf(`display notification "%s completed
%s
duration = %s" with title "Command complete"`,
		strings.Replace(strings.Join(os.Args[1:], " "), `"`, `\"`, -1),
		strings.Replace(status, `"`, `\"`, -1),
		strings.Replace(time.Since(start).String(), `"`, `\"`, -1))
	out, err := exec.Command("osascript", "-e", notifycmd).CombinedOutput()

	if err != nil {
		fmt.Println("notify error:", err, "\noutput:", string(out))
		os.Exit(1)
	}
}
