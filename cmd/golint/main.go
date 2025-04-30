package main

import (
	"fmt"
	"os"

	"github.com/nsamartsev/go-layer-linter/internal/analyzer"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "golint",
	Short: "Go Arch Linter",
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run linter",
	Run: func(cmd *cobra.Command, args []string) {
		dir := cmd.Flag("dir").Value.String()
		analyzer.Run(dir)
	},
}

func init() {
	runCmd.Flags().StringP("dir", "d", ".", "Project directory")
	rootCmd.AddCommand(runCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
