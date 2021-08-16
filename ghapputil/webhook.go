package ghapputil

import (
	"fmt"

	"github.com/google/go-github/v38/github"
)

func ValidateSignature(event *Event, webhookSecret []byte) error {
	if webhookSecret == nil {
		return nil
	}
	if err := github.ValidateSignature(event.Headers.Signature, []byte(event.Body), webhookSecret); err != nil {
		return fmt.Errorf("validate GitHub Webhook with signature: %w", err)
	}
	return nil
}

func ParseWebHook(event, body string) (interface{}, error) {
	payload, err := github.ParseWebHook(event, []byte(body))
	if err != nil {
		return nil, fmt.Errorf("parse webhook payload: %w", err)
	}
	return payload, nil
}

func ExcludePREventByAction(prEvent *github.PullRequestEvent) bool {
	switch action := prEvent.GetAction(); action {
	case "opened", "synchronize", "reopened":
		return true
	case "closed":
		if prEvent.PullRequest.GetMerged() {
			return true
		}
	}
	return false
}
