package rebase

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/fatih/color"
	"github.com/valeriobelli/gh-branch/internal/pkg/infrastructure/gh"
)

func Rebase(rebaseType string) {
	pullRequestInfo, err := gh.RetrievePullRequestInformation()

	if err != nil {
		fmt.Printf("%s %s\n", color.RedString("✘"), err.Error())

		return
	}

	var stdErr bytes.Buffer

	cmd := exec.Command("git", "pull", "origin", pullRequestInfo.BaseBranch, "--rebase", rebaseType)

	cmd.Stderr = &stdErr

	if err = cmd.Run(); err != nil {
		fmt.Printf("%s %s\n", color.RedString("✘"), stdErr.String())

		return
	}

	fmt.Printf("%s Branch rebased.\n", color.GreenString("✓"))
}
