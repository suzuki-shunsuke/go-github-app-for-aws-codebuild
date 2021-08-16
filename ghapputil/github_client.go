package ghapputil

import (
	"context"
	"fmt"

	"github.com/google/go-github/v38/github"
)

const maxPerPage = 100

func GetPRFiles(ctx context.Context, client *github.Client, owner, repo string, prNumber, fileSize int) ([]*github.CommitFile, error) {
	ret := []*github.CommitFile{}
	if fileSize == 0 {
		return nil, nil
	}
	var n int
	if fileSize < 0 {
		n = 30
	} else {
		n = (fileSize / maxPerPage) + 1
		if n > 30 { //nolint:gomnd
			// https://docs.github.com/en/rest/reference/pulls#list-pull-requests-files
			// > Note: Responses include a maximum of 3000 files.
			n = 30
		}
	}
	for i := 1; i <= n; i++ {
		opts := &github.ListOptions{
			Page:    i,
			PerPage: maxPerPage,
		}
		files, _, err := client.PullRequests.ListFiles(ctx, owner, repo, prNumber, opts)
		if err != nil {
			return files, fmt.Errorf("get pull request files (page: %d, per_page: %d): %w", opts.Page, opts.PerPage, err)
		}
		ret = append(ret, files...)
		if len(files) != maxPerPage {
			return ret, nil
		}
	}

	return ret, nil
}

func GetLabelNames(labels []*github.Label) []string {
	arr := make([]string, len(labels))
	for i, label := range labels {
		arr[i] = label.GetName()
	}
	return arr
}
