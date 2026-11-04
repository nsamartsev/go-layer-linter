// SPDX-License-Identifier: MIT
// Copyright (c) 2026 Nikita Samartsev

package cli

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:           "go-layer-linter",
	Short:         "Check package import rules in layered Go projects",
	SilenceUsage:  true,
	SilenceErrors: true,
}

// Execute runs the command-line application.
func Execute() error {
	return rootCmd.Execute()
}
