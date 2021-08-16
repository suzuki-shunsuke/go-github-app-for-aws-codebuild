package ghapputil

import (
	"fmt"
	"net/http"

	"github.com/bradleyfalzon/ghinstallation"
	"github.com/google/go-github/v38/github"
)

func NewGitHubClient(appID int64, event *github.PullRequestEvent, keyFile []byte) (*github.Client, error) {
	inst := event.GetInstallation()
	itr, err := ghinstallation.New(http.DefaultTransport, appID, inst.GetID(), keyFile)
	if err != nil {
		return nil, fmt.Errorf("create a transport with private key: %w", err)
	}
	return github.NewClient(&http.Client{Transport: itr}), nil
}
