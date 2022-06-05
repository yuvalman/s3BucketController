package s3runtime

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/go-logr/logr"
	awsv1 "github.com/yuvalman/s3BucketController/api/v1"
)

// S3Client interface for s3.Client
type S3Client interface {
	PutPublicAccessBlock(ctx context.Context, params *s3.PutPublicAccessBlockInput, optFns ...func(*s3.Options)) (*s3.PutPublicAccessBlockOutput, error)
}

type S3ops interface {
	UpdatePublicAccessBlock(ctx context.Context, bucket *awsv1.S3Bucket, log logr.Logger) error
}
