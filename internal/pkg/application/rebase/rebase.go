package rebase

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/fatih/color"
	"github.com/valeriobelli/gh-branch/internal/pkg/infrastructure/gh"
	"github.com/valeriobelli/gh-branch/internal/pkg/infrastructure/spinner"
)

func Rebase(rebaseType string) {
	spinner := spinner.NewSpinner()

	spinner.Start()

	pullRequestInfo, err := gh.RetrievePullRequestInformation()

	if err != nil {
		spinner.Stop()

		fmt.Printf("%s %s\n", color.RedString("✘"), err.Error())

		return
	}

	var stdErr bytes.Buffer

	cmd := exec.Command("git", "pull", fmt.Sprintf("--rebase=%s", rebaseType), "origin", pullRequestInfo.BaseBranch)

	cmd.Stderr = &stdErr

	if err = cmd.Run(); err != nil {
		spinner.Stop()

		fmt.Printf("%s %s\n", color.RedString("✘"), stdErr.String())

		return
	}

	spinner.Stop()

	fmt.Printf("%s Branch rebased.\n", color.GreenString("✓"))
}
