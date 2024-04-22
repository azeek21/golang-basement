package main

import (
	"fmt"
	"os"
	"os/exec"
)

func getStdinArgs() []string {
	inp := ""
	res := []string{}
	for {
		if _, err := fmt.Scanf("%s\n", &inp); err != nil {
			break
		}
		res = append(res, inp)
	}
	return res
}

func Xargs(args []string) {
	stdinArgs := getStdinArgs()
	app := args[0]
	appArgs := append(args[1:], stdinArgs...)
	cmd := exec.Command(app, appArgs...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Printf("error during execution: %s\n", err.Error())
	}
}
