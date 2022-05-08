package cmd

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
	"github.com/valeriobelli/gh-branch/internal/pkg/application/rebase"
	"github.com/valeriobelli/gh-branch/internal/pkg/domain/slices"
)

func NewRebase() *cobra.Command {
	defaultRebaseType := "true"
	validRebaseTypes := slices.NewStringSlice([]string{defaultRebaseType, "false", "interactive", "merges"})

	var getRebaseType = func(command *cobra.Command) string {
		branchType, err := command.Flags().GetString("type")

		if err != nil || !validRebaseTypes.Contains(branchType) {
			return "feature"
		}

		return branchType
	}

	createCommand := &cobra.Command{
		Use:   "rebase",
		Short: "Rebase the current Pull Request's branch",
		Long: heredoc.Doc(`
			Rebase the current Pull Request's branch.

			Optionally, the rebase type can be specified.
		`),
		Run: func(command *cobra.Command, args []string) {
			rebase.Rebase(getRebaseType(command))
		},
	}

	createCommand.Flags().StringP("type", "t", defaultRebaseType, fmt.Sprintf("Define the type of rebase [%s]", validRebaseTypes.Join(", ")))

	return createCommand
}

func init() {
	rootCommand.AddCommand(NewRebase())
}
