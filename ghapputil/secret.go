package ghapputil

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

type Secret struct {
	WebhookSecret       string `json:"webhook_secret"`
	GitHubAppPrivateKey string `json:"github_app_private_key"`
}

type SecretsManager interface {
	GetSecretValueWithContext(ctx aws.Context, input *secretsmanager.GetSecretValueInput, opts ...request.Option) (*secretsmanager.GetSecretValueOutput, error)
}

func ReadSecretFromSecretsManager(ctx context.Context, svc SecretsManager, input *secretsmanager.GetSecretValueInput, secret interface{}) error {
	output, err := svc.GetSecretValueWithContext(ctx, input)
	if err != nil {
		return fmt.Errorf("get secret value from AWS SecretsManager: %w", err)
	}
	if err := json.Unmarshal([]byte(*output.SecretString), secret); err != nil {
		return fmt.Errorf("parse secret value: %w", err)
	}
	return nil
}
