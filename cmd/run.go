package cmd

import (
	"fmt"
	"os"

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
			os.Exit(0)
		} else {
			for _, issue := range issues {
				fmt.Println(issue)
			}
			fmt.Printf("\nНайдено %d нарушений.\n", len(issues))
			os.Exit(1)
		}
	},
}

func init() {
	runCmd.Flags().StringP("dir", "d", ".", "Директория проекта")
	RootCmd.AddCommand(runCmd)
}
