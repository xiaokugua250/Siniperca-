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
	"log"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	webappv1 "Siniperca/api/v1"
)

// SiteholdReconciler reconciles a Sitehold object
type SiteholdReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=webapp.z-gour.com,resources=siteholdren,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=webapp.z-gour.com,resources=siteholdren/status,verbs=get;update;patch

func (r *SiteholdReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("sitehold", req.NamespacedName)

	// your logic here
	ctx := context.Background()
	_ = r.Log.WithValues("apiexample", req.NamespacedName)

	obj := &webappv1.Sitehold{}
	if err := r.Get(ctx, req.NamespacedName, obj); err != nil {
		log.Println(err, "unable to fetch Obj")
	} else {
		log.Println("greeting from kubebuilder to", obj.Spec.DBService, obj.Spec.MicroServices)
	}
	obj.Status.Status = "Running"
	if err := r.Status().Update(ctx, obj); err != nil {
		log.Println("unable to update status....")

	}
	return ctrl.Result{}, nil
}

func (r *SiteholdReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&webappv1.Sitehold{}).
		Complete(r)
}
