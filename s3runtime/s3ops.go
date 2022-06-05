package s3runtime

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3types "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/go-logr/logr"
	awsv1 "github.com/yuvalman/s3BucketController/api/v1"
	"time"
)

func Retryer() aws.Retryer {
	return retry.AddWithMaxAttempts(retry.AddWithMaxBackoffDelay(retry.NewStandard(), time.Second*100), 12)
}

func ConfigureS3Client(ctx context.Context) (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRetryer(Retryer))
	if err != nil {
		return nil, err
	}
	return s3.NewFromConfig(cfg), nil
}

type s3ops struct {
	s3c S3Client
}

// NewS3Ops create a default S3 operations helper
func NewS3Ops(
	s3c S3Client,
) S3ops {
	return &s3ops{
		s3c: s3c,
	}
}

func (t *s3ops) UpdatePublicAccessBlock(ctx context.Context, bucket *awsv1.S3Bucket, log logr.Logger) error {
	// find the endpoint for the bucket's region
	defres := s3.NewDefaultEndpointResolver()
	epr := s3.WithEndpointResolver(s3.EndpointResolverFunc(func(region string, options s3.EndpointResolverOptions) (aws.Endpoint, error) {
		return defres.ResolveEndpoint(bucket.Spec.Region, options)
	}))
	publicAccessBlock := &s3types.PublicAccessBlockConfiguration{
		BlockPublicAcls:       bucket.Spec.PublicAccessBlock.BlockPublicACLs,
		BlockPublicPolicy:     bucket.Spec.PublicAccessBlock.BlockPublicPolicy,
		IgnorePublicAcls:      bucket.Spec.PublicAccessBlock.IgnorePublicACLs,
		RestrictPublicBuckets: bucket.Spec.PublicAccessBlock.RestrictPublicBuckets,
	}
	log.Info("put public access block to bucket " + bucket.Spec.Name)
	_, err := t.s3c.PutPublicAccessBlock(ctx, &s3.PutPublicAccessBlockInput{Bucket: &bucket.Spec.Name, PublicAccessBlockConfiguration: publicAccessBlock}, epr)
	return err
}
