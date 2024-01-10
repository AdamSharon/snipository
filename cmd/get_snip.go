package main

import (
	"fmt"

	"snipository/common"
	"snipository/config"
	"snipository/storage"

	"github.com/atotto/clipboard"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/spf13/cobra"
)

func newSnipositoryCommand() *cobra.Command {
	newConfig := config.NewConfig()

	c := &cobra.Command{
		Use:     "get-snip",
		Aliases: []string{"cp", "get"},
		Short:   "get a snippet and copy it to the clipboard. use TAB to autocomplete the snippet name",
		Long: "get a snippet and copy it to the clipboard. use TAB to autocomplete the snippet name. " +
			"the snippet name is the first argument, and it works with fuzzy search." +
			"this command meant to work with the zsh completion script, so if you enter a command that doesn't exist, " + "" +
			"you might get an error.",
		Args:         cobra.ExactArgs(1),
		Example:      "",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			_ = cmd.Context()

			inputSnippetName := args[0]
			snippets, err := getSnippetsResponse(newConfig.DataFilePath)

			snippetsMap := common.ConvertToMap(snippets)
			snippetCommand, ok := snippetsMap[inputSnippetName]
			if !ok {
				return fmt.Errorf("snippet not found in the DB")
			}

			err = clipboard.WriteAll(snippetCommand)
			if err != nil {
				return err
			}

			fmt.Println("\nThe command has been copied to the clipboard. You can paste it anywhere.")

			return nil
		},
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string,
		) ([]string, cobra.ShellCompDirective) {
			snippets, err := getSnippetsResponse(newConfig.DataFilePath)
			if err != nil {
				return nil, cobra.ShellCompDirectiveNoFileComp
			}

			snippetsMap := common.ConvertToMap(snippets)

			snippetSelectors := make([]string, 0, len(snippetsMap))
			for k, v := range snippetsMap {
				snippetSelection := fmt.Sprintf("%s\t%s", k, v)
				snippetSelectors = append(snippetSelectors, snippetSelection)
			}

			return fuzzy.Find(toComplete, snippetSelectors), cobra.ShellCompDirectiveNoFileComp
		},
	}

	return c
}

func getSnippetsResponse(path string) (*common.SnippetsResponse, error) {
	data, err := storage.GetDataFromFile(path)
	if err != nil {
		return nil, err
	}

	snippets, err := common.GetSnippetResponse(data)
	if err != nil {
		return nil, err
	}

	return snippets, nil
}
