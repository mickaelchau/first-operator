/*
Copyright 2024.

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

package controller

import (
	"context"
	webappv1 "github.com/mickaelchau/first-operator/api/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// MickaDeploymentReconciler reconciles a MickaDeployment object
type MickaDeploymentReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=webapp.micka.com,resources=mickadeployments,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=webapp.micka.com,resources=mickadeployments/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=webapp.micka.com,resources=mickadeployments/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the MickaDeployment object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.16.3/pkg/reconcile
func (r *MickaDeploymentReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	// Log request
	log := log.FromContext(ctx)
	log.Info("Reconciling MickaDeployment")

	// Fetch the MickaDeployment instance
	instance := &webappv1.MickaDeployment{}
	err := r.Get(ctx, req.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Object not found, return.  Created objects are automatically garbage collected.
			// For additional cleanup logic use finalizers.
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// Check if the Foo field is already set to "test"
	if instance.Spec.Foo != "test" {
		// Update the Foo field
		instance.Spec.Foo = "test"
		err := r.Update(ctx, instance)
		if err != nil {
			log.Error(err, "Failed to update MickaDeployment instance")
			return reconcile.Result{}, err
		}
		log.Info("Updated Foo field to 'test'")
	} else {
		log.Info("Foo field is already set to 'test', no update required")
	}

	// No requeue, we only requeue if the Foo field was not already set to "test"
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *MickaDeploymentReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&webappv1.MickaDeployment{}).
		Complete(r)
}
