package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"jzsg.com/mca/core/common"
	"jzsg.com/mca/core/server"
)

const (
	CmdRoot = "server"
	)

// The main command describes the service and
// defaults to printing the help message.
var mainCmd = &cobra.Command{Use: CmdRoot}

func main() {
	// For environment variables.
	viper.SetEnvPrefix(CmdRoot)
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// Define command-line flags that are valid for all peer commands and
	// subcommands.
	//mainFlags := mainCmd.PersistentFlags()
	//
	//mainFlags.String("logging-level", "", "Legacy logging level flag")
	//viper.BindPFlag("logging_level", mainFlags.Lookup("logging-level"))
	//mainFlags.MarkHidden("logging-level")
	mainCmd.PersistentFlags()

	mainCmd.AddCommand(versionCommand)
	mainCmd.AddCommand(server.StopCmd())
	mainCmd.AddCommand(server.StartCmd())

	// On failure Cobra prints the usage message and error string, so we only
	// need to exit with a non-0 status
	if err := mainCmd.Execute(); err != nil {
		panic(err)
		os.Exit(1)
	}
}

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Print version.",
	Long:  `Print current version of the server.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return fmt.Errorf("trailing args detected")
		}
		// Parsing of the command line is done so silence cmd usage
		cmd.SilenceUsage = true
		fmt.Print(GetVersionInfo())
		return nil
	},
}

func GetVersionInfo() string {
	return fmt.Sprintf(
		"%s:\n Version: %s\n Commit SHA: %s\n Go version: %s\n OS/Arch: %s\n",
		CmdRoot,
		common.Version,
		common.CommitSHA,
		runtime.Version(),
		fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	)
}