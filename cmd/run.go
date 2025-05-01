package cmd

import (
	"fmt"

	"github.com/nsamartsev/go-layer-linter/internal/rules"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Запуск анализа DDD-проекта",
	Run: func(cmd *cobra.Command, args []string) {
		dir, _ := cmd.Flags().GetString("dir")
		issues := rules.RunAnalysis(dir)
		if len(issues) == 0 {
			fmt.Println("✅ Архитектурных нарушений не найдено.")
		} else {
			for _, issue := range issues {
				fmt.Println(issue)
			}
		}
	},
}

func init() {
	runCmd.Flags().StringP("dir", "d", ".", "Директория проекта")
	RootCmd.AddCommand(runCmd)
}
