package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	conf "github.com/slackhq/nebula/config"
	"github.com/spf13/cobra"
)

var (
	l          *logrus.Logger
	configPath string
	config     *conf.C

	rootCmd = &cobra.Command{}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.Use = os.Args[0]
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "", "Path to either a file or directory to load configuration from")

	rootCmd.AddCommand(initCmd, unsealCmd, networkCmd, caCmd, enrollCmd, userCmd)
}

func initConfig() {
	l = logrus.New()
	l.Out = os.Stdout
	config = conf.NewC(l)
	if configPath != "" {
		err := config.Load(configPath)
		if err != nil {
			l.WithError(err).Printf("failed to load config")
			os.Exit(1)
		}
	}
}

func main() {
	Execute()
}
