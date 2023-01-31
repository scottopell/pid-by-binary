/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/scottopell/pid-by-binary/pkg/procfinder"
	"github.com/shirou/gopsutil/v3/process"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pid-by-binary",
	Short: "Given an executable binary, find all running instances of that binary",
	Long:  `Given an executable binary, find all running instances of that binary`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			pids, err := procfinder.FindProcesses(arg)
			if err != nil {
				fmt.Println("Error: ", err)
				continue
			}
			requestedMemoryStats, _ := cmd.Flags().GetBool("memory")

			if requestedMemoryStats {
				for _, pid := range pids {
					proc, err := process.NewProcess(pid)
					if err != nil {
						fmt.Println("Could not get process data for pid ", pid)
						continue
					}
					memoryStats, err := proc.MemoryInfo()
					if err != nil {
						fmt.Println("Could not get process data for pid ", pid)
						continue
					}
					fmt.Printf("Pid: %v\tMemInfo: %s\n", pid, memoryStats.String())
				}
			} else {
				for _, pid := range pids {
					fmt.Printf("%d\n", pid)
				}
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pid-by-binary.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("memory", "m", false, "Include Memory Stats")
	rootCmd.ValidArgs = []string{"binary-path"}
}
