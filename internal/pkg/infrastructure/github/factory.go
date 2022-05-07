package github

import (
	"net/http"

	"github.com/google/go-github/v44/github"
)

func NewRestClient(httpClient *http.Client) *github.Client {
	return github.NewClient(httpClient)
}
