package cmd

import (
	"errors"
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
	"github.com/valeriobelli/gh-branch/internal/pkg/application/create"
	"github.com/valeriobelli/gh-branch/internal/pkg/domain/slices"
)

func NewCreate() *cobra.Command {
	validBranchTypes := slices.NewStringSlice([]string{"feature", "feat", "hotfix", "fix", "release"})

	var getBranchType = func(command *cobra.Command) string {
		branchType, err := command.Flags().GetString("type")

		if err != nil || !validBranchTypes.Contains(branchType) {
			return "feature"
		}

		return branchType
	}

	createCommand := &cobra.Command{
		Use:   "create",
		Short: "Create a branch from an issue",
		Long: heredoc.Doc(`
			Create a branch starting from an issue.

			Optionally, the branch type can be specified.
		`),
		Run: func(command *cobra.Command, args []string) {
			create.NewCreate(create.CreateConfig{
				Type: getBranchType(command),
			}).FromIssue(args[0])
		},
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("This command expects an issue URL or an issue number to properly work.")
			}

			return nil
		},
	}

	createCommand.Flags().StringP("type", "t", "feature", fmt.Sprintf("Define the type of branch [%s]", validBranchTypes.Join(", ")))

	return createCommand
}

func init() {
	rootCommand.AddCommand(NewCreate())
}
