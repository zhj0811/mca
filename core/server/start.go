package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"jzsg.com/mca/common/zlogging"
	"jzsg.com/mca/core/db"
	"jzsg.com/mca/core/server/config"
	"jzsg.com/mca/core/server/handler"
	"jzsg.com/mca/core/utils"
)

var confPath string

func StartCmd() *cobra.Command {
	var daemon bool
	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Starts the node.",
		Long:  `Starts a node that interacts with the network.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			//if len(args) != 0 {
			//	return fmt.Errorf("trailing args detected")
			//}
			// Parsing of the command line is done so silence cmd usage
			cmd.SilenceUsage = true
			fmt.Printf("Args: %v\n", os.Args)
			if daemon {
				fmt.Println("Start app as a daemon")
				as := os.Args[1:]
				for i := 0; i < len(as); i++ {
					if as[i] == "-d=true" || as[i] == "-d" {
						as[i] = "-d=false"
						break
					}
				}
				command := exec.Command(os.Args[0], as...)
				command.Start()

				if command.Process == nil {
					panic("process is nil")
				}
				//daemon = false
				os.Exit(0)
			}
			pidFile := fmt.Sprintf("%s.pid", CmdRoot)
			if utils.CheckFileIsExist(pidFile) {
				os.Remove(pidFile)
			}
			ioutil.WriteFile(pidFile, []byte(fmt.Sprintf("%d", os.Getpid())), 0666)
			return serve(args)
		},
	}
	startCmd.Flags().BoolVarP(&daemon, "daemon", "d", false, "is daemon?")
	startCmd.Flags().StringVarP(&confPath, "config", "c", "./config.yaml", "config file path")
	return startCmd
}

func serve(args []string) error {
	fmt.Println("start ...")
	//f, err := os.OpenFile("./test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//if err != nil {
	//	panic(err)
	//}
	err := config.InitConfig(confPath)
	if err != nil {
		panic(err)
	}
	zlogging.SetWriter(config.GetDefaultLogWriter())
	//log.SetOutput(config.GetDefaultLogWriter())

	err = db.InitDb()
	if err != nil {
		panic(err)
	}
	err = handler.InitRouter()
	if err != nil {
		panic(err)
	}

	//startHttp()
	return nil
}

func startHttp() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello cmd!")
	})
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
