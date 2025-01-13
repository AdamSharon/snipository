package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"snipository/config"
)

func getSaveFile() *cobra.Command {
	newConfig := config.NewConfig()
	c := &cobra.Command{
		Use:          "get-file",
		Short:        "get the local file path",
		Args:         cobra.ExactArgs(0),
		Example:      "snipository get-file",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			_ = cmd.Context()
			fmt.Println(newConfig.DataFilePath)
			return nil
		},
	}

	return c
}
