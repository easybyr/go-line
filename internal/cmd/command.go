package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)


var (
	rootCmd *cobra.Command
	Env string
	Config string
)

func init() {
	rootCmd = newCommand()
	flagCommand()
}

func newCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "go-line",
		Long: "go-line project command line",
		Version: shortVersion(),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("env:%v, config:%v\n", Env, Config)
		},
	}
	return cmd
}

func flagCommand() {
	rootCmd.PersistentFlags().StringVarP(&Env, "env", "e", "", 
		"environment, can be 'dev', 'tst', or 'prod', default to 'dev'.")
	rootCmd.PersistentFlags().StringVarP(&Config, "config", "c", "", 
		"configuration file path, can be file path or absolute directory")
}

func shortVersion() string {
	return ""
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(-1)
	}
}
