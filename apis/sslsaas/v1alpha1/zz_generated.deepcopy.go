// +build !ignore_autogenerated

/*
Copyright 2020 The Crossplane Authors.

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
	cloudflare_go "github.com/cloudflare/cloudflare-go"
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomHostname) DeepCopyInto(out *CustomHostname) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomHostname.
func (in *CustomHostname) DeepCopy() *CustomHostname {
	if in == nil {
		return nil
	}
	out := new(CustomHostname)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomHostname) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomHostnameList) DeepCopyInto(out *CustomHostnameList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CustomHostname, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomHostnameList.
func (in *CustomHostnameList) DeepCopy() *CustomHostnameList {
	if in == nil {
		return nil
	}
	out := new(CustomHostnameList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomHostnameList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomHostnameObservation) DeepCopyInto(out *CustomHostnameObservation) {
	*out = *in
	if in.VerificationErrors != nil {
		in, out := &in.VerificationErrors, &out.VerificationErrors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.SSL.DeepCopyInto(&out.SSL)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomHostnameObservation.
func (in *CustomHostnameObservation) DeepCopy() *CustomHostnameObservation {
	if in == nil {
		return nil
	}
	out := new(CustomHostnameObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomHostnameOwnershipVerification) DeepCopyInto(out *CustomHostnameOwnershipVerification) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomHostnameOwnershipVerification.
func (in *CustomHostnameOwnershipVerification) DeepCopy() *CustomHostnameOwnershipVerification {
	if in == nil {
		return nil
	}
	out := new(CustomHostnameOwnershipVerification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomHostnameParameters) DeepCopyInto(out *CustomHostnameParameters) {
	*out = *in
	if in.Hostname != nil {
		in, out := &in.Hostname, &out.Hostname
		*out = new(string)
		**out = **in
	}
	in.SSL.DeepCopyInto(&out.SSL)
	if in.CustomOriginServer != nil {
		in, out := &in.CustomOriginServer, &out.CustomOriginServer
		*out = new(string)
		**out = **in
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
	if in.ZoneRef != nil {
		in, out := &in.ZoneRef, &out.ZoneRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.ZoneSelector != nil {
		in, out := &in.ZoneSelector, &out.ZoneSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomHostnameParameters.
func (in *CustomHostnameParameters) DeepCopy() *CustomHostnameParameters {
	if in == nil {
		return nil
	}
	out := new(CustomHostnameParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomHostnameSSL) DeepCopyInto(out *CustomHostnameSSL) {
	*out = *in
	if in.Method != nil {
		in, out := &in.Method, &out.Method
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	in.Settings.DeepCopyInto(&out.Settings)
	if in.Wildcard != nil {
		in, out := &in.Wildcard, &out.Wildcard
		*out = new(bool)
		**out = **in
	}
	if in.CustomCertificate != nil {
		in, out := &in.CustomCertificate, &out.CustomCertificate
		*out = new(string)
		**out = **in
	}
	if in.CustomKey != nil {
		in, out := &in.CustomKey, &out.CustomKey
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomHostnameSSL.
func (in *CustomHostnameSSL) DeepCopy() *CustomHostnameSSL {
	if in == nil {
		return nil
	}
	out := new(CustomHostnameSSL)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomHostnameSSLObserved) DeepCopyInto(out *CustomHostnameSSLObserved) {
	*out = *in
	if in.ValidationErrors != nil {
		in, out := &in.ValidationErrors, &out.ValidationErrors
		*out = make([]cloudflare_go.CustomHostnameSSLValidationErrors, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomHostnameSSLObserved.
func (in *CustomHostnameSSLObserved) DeepCopy() *CustomHostnameSSLObserved {
	if in == nil {
		return nil
	}
	out := new(CustomHostnameSSLObserved)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomHostnameSSLSettings) DeepCopyInto(out *CustomHostnameSSLSettings) {
	*out = *in
	if in.HTTP2 != nil {
		in, out := &in.HTTP2, &out.HTTP2
		*out = new(string)
		**out = **in
	}
	if in.TLS13 != nil {
		in, out := &in.TLS13, &out.TLS13
		*out = new(string)
		**out = **in
	}
	if in.MinTLSVersion != nil {
		in, out := &in.MinTLSVersion, &out.MinTLSVersion
		*out = new(string)
		**out = **in
	}
	if in.Ciphers != nil {
		in, out := &in.Ciphers, &out.Ciphers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomHostnameSSLSettings.
func (in *CustomHostnameSSLSettings) DeepCopy() *CustomHostnameSSLSettings {
	if in == nil {
		return nil
	}
	out := new(CustomHostnameSSLSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomHostnameSSLValidationErrors) DeepCopyInto(out *CustomHostnameSSLValidationErrors) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomHostnameSSLValidationErrors.
func (in *CustomHostnameSSLValidationErrors) DeepCopy() *CustomHostnameSSLValidationErrors {
	if in == nil {
		return nil
	}
	out := new(CustomHostnameSSLValidationErrors)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomHostnameSpec) DeepCopyInto(out *CustomHostnameSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomHostnameSpec.
func (in *CustomHostnameSpec) DeepCopy() *CustomHostnameSpec {
	if in == nil {
		return nil
	}
	out := new(CustomHostnameSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomHostnameStatus) DeepCopyInto(out *CustomHostnameStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomHostnameStatus.
func (in *CustomHostnameStatus) DeepCopy() *CustomHostnameStatus {
	if in == nil {
		return nil
	}
	out := new(CustomHostnameStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FallbackOrigin) DeepCopyInto(out *FallbackOrigin) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FallbackOrigin.
func (in *FallbackOrigin) DeepCopy() *FallbackOrigin {
	if in == nil {
		return nil
	}
	out := new(FallbackOrigin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FallbackOrigin) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FallbackOriginList) DeepCopyInto(out *FallbackOriginList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FallbackOrigin, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FallbackOriginList.
func (in *FallbackOriginList) DeepCopy() *FallbackOriginList {
	if in == nil {
		return nil
	}
	out := new(FallbackOriginList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FallbackOriginList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FallbackOriginObservation) DeepCopyInto(out *FallbackOriginObservation) {
	*out = *in
	if in.Errors != nil {
		in, out := &in.Errors, &out.Errors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FallbackOriginObservation.
func (in *FallbackOriginObservation) DeepCopy() *FallbackOriginObservation {
	if in == nil {
		return nil
	}
	out := new(FallbackOriginObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FallbackOriginParameters) DeepCopyInto(out *FallbackOriginParameters) {
	*out = *in
	if in.Origin != nil {
		in, out := &in.Origin, &out.Origin
		*out = new(string)
		**out = **in
	}
	if in.OriginRef != nil {
		in, out := &in.OriginRef, &out.OriginRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.OriginSelector != nil {
		in, out := &in.OriginSelector, &out.OriginSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
	if in.ZoneRef != nil {
		in, out := &in.ZoneRef, &out.ZoneRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.ZoneSelector != nil {
		in, out := &in.ZoneSelector, &out.ZoneSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FallbackOriginParameters.
func (in *FallbackOriginParameters) DeepCopy() *FallbackOriginParameters {
	if in == nil {
		return nil
	}
	out := new(FallbackOriginParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FallbackOriginSpec) DeepCopyInto(out *FallbackOriginSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FallbackOriginSpec.
func (in *FallbackOriginSpec) DeepCopy() *FallbackOriginSpec {
	if in == nil {
		return nil
	}
	out := new(FallbackOriginSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FallbackOriginStatus) DeepCopyInto(out *FallbackOriginStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FallbackOriginStatus.
func (in *FallbackOriginStatus) DeepCopy() *FallbackOriginStatus {
	if in == nil {
		return nil
	}
	out := new(FallbackOriginStatus)
	in.DeepCopyInto(out)
	return out
}
