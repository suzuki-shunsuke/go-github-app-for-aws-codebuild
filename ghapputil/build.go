package ghapputil

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codebuild"
)

type CodeBuild interface {
	StartBuildWithContext(ctx aws.Context, input *codebuild.StartBuildInput, opts ...request.Option) (*codebuild.StartBuildOutput, error)
}

type Builder struct {
	clients map[string]CodeBuild
	session *session.Session
}

func NewBuilder(sess *session.Session) *Builder {
	return &Builder{
		clients: map[string]CodeBuild{},
		session: sess,
	}
}

func (builder *Builder) newClient(input *StartBuildInput) CodeBuild {
	if input.AssumeRoleARN == "" && input.Region == "" {
		return codebuild.New(builder.session)
	}
	cfg := &aws.Config{}
	if input.AssumeRoleARN != "" {
		cfg.Credentials = stscreds.NewCredentials(builder.session, input.AssumeRoleARN)
	}
	if input.Region != "" {
		cfg.Region = aws.String(input.Region)
	}
	return codebuild.New(builder.session, cfg)
}

func (builder *Builder) getClient(input *StartBuildInput) CodeBuild {
	key := input.AssumeRoleARN + "/" + input.Region
	client, ok := builder.clients[key]
	if ok {
		return client
	}
	client = builder.newClient(input)
	builder.clients[key] = client // Note that this function isn't thread safe.
	return client
}

func (builder *Builder) Start(ctx aws.Context, input *StartBuildInput) (*codebuild.StartBuildOutput, error) {
	return builder.getClient(input).StartBuildWithContext(ctx, input.Input, input.Options...) //nolint:wrapcheck
}

type StartBuildInput struct {
	Input         *codebuild.StartBuildInput
	Region        string
	AssumeRoleARN string
	Options       []request.Option
}

type StartBuildError struct {
	Input  *StartBuildInput
	Output *codebuild.StartBuildOutput
	Error  error
}

func StartBuild(ctx context.Context, builder *Builder, inputs []*StartBuildInput) []*StartBuildError {
	var startBuildErrors []*StartBuildError
	for _, input := range inputs {
		input := input
		output, err := builder.Start(ctx, input)
		if err != nil {
			startBuildErrors = append(startBuildErrors, &StartBuildError{
				Input:  input,
				Output: output,
				Error:  err,
			})
		}
	}

	return startBuildErrors
}
