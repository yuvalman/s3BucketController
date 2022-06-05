/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"github.com/yuvalman/s3BucketController/s3runtime"
	"k8s.io/apimachinery/pkg/api/errors"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	awsv1 "github.com/yuvalman/s3BucketController/api/v1"
)

// S3BucketReconciler reconciles a S3Bucket object
type S3BucketReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	S3Ops  s3runtime.S3ops
}

//+kubebuilder:rbac:groups=aws.services.io,resources=s3buckets,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=aws.services.io,resources=s3buckets/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=aws.services.io,resources=s3buckets/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the S3Bucket object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.2/pkg/reconcile
func (r *S3BucketReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	l := log.FromContext(ctx)
	var bucket awsv1.S3Bucket
	if err := r.Get(ctx, req.NamespacedName, &bucket); err != nil {
		if !errors.IsNotFound(err) {
			l.Error(err, "unable to fetch bucket")
		}
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	l.Info("=== Reconciling s3Bucket resource: " + bucket.Name)

	if err := r.S3Ops.UpdatePublicAccessBlock(ctx, &bucket, l); err != nil {
		return ctrl.Result{RequeueAfter: 1 * time.Second}, err
	}
	return ctrl.Result{RequeueAfter: 10 * time.Second}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *S3BucketReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&awsv1.S3Bucket{}).
		Complete(r)
}
