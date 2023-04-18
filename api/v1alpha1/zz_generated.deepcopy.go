//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NimbleOpticAdapterConfig) DeepCopyInto(out *NimbleOpticAdapterConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NimbleOpticAdapterConfig.
func (in *NimbleOpticAdapterConfig) DeepCopy() *NimbleOpticAdapterConfig {
	if in == nil {
		return nil
	}
	out := new(NimbleOpticAdapterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NimbleOpticAdapterConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NimbleOpticAdapterConfigList) DeepCopyInto(out *NimbleOpticAdapterConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NimbleOpticAdapterConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NimbleOpticAdapterConfigList.
func (in *NimbleOpticAdapterConfigList) DeepCopy() *NimbleOpticAdapterConfigList {
	if in == nil {
		return nil
	}
	out := new(NimbleOpticAdapterConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NimbleOpticAdapterConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NimbleOpticAdapterConfigSpec) DeepCopyInto(out *NimbleOpticAdapterConfigSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NimbleOpticAdapterConfigSpec.
func (in *NimbleOpticAdapterConfigSpec) DeepCopy() *NimbleOpticAdapterConfigSpec {
	if in == nil {
		return nil
	}
	out := new(NimbleOpticAdapterConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NimbleOpticAdapterConfigStatus) DeepCopyInto(out *NimbleOpticAdapterConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NimbleOpticAdapterConfigStatus.
func (in *NimbleOpticAdapterConfigStatus) DeepCopy() *NimbleOpticAdapterConfigStatus {
	if in == nil {
		return nil
	}
	out := new(NimbleOpticAdapterConfigStatus)
	in.DeepCopyInto(out)
	return out
}
