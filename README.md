# s3BucketController

This is a tutorial of how to build a controller in k8s.

Note: For your convenient, every step that should be done is a commit in this repo.
## Overview
AWS S3 bucket controller for keeping desired state of public access permission for the bucket.

## Prerequisites

- The controller should be built on Linux.
- kubebuilder should be [installed](https://book.kubebuilder.io/quick-start.html#installation) and it's [prerequisites](https://book.kubebuilder.io/quick-start.html#prerequisites).
- For generating s3 ops mocks for unitest, mockgen should be [installed](https://github.com/golang/mock#installation)
- For communicating with s3 client, we should pass a secret containing AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY to a user that has the permission to update s3 buckets. For your secret, you should replace "your-aws-access-key" and "your-aws-secret-access-key" in config/manager/aws-secret.yaml
