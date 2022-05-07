package create

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/gosimple/slug"
	"github.com/valeriobelli/gh-branch/internal/pkg/infrastructure/gh"
	"github.com/valeriobelli/gh-branch/internal/pkg/infrastructure/github"
	"github.com/valeriobelli/gh-branch/internal/pkg/infrastructure/http"
	"github.com/valeriobelli/gh-branch/internal/pkg/infrastructure/spinner"
)

type CreateConfig struct {
	Type string
}

type Create struct {
	config CreateConfig
}

func NewCreate(config CreateConfig) *Create {
	return &Create{
		config: config,
	}
}

func (cb Create) getIssueNumber(issueOrIssueNumber string) (*int, error) {
	issueNumber, err := strconv.Atoi(issueOrIssueNumber)

	if err == nil {
		return &issueNumber, nil
	}

	issueUrlRegexp := regexp.MustCompile("/.+/.+/issues/(?P<Year>\\d+)")

	match := issueUrlRegexp.FindStringSubmatch(issueOrIssueNumber)

	if len(match) == 2 {
		number, err := strconv.Atoi(match[1])

		if err != nil {
			return nil, err
		}

		return &number, nil
	}

	return nil, errors.New(fmt.Sprintf("Issue number not found in %s", issueOrIssueNumber))
}

func (cb Create) FromIssue(issueOrIssueNumber string) {
	issueNumber, err := cb.getIssueNumber(issueOrIssueNumber)

	if err != nil {
		fmt.Println(err.Error())

		return
	}

	repoInfo, err := gh.RetrieveRepoInformation()

	if err != nil {
		fmt.Println(err.Error())

		return
	}

	spinner := spinner.NewSpinner()
	githubClient := github.NewRestClient(http.NewClient())

	spinner.Start()

	issue, response, err := githubClient.Issues.Get(context.Background(), repoInfo.Owner, repoInfo.Name, *issueNumber)

	spinner.Stop()

	if err != nil {
		switch response.StatusCode {
		case 404:
			fmt.Printf("Issue #%d not found.\n", *issueNumber)
		default:
			fmt.Println("Unknown error.")
		}

		return
	}

	branch := fmt.Sprintf("%s/%d-%s", cb.config.Type, *issueNumber, slug.Make(issue.GetTitle()))

	var listOfBranchStdOut bytes.Buffer
	var listOfBranchStdErr bytes.Buffer

	cmd := exec.Command("git", "branch", "--list")

	cmd.Stderr = &listOfBranchStdErr
	cmd.Stdout = &listOfBranchStdOut

	if err = cmd.Run(); err != nil {
		fmt.Printf("%s %s\n", color.RedString("✘"), listOfBranchStdErr.String())

		return
	}

	if strings.Contains(listOfBranchStdOut.String(), branch) {
		fmt.Printf("%s A branch named '%s' already exists.\n", color.RedString("✘"), branch)

		return
	}

	var createBranchStdErr bytes.Buffer

	cmd = exec.Command("git", "switch", "-c", branch)

	cmd.Stderr = &createBranchStdErr

	if err = cmd.Run(); err != nil {
		fmt.Printf("%s %s", color.RedString("✘"), createBranchStdErr.String())

		return
	}

	fmt.Printf("%s Branch %s created.\n", color.GreenString("✓"), branch)
}
