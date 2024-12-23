package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	for _, cmd := range []*cobra.Command{runCmd, initCmd} {
		rootCmd.AddCommand(cmd)
	}
}

var (
	rootCmd = &cobra.Command{
		Use:   "app",
		Short: "app应用",
		Long:  "app应用",
	}

	runCmd = &cobra.Command{
		Use:   "run",
		Short: "启动命令",
		Long:  "启动命令",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("                    _     _              __                                             _    ")
			fmt.Println("   __ _  ___       | |__ (_)_ __        / _|_ __ __ _ _ __ ___   _____      _____  _ __| | __")
			fmt.Println("  / _` |/ _ \\ _____| '_ \\| | '_ \\ _____| |_| '__/ _` | '_ ` _ \\ / _ \\ \\ /\\ / / _ \\| '__| |/ /")
			fmt.Println(" | (_| | (_) |_____| |_) | | | | |_____|  _| | | (_| | | | | | |  __/\\ V  V / (_) | |  |   < ")
			fmt.Println("  \\__, |\\___/      |_.__/|_|_| |_|     |_| |_|  \\__,_|_| |_| |_|\\___| \\_/\\_/ \\___/|_|  |_|\\_\\")
			fmt.Println("  |___/                                                                                       ")
			RunApp()
		},
	}

	initCmd = &cobra.Command{
		Use:   "init",
		Short: "初始化命令",
		Long:  "初始化命令",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
