/*


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

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	cnatv1alpha1 "cnat/api/v1alpha1"
)

// AtReconciler reconciles a At object
type AtReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=cnat.ishankhare.dev,resources=ats,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cnat.ishankhare.dev,resources=ats/status,verbs=get;update;patch

func (r *AtReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("at", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *AtReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&cnatv1alpha1.At{}).
		Complete(r)
}
