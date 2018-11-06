package util

import (
	"fmt"
	"os/exec"
)

func Run(cmd *exec.Cmd, dir string) error {
	fmt.Printf("Path: %s, Command: %s\n", dir, cmd.Args)
	cmd.Dir = dir
	return cmd.Run()
}
