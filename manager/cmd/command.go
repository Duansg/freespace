package freespacecmd

import (
	"fmt"
	"os"

	"github.com/segmentfault/pacman/log"

	"github.com/Duansg/freespace/manager/cli"
	"github.com/spf13/cobra"
)

var (
	// dataDirPath save all answer application data in this directory. like config file, upload file...
	dataDirPath string
)

func init() {
	rootCmd.Version = fmt.Sprintf("%s\nrevision: %s\nbuild time: %s", Version)
	rootCmd.PersistentFlags().StringVarP(&dataDirPath, "data-path", "C", "./data", "data path, eg: -C ./data/")

	for _, cmd := range []*cobra.Command{runCmd} {
		rootCmd.AddCommand(cmd)
	}
}

var (
	rootCmd = &cobra.Command{
		Use:   "freespace",
		Short: "Free Space is a Simple blog platform.",
		Long: `Free Space is a Simple blog platform. To run freespace, use:
		- 'freespace init' to initialize the required environment.
		- 'freespace run' to launch application.`,
	}

	// runCmd represents the run command
	runCmd = &cobra.Command{
		Use:   "run",
		Short: "Run the application",
		Long:  `Run the application`,
		Run: func(_ *cobra.Command, _ []string) {
			log.Info("dataDirPath: ", dataDirPath)
			cli.FormatAllPath(dataDirPath)

			log.Info("Freespace is starting..........................")
			runApp()
		},
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
