package gh

import "strings"

type PullRequestInfo struct {
	BaseBranch string
	Branch     string
}

func RetrievePullRequestInformation() (PullRequestInfo, error) {
	branch, err := Execute([]string{"pr", "view", "--json", "headRefName", "--jq", ".headRefName"})
	baseBranch, err := Execute([]string{"pr", "view", "--json", "baseRefName", "--jq", ".baseRefName"})

	return PullRequestInfo{BaseBranch: strings.TrimSpace(baseBranch), Branch: strings.TrimSpace(branch)}, err
}
