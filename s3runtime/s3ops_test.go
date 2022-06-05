package s3runtime_test

import (
	"context"
	"fmt"
	"github.com/go-logr/logr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3types "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	awsv1 "github.com/yuvalman/s3BucketController/api/v1"
	. "github.com/yuvalman/s3BucketController/s3runtime"
	. "github.com/yuvalman/s3BucketController/s3runtime/mocks"
)

func TestS3Ops(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "s3 ops tests")
}

var _ = Describe("S3Ops", func() {
	var mockc *gomock.Controller
	var s3c *MockS3Client
	var ops S3ops
	ctx := context.Background()
	err := fmt.Errorf("badError")
	var bucket awsv1.S3Bucket
	var publicAccessBlock *awsv1.PublicAccessBlockConfiguration
	var l logr.Logger

	BeforeEach(func() {
		mockc = gomock.NewController(GinkgoT())
		s3c = NewMockS3Client(mockc)
		ops = NewS3Ops(s3c)
		publicAccessBlock = &awsv1.PublicAccessBlockConfiguration{
			BlockPublicACLs:       true,
			BlockPublicPolicy:     true,
			IgnorePublicACLs:      true,
			RestrictPublicBuckets: true,
		}
		bucket = awsv1.S3Bucket{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "myBucketCrd",
				Namespace: "myNs",
			},
			Spec: awsv1.S3BucketSpec{
				Name:              "myBucket",
				Region:            "myRegion",
				PublicAccessBlock: publicAccessBlock,
			},
		}
		l = log.FromContext(ctx)
	})
	AfterEach(func() {
		mockc.Finish()
	})

	Describe("PutPublicAccessBlock Tests", func() {
		It("PutPublicAccessBlock - fail", func() {
			s3c.EXPECT().PutPublicAccessBlock(
				ctx,
				&s3.PutPublicAccessBlockInput{
					Bucket: &bucket.Spec.Name,
					PublicAccessBlockConfiguration: &s3types.PublicAccessBlockConfiguration{
						BlockPublicAcls:       bucket.Spec.PublicAccessBlock.BlockPublicACLs,
						BlockPublicPolicy:     bucket.Spec.PublicAccessBlock.BlockPublicPolicy,
						IgnorePublicAcls:      bucket.Spec.PublicAccessBlock.IgnorePublicACLs,
						RestrictPublicBuckets: bucket.Spec.PublicAccessBlock.RestrictPublicBuckets,
					},
				},
				&optFnMatcher{endpoint: "https://s3.myRegion.amazonaws.com"},
			).Return(nil, err)
			Expect(ops.UpdatePublicAccessBlock(ctx, &bucket, l)).To(HaveOccurred())
		})
		It("PutPublicAccessBlock - success", func() {
			s3c.EXPECT().PutPublicAccessBlock(
				ctx,
				&s3.PutPublicAccessBlockInput{
					Bucket: &bucket.Spec.Name,
					PublicAccessBlockConfiguration: &s3types.PublicAccessBlockConfiguration{
						BlockPublicAcls:       bucket.Spec.PublicAccessBlock.BlockPublicACLs,
						BlockPublicPolicy:     bucket.Spec.PublicAccessBlock.BlockPublicPolicy,
						IgnorePublicAcls:      bucket.Spec.PublicAccessBlock.IgnorePublicACLs,
						RestrictPublicBuckets: bucket.Spec.PublicAccessBlock.RestrictPublicBuckets,
					},
				},
				&optFnMatcher{endpoint: "https://s3.myRegion.amazonaws.com"},
			).Return(nil, nil)
			Expect(ops.UpdatePublicAccessBlock(ctx, &bucket, l)).To(Succeed())
		})
	})
})

// mock matcher for the opt function
type optFnMatcher struct {
	endpoint string
}

func (t *optFnMatcher) Matches(x interface{}) bool {
	if _, ok := x.(func(*s3.Options)); !ok {
		return false
	}
	optFunc := x.(func(*s3.Options))
	opt := &s3.Options{}
	optFunc(opt)
	if opt.EndpointResolver != nil {
		ep, err := opt.EndpointResolver.ResolveEndpoint("abc", s3.EndpointResolverOptions{})
		if err == nil {
			return ep.URL == t.endpoint
		}
	}

	return false
}
func (t *optFnMatcher) String() string {
	return "optfn epr " + t.endpoint
}
