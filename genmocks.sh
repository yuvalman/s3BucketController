set -e

mockgen -package s3_mocks \
-destination s3runtime/mocks/types_mocks.go \
-source s3runtime/types.go
