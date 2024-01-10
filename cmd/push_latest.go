package main

import (
	"fmt"

	"snipository/common"
	"snipository/config"
	"snipository/history"
	"snipository/storage"

	"github.com/spf13/cobra"
)

func pushLatestCommand() *cobra.Command {
	newConfig := config.NewConfig()
	c := &cobra.Command{
		Use:     "push-latest",
		Aliases: []string{"push"},
		Short: "Push the latest command from the history to the DB (local file). " +
			"input the name of the snippet as an argument",
		Args:         cobra.ExactArgs(1),
		Example:      "snipository push-latest get-pods",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			_ = cmd.Context()

			inputSnippetName := args[0]
			if inputSnippetName == "" {
				return fmt.Errorf("snippet name cannot be empty")
			}

			latestCommand, err := history.GetLatestFromHistory(newConfig.HistoryEnvVarName)
			if err != nil {
				return err
			}

			snippet := common.Snippet{
				Name:    inputSnippetName,
				Command: latestCommand,
			}

			err = storage.PushSnippetToFile(snippet, newConfig.DataFilePath)
			if err != nil {
				return err
			}

			return nil
		},
	}

	return c
}
