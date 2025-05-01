package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "ddd-linter",
	Short: "DDD Linter — проверяет архитектурные соглашения в DDD проектах",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		panic(err)
	}
}
