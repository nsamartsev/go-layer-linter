// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Nikita Samartsev

package cli

import (
	"fmt"

	"github.com/nsamartsev/go-layer-linter/internal/rules"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Analyze a Go project",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		dir, err := cmd.Flags().GetString("dir")
		if err != nil {
			return fmt.Errorf("read --dir flag: %w", err)
		}
		issues, err := rules.RunAnalysis(dir)
		if err != nil {
			return err
		}
		if len(issues) == 0 {
			fmt.Fprintln(cmd.OutOrStdout(), "No architecture violations found.")
			return nil
		}
		for _, issue := range issues {
			fmt.Fprintln(cmd.OutOrStdout(), issue)
		}
		return fmt.Errorf("found %d architecture violation(s)", len(issues))
	},
}

func init() {
	runCmd.Flags().StringP("dir", "d", ".", "Project directory to analyze (default: current directory)")
	rootCmd.AddCommand(runCmd)
}
