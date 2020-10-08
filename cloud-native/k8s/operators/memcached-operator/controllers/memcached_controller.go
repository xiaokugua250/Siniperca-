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
	appv1alpha1 "github.com/example-org/app-operator/api/v1alpha1"
	"github.com/go-logr/logr"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	cachev1 "github.com/Siniperca/memcached-operator/api/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const memcachedFinalizer = "finalizer.cache.example.com"

// MemcachedReconciler reconciles a Memcached object
type MemcachedReconciler struct {
	// client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client.Client
	Log logr.Logger
	// scheme defines methods for serializing and deserializing API objects,
	// a type registry for converting group, version, and kind information
	// to and from Go schemas, and mappings between Go schemas of different
	// versions. A scheme is the foundation for a versioned API and versioned
	// configuration over time.
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=cache.siniperca.cloudnative.com,resources=memcacheds,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cache.siniperca.cloudnative.com,resources=memcacheds/status,verbs=get;update;patch

func (r *MemcachedReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	_ = r.Log.WithValues("memcached", req.NamespacedName)
	//Logger.Info("Reconciling Memcached")

	memcached := &cachev1.Memcached{}
	err := r.Get(ctx, req.NamespacedName, memcached)
	// your logic here
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			//reqLogger.Info("Memcached resource not found. Ignoring since object must be deleted.")
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request.
		//reqLogger.Error(err, "Failed to get Memcached.")
		return ctrl.Result{}, err
	}
	// Check if the Memcached instance is marked to be deleted, which is
	// indicated by the deletion timestamp being set.
	isMemcachedMarkedToBeDeleted := memcached.GetDeletionTimestamp() != nil
	if isMemcachedMarkedToBeDeleted {
		if contains(memcached.GetFinalizers(), memcachedFinalizer) {
			// Run finalization logic for memcachedFinalizer. If the
			// finalization logic fails, don't remove the finalizer so
			// that we can retry during the next reconciliation.
			if err := r.finalizeMemcached(reqLogger, memcached); err != nil {
				return ctrl.Result{}, err
			}
			// Remove memcachedFinalizer. Once all finalizers have been
			// removed, the object will be deleted.
			controllerutil.RemoveFinalizer(memcached, memcachedFinalizer)
			err := r.Update(ctx, memcached)
			if err != nil {
				return ctrl.Result{}, err
			}
		}
	}
	// Add finalizer for this CR
	if !contains(memcached.GetFinalizers(), memcachedFinalizer) {
		if err := r.addFinalizer(reqLogger, memcached); err != nil {
			return ctrl.Result{}, err
		}
	}
	return ctrl.Result{}, nil
}

// deploymentForApp returns a app Deployment object.
func (r *MemcachedReconciler) deploymentForApp(m *appv1alpha1.App) *v1.Deployment {
	lbls := labelsForApp(m.Name)
	replicas := m.Spec.Size

	dep := &v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      m.Name,
			Namespace: m.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: lbls,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: lbls,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image:   "app:alpine",
						Name:    "app",
						Command: []string{"app", "-a=64", "-b"},
						Ports: []corev1.ContainerPort{{
							ContainerPort: 10000,
							Name:          "app",
						}},
					}},
				},
			},
		},
	}

	// Set App instance as the owner and controller.
	// NOTE: calling SetControllerReference, and setting owner references in
	// general, is important as it allows deleted objects to be garbage collected.
	controllerutil.SetControllerReference(m, dep, r.scheme)
	return dep
}
func (r *MemcachedReconciler) finalizeMemcached(reqLogger logr.Logger, m *cachev1alpha1.Memcached) error {
	// TODO(user): Add the cleanup steps that the operator
	// needs to do before the CR can be deleted. Examples
	// of finalizers include performing backups and deleting
	// resources that are not owned by this CR, like a PVC.
	reqLogger.Info("Successfully finalized memcached")
	return nil
}

func (r *MemcachedReconciler) addFinalizer(reqLogger logr.Logger, m *cachev1alpha1.Memcached) error {
	reqLogger.Info("Adding Finalizer for the Memcached")
	controllerutil.AddFinalizer(m, memcachedFinalizer)

	// Update CR
	err := r.Update(context.TODO(), m)
	if err != nil {
		reqLogger.Error(err, "Failed to update Memcached with finalizer")
		return err
	}
	return nil
}

func contains(list []string, s string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}
func (r *MemcachedReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&cachev1.Memcached{}).
		Owns(&v1.Deployment{}).
		WithOptions(controller.Options{
			MaxConcurrentReconciles: 2,
		}).
		Complete(r)
}

func ignoreDeletionPredicate() predicate.Predicate {
	return predicate.Funcs{
		UpdateFunc: func(event event.UpdateEvent) bool {
			return event.MetaOld.GetGeneration() != event.MetaNew.GetGeneration()
		},
		DeleteFunc: func(deleteEvent event.DeleteEvent) bool {
			return !deleteEvent.DeleteStateUnknown
		},
	}
}
