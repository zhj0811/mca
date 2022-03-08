package server

import (
	"fmt"
	"io/ioutil"
	"os/exec"

	"github.com/spf13/cobra"
)

func StopCmd() *cobra.Command {
	stopCmd := &cobra.Command{
		Use:   "stop",
		Short: "Stop service",
		Run: func(cmd *cobra.Command, args []string) {
			pid, err := ioutil.ReadFile(fmt.Sprintf("%s.pid", CmdRoot))
			if err != nil {
				panic(err)
			}
			command := exec.Command("kill", string(pid))
			command.Start()
			fmt.Println(fmt.Sprintf("%s stop", CmdRoot))
		},
	}
	return stopCmd
}
