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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Zone) DeepCopyInto(out *Zone) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Zone.
func (in *Zone) DeepCopy() *Zone {
	if in == nil {
		return nil
	}
	out := new(Zone)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Zone) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneList) DeepCopyInto(out *ZoneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Zone, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneList.
func (in *ZoneList) DeepCopy() *ZoneList {
	if in == nil {
		return nil
	}
	out := new(ZoneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ZoneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneObservation) DeepCopyInto(out *ZoneObservation) {
	*out = *in
	if in.OriginalNS != nil {
		in, out := &in.OriginalNS, &out.OriginalNS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NameServers != nil {
		in, out := &in.NameServers, &out.NameServers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Betas != nil {
		in, out := &in.Betas, &out.Betas
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.VanityNameServers != nil {
		in, out := &in.VanityNameServers, &out.VanityNameServers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneObservation.
func (in *ZoneObservation) DeepCopy() *ZoneObservation {
	if in == nil {
		return nil
	}
	out := new(ZoneObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneParameters) DeepCopyInto(out *ZoneParameters) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(string)
		**out = **in
	}
	if in.JumpStart != nil {
		in, out := &in.JumpStart, &out.JumpStart
		*out = new(bool)
		**out = **in
	}
	if in.Paused != nil {
		in, out := &in.Paused, &out.Paused
		*out = new(bool)
		**out = **in
	}
	if in.PlanID != nil {
		in, out := &in.PlanID, &out.PlanID
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	in.Settings.DeepCopyInto(&out.Settings)
	if in.VanityNameServers != nil {
		in, out := &in.VanityNameServers, &out.VanityNameServers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneParameters.
func (in *ZoneParameters) DeepCopy() *ZoneParameters {
	if in == nil {
		return nil
	}
	out := new(ZoneParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneSettings) DeepCopyInto(out *ZoneSettings) {
	*out = *in
	if in.AlwaysOnline != nil {
		in, out := &in.AlwaysOnline, &out.AlwaysOnline
		*out = new(bool)
		**out = **in
	}
	if in.AdvancedDDOS != nil {
		in, out := &in.AdvancedDDOS, &out.AdvancedDDOS
		*out = new(bool)
		**out = **in
	}
	if in.AlwaysUseHTTPS != nil {
		in, out := &in.AlwaysUseHTTPS, &out.AlwaysUseHTTPS
		*out = new(bool)
		**out = **in
	}
	if in.AutomaticHTTPSRewrites != nil {
		in, out := &in.AutomaticHTTPSRewrites, &out.AutomaticHTTPSRewrites
		*out = new(bool)
		**out = **in
	}
	if in.Brotli != nil {
		in, out := &in.Brotli, &out.Brotli
		*out = new(bool)
		**out = **in
	}
	if in.BrowserCacheTTL != nil {
		in, out := &in.BrowserCacheTTL, &out.BrowserCacheTTL
		*out = new(int)
		**out = **in
	}
	if in.BrowserCheck != nil {
		in, out := &in.BrowserCheck, &out.BrowserCheck
		*out = new(bool)
		**out = **in
	}
	if in.CacheLevel != nil {
		in, out := &in.CacheLevel, &out.CacheLevel
		*out = new(string)
		**out = **in
	}
	if in.ChallengeTTL != nil {
		in, out := &in.ChallengeTTL, &out.ChallengeTTL
		*out = new(int)
		**out = **in
	}
	if in.CnameFlattening != nil {
		in, out := &in.CnameFlattening, &out.CnameFlattening
		*out = new(string)
		**out = **in
	}
	if in.DevelopmentMode != nil {
		in, out := &in.DevelopmentMode, &out.DevelopmentMode
		*out = new(bool)
		**out = **in
	}
	if in.EdgeCacheTTL != nil {
		in, out := &in.EdgeCacheTTL, &out.EdgeCacheTTL
		*out = new(int)
		**out = **in
	}
	if in.EmailObfuscation != nil {
		in, out := &in.EmailObfuscation, &out.EmailObfuscation
		*out = new(bool)
		**out = **in
	}
	if in.HotlinkProtection != nil {
		in, out := &in.HotlinkProtection, &out.HotlinkProtection
		*out = new(bool)
		**out = **in
	}
	if in.HTTP2 != nil {
		in, out := &in.HTTP2, &out.HTTP2
		*out = new(bool)
		**out = **in
	}
	if in.HTTP3 != nil {
		in, out := &in.HTTP3, &out.HTTP3
		*out = new(bool)
		**out = **in
	}
	if in.IPGeolocation != nil {
		in, out := &in.IPGeolocation, &out.IPGeolocation
		*out = new(bool)
		**out = **in
	}
	if in.IPv6 != nil {
		in, out := &in.IPv6, &out.IPv6
		*out = new(bool)
		**out = **in
	}
	if in.LogToCloudflare != nil {
		in, out := &in.LogToCloudflare, &out.LogToCloudflare
		*out = new(bool)
		**out = **in
	}
	if in.MaxUpload != nil {
		in, out := &in.MaxUpload, &out.MaxUpload
		*out = new(int)
		**out = **in
	}
	if in.MinTLSVersion != nil {
		in, out := &in.MinTLSVersion, &out.MinTLSVersion
		*out = new(string)
		**out = **in
	}
	if in.Mirage != nil {
		in, out := &in.Mirage, &out.Mirage
		*out = new(bool)
		**out = **in
	}
	if in.OpportunisticEncryption != nil {
		in, out := &in.OpportunisticEncryption, &out.OpportunisticEncryption
		*out = new(bool)
		**out = **in
	}
	if in.OpportunisticOnion != nil {
		in, out := &in.OpportunisticOnion, &out.OpportunisticOnion
		*out = new(bool)
		**out = **in
	}
	if in.OrangeToOrange != nil {
		in, out := &in.OrangeToOrange, &out.OrangeToOrange
		*out = new(bool)
		**out = **in
	}
	if in.OriginErrorPagePassThru != nil {
		in, out := &in.OriginErrorPagePassThru, &out.OriginErrorPagePassThru
		*out = new(bool)
		**out = **in
	}
	if in.Polish != nil {
		in, out := &in.Polish, &out.Polish
		*out = new(string)
		**out = **in
	}
	if in.PrefetchPreload != nil {
		in, out := &in.PrefetchPreload, &out.PrefetchPreload
		*out = new(bool)
		**out = **in
	}
	if in.PrivacyPass != nil {
		in, out := &in.PrivacyPass, &out.PrivacyPass
		*out = new(bool)
		**out = **in
	}
	if in.PseudoIPv4 != nil {
		in, out := &in.PseudoIPv4, &out.PseudoIPv4
		*out = new(string)
		**out = **in
	}
	if in.ResponseBuffering != nil {
		in, out := &in.ResponseBuffering, &out.ResponseBuffering
		*out = new(bool)
		**out = **in
	}
	if in.RocketLoader != nil {
		in, out := &in.RocketLoader, &out.RocketLoader
		*out = new(bool)
		**out = **in
	}
	if in.SecurityLevel != nil {
		in, out := &in.SecurityLevel, &out.SecurityLevel
		*out = new(string)
		**out = **in
	}
	if in.ServerSideExclude != nil {
		in, out := &in.ServerSideExclude, &out.ServerSideExclude
		*out = new(bool)
		**out = **in
	}
	if in.SortQueryStringForCache != nil {
		in, out := &in.SortQueryStringForCache, &out.SortQueryStringForCache
		*out = new(bool)
		**out = **in
	}
	if in.SSL != nil {
		in, out := &in.SSL, &out.SSL
		*out = new(string)
		**out = **in
	}
	if in.TLS13 != nil {
		in, out := &in.TLS13, &out.TLS13
		*out = new(string)
		**out = **in
	}
	if in.TLSClientAuth != nil {
		in, out := &in.TLSClientAuth, &out.TLSClientAuth
		*out = new(bool)
		**out = **in
	}
	if in.TrueClientIPHeader != nil {
		in, out := &in.TrueClientIPHeader, &out.TrueClientIPHeader
		*out = new(bool)
		**out = **in
	}
	if in.VisitorIP != nil {
		in, out := &in.VisitorIP, &out.VisitorIP
		*out = new(bool)
		**out = **in
	}
	if in.WAF != nil {
		in, out := &in.WAF, &out.WAF
		*out = new(bool)
		**out = **in
	}
	if in.WebP != nil {
		in, out := &in.WebP, &out.WebP
		*out = new(bool)
		**out = **in
	}
	if in.WebSockets != nil {
		in, out := &in.WebSockets, &out.WebSockets
		*out = new(bool)
		**out = **in
	}
	if in.ZeroRTT != nil {
		in, out := &in.ZeroRTT, &out.ZeroRTT
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneSettings.
func (in *ZoneSettings) DeepCopy() *ZoneSettings {
	if in == nil {
		return nil
	}
	out := new(ZoneSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneSpec) DeepCopyInto(out *ZoneSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneSpec.
func (in *ZoneSpec) DeepCopy() *ZoneSpec {
	if in == nil {
		return nil
	}
	out := new(ZoneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneStatus) DeepCopyInto(out *ZoneStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneStatus.
func (in *ZoneStatus) DeepCopy() *ZoneStatus {
	if in == nil {
		return nil
	}
	out := new(ZoneStatus)
	in.DeepCopyInto(out)
	return out
}