package main

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	ctx := context.Background()
	if err := execute(ctx); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func execute(ctx context.Context) error {
	rootCmd := newRootCmd()
	rootCmd, err := rootCmd.ExecuteContextC(ctx)
	if err != nil {
		return err
	}
	return nil
}

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "snipository",
		Short: "Snipository is a CLI tool for managing snippets.",
		Long: "Snipository is a CLI tool for managing snippets. \n" +
			"snipository uses a local file to store snippets, at ~/.config/snipository/snipository_data.json. " +
			"please make sure the file exists before using the CLI. \n" +
			"you can use the following command to create the file: \n" +
			"mkdir -p ~/.config/snipository && touch ~/.config/snipository/snipository_data.json",
		SilenceErrors: true,
	}

	cmd.AddCommand(newSnipositoryCommand())
	cmd.AddCommand(zshCompletionCommand())
	cmd.AddCommand(pushLatestCommand())

	return cmd
}
