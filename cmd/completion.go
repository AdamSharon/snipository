package main

import (
	"os"

	"github.com/spf13/cobra"
)

func zshCompletionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "completion",
		Short: "Generates zsh completion scripts",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Root().GenZshCompletion(os.Stdout)
		},
	}
}
