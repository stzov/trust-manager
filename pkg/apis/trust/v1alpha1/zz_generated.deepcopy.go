// +build !ignore_autogenerated

/*
Copyright 2021 The cert-manager Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bundle) DeepCopyInto(out *Bundle) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bundle.
func (in *Bundle) DeepCopy() *Bundle {
	if in == nil {
		return nil
	}
	out := new(Bundle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Bundle) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BundleCondition) DeepCopyInto(out *BundleCondition) {
	*out = *in
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BundleCondition.
func (in *BundleCondition) DeepCopy() *BundleCondition {
	if in == nil {
		return nil
	}
	out := new(BundleCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BundleList) DeepCopyInto(out *BundleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Bundle, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BundleList.
func (in *BundleList) DeepCopy() *BundleList {
	if in == nil {
		return nil
	}
	out := new(BundleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BundleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BundleSource) DeepCopyInto(out *BundleSource) {
	*out = *in
	if in.ConfigMap != nil {
		in, out := &in.ConfigMap, &out.ConfigMap
		*out = new(ObjectKeySelector)
		**out = **in
	}
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(ObjectKeySelector)
		**out = **in
	}
	if in.InLine != nil {
		in, out := &in.InLine, &out.InLine
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BundleSource.
func (in *BundleSource) DeepCopy() *BundleSource {
	if in == nil {
		return nil
	}
	out := new(BundleSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BundleSpec) DeepCopyInto(out *BundleSpec) {
	*out = *in
	if in.Sources != nil {
		in, out := &in.Sources, &out.Sources
		*out = make([]BundleSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Target.DeepCopyInto(&out.Target)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BundleSpec.
func (in *BundleSpec) DeepCopy() *BundleSpec {
	if in == nil {
		return nil
	}
	out := new(BundleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BundleStatus) DeepCopyInto(out *BundleStatus) {
	*out = *in
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(BundleTarget)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]BundleCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BundleStatus.
func (in *BundleStatus) DeepCopy() *BundleStatus {
	if in == nil {
		return nil
	}
	out := new(BundleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BundleTarget) DeepCopyInto(out *BundleTarget) {
	*out = *in
	if in.ConfigMap != nil {
		in, out := &in.ConfigMap, &out.ConfigMap
		*out = new(LocalKeySelector)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BundleTarget.
func (in *BundleTarget) DeepCopy() *BundleTarget {
	if in == nil {
		return nil
	}
	out := new(BundleTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalKeySelector) DeepCopyInto(out *LocalKeySelector) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalKeySelector.
func (in *LocalKeySelector) DeepCopy() *LocalKeySelector {
	if in == nil {
		return nil
	}
	out := new(LocalKeySelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectKeySelector) DeepCopyInto(out *ObjectKeySelector) {
	*out = *in
	out.LocalKeySelector = in.LocalKeySelector
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectKeySelector.
func (in *ObjectKeySelector) DeepCopy() *ObjectKeySelector {
	if in == nil {
		return nil
	}
	out := new(ObjectKeySelector)
	in.DeepCopyInto(out)
	return out
}
