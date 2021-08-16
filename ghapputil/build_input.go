package ghapputil

import (
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/codebuild"
	"github.com/google/go-github/v38/github"
)

func GetSourceVersion(prEvent *github.PullRequestEvent, isSHA bool) string {
	pr := prEvent.PullRequest
	if pr.GetMerged() {
		return pr.GetMergeCommitSHA()
	}
	if isSHA {
		return prEvent.GetAfter()
	}
	return "pr/" + strconv.Itoa(prEvent.GetNumber())
}

func GetPREnv(prEvent *github.PullRequestEvent) []*codebuild.EnvironmentVariable {
	labels := GetLabelNames(prEvent.PullRequest.Labels)
	repo := prEvent.Repo
	return []*codebuild.EnvironmentVariable{
		{
			Name:  aws.String("CODEBUILDER_REPO_OWNER"),
			Value: aws.String(repo.Owner.GetLogin()),
		},
		{
			Name:  aws.String("CODEBUILDER_REPO_NAME"),
			Value: aws.String(repo.GetName()),
		},
		{
			Name:  aws.String("CODEBUILDER_PR_NUMBER"),
			Value: aws.String(strconv.Itoa(prEvent.GetNumber())),
		},
		{
			Name:  aws.String("CODEBUILDER_PR_AUTHOR"),
			Value: aws.String(prEvent.PullRequest.User.GetLogin()),
		},
		{
			Name:  aws.String("CODEBUILDER_PR_LABELS"),
			Value: aws.String(strings.Join(labels, "\n")),
		},
		{
			Name:  aws.String("CODEBUILDER_EVENT_ACTION"),
			Value: aws.String(prEvent.GetAction()),
		},
		{
			Name:  aws.String("CODEBUILDER_PR_MERGED"),
			Value: aws.String(strconv.FormatBool(prEvent.PullRequest.GetMerged())),
		},
	}
}

func GetGitHubAppEnv(appID, instID int64) []*codebuild.EnvironmentVariable {
	return []*codebuild.EnvironmentVariable{
		{
			Name:  aws.String("CODEBUILDER_GITHUB_APP_APP_ID"),
			Value: aws.String(strconv.FormatInt(appID, 10)), //nolint:gomnd
		},
		{
			Name:  aws.String("CODEBUILDER_GITHUB_APP_INSTALLATION_ID"),
			Value: aws.String(strconv.FormatInt(instID, 10)), //nolint:gomnd
		},
	}
}
