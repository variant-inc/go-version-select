package cmd

import (
	"context"
	"fmt"
	"go-version-select/internal/handlers"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-version-select",
	Short: "Selects the latest version matching a given constraint",
	Long: `go-version-select is a command-line tool that helps you determine the latest version
that satisfies a specified version constraint.

You provide a list of available versions and a constraint (e.g., "^1.0.0"),
and the tool selects the most recent compatible version.`,
	Run: func(cmd *cobra.Command, _ []string) {
		versionList, err := cmd.Flags().GetString("versions")
		if err != nil {
			fmt.Fprintf(os.Stderr, fmt.Sprintf("error occurred while handling versions flag %+v", err))
			os.Exit(1)
		}

		constraint, err := cmd.Flags().GetString("constraint")
		if err != nil {
			fmt.Fprintf(os.Stderr, fmt.Sprintf("error occurred while handling constraint flag %+v", err))
			os.Exit(1)
		}

		if versionList == "" || constraint == "" {
			fmt.Fprintf(os.Stderr, "Both --versions and --constraint flags are required")
			os.Exit(1)
		}

		version, err := handlers.ProcessVersions(context.Background(), versionList, constraint)
		if err != nil {
			os.Exit(1)
		}
		fmt.Println(version)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().String("versions", "", "Comma-separated list of available versions")
	rootCmd.Flags().String("constraint", "", "Version constraint (e.g., ^1.0.0)")
}
