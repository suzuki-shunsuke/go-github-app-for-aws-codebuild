package ghapputil

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/codebuild"
)

type CodeBuild interface {
	StartBuildWithContext(ctx aws.Context, input *codebuild.StartBuildInput, opts ...request.Option) (*codebuild.StartBuildOutput, error)
}
