/*
Copyright 2023.

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

// api/v1/nimbleopti_webhook.go

package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// log is for logging in this package.
var nimbleoptilog = logf.Log.WithName("nimbleopti-resource")

func (r *NimbleOpti) SetupWebhookWithManager(mgr ctrl.Manager) error {
	// debug
	klog.Info("debug - SetupWebhookWithManager")

	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-adapter-uri-tech-github-io-v1-nimbleopti,mutating=true,failurePolicy=fail,sideEffects=None,groups=adapter.uri-tech.github.io,resources=nimbleoptis,verbs=create;update,versions=v1,name=mnimbleopti.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &NimbleOpti{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *NimbleOpti) Default() {
	// debug
	klog.Info("debug - Default")

	nimbleoptilog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-adapter-uri-tech-github-io-v1-nimbleopti,mutating=false,failurePolicy=fail,sideEffects=None,groups=adapter.uri-tech.github.io,resources=nimbleoptis,verbs=create;update,versions=v1,name=vnimbleopti.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &NimbleOpti{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *NimbleOpti) ValidateCreate() (admission.Warnings, error) {
	nimbleoptilog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil, nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *NimbleOpti) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	// debug
	klog.Info("debug - ValidateUpdate")

	nimbleoptilog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil, nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *NimbleOpti) ValidateDelete() (admission.Warnings, error) {
	// debug
	klog.Info("debug - ValidateDelete")

	nimbleoptilog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil, nil
}
