//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by deepcopy-gen. DO NOT EDIT.

package api

import (
	labels "github.com/cilium/cilium/pkg/k8s/slim/k8s/apis/labels"
	v1 "github.com/cilium/cilium/pkg/k8s/slim/k8s/apis/meta/v1"
	kafka "github.com/cilium/proxy/pkg/policy/api/kafka"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSGroup) DeepCopyInto(out *AWSGroup) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SecurityGroupsIds != nil {
		in, out := &in.SecurityGroupsIds, &out.SecurityGroupsIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroupsNames != nil {
		in, out := &in.SecurityGroupsNames, &out.SecurityGroupsNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSGroup.
func (in *AWSGroup) DeepCopy() *AWSGroup {
	if in == nil {
		return nil
	}
	out := new(AWSGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Authentication) DeepCopyInto(out *Authentication) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Authentication.
func (in *Authentication) DeepCopy() *Authentication {
	if in == nil {
		return nil
	}
	out := new(Authentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CIDRRule) DeepCopyInto(out *CIDRRule) {
	*out = *in
	if in.ExceptCIDRs != nil {
		in, out := &in.ExceptCIDRs, &out.ExceptCIDRs
		*out = make([]CIDR, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CIDRRule.
func (in *CIDRRule) DeepCopy() *CIDRRule {
	if in == nil {
		return nil
	}
	out := new(CIDRRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in CIDRRuleSlice) DeepCopyInto(out *CIDRRuleSlice) {
	{
		in := &in
		*out = make(CIDRRuleSlice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CIDRRuleSlice.
func (in CIDRRuleSlice) DeepCopy() CIDRRuleSlice {
	if in == nil {
		return nil
	}
	out := new(CIDRRuleSlice)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in CIDRSlice) DeepCopyInto(out *CIDRSlice) {
	{
		in := &in
		*out = make(CIDRSlice, len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CIDRSlice.
func (in CIDRSlice) DeepCopy() CIDRSlice {
	if in == nil {
		return nil
	}
	out := new(CIDRSlice)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EgressCommonRule) DeepCopyInto(out *EgressCommonRule) {
	*out = *in
	if in.ToEndpoints != nil {
		in, out := &in.ToEndpoints, &out.ToEndpoints
		*out = make([]EndpointSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ToRequires != nil {
		in, out := &in.ToRequires, &out.ToRequires
		*out = make([]EndpointSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ToCIDR != nil {
		in, out := &in.ToCIDR, &out.ToCIDR
		*out = make(CIDRSlice, len(*in))
		copy(*out, *in)
	}
	if in.ToCIDRSet != nil {
		in, out := &in.ToCIDRSet, &out.ToCIDRSet
		*out = make(CIDRRuleSlice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ToEntities != nil {
		in, out := &in.ToEntities, &out.ToEntities
		*out = make(EntitySlice, len(*in))
		copy(*out, *in)
	}
	if in.ToServices != nil {
		in, out := &in.ToServices, &out.ToServices
		*out = make([]Service, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ToGroups != nil {
		in, out := &in.ToGroups, &out.ToGroups
		*out = make([]ToGroups, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.aggregatedSelectors != nil {
		in, out := &in.aggregatedSelectors, &out.aggregatedSelectors
		*out = make(EndpointSelectorSlice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EgressCommonRule.
func (in *EgressCommonRule) DeepCopy() *EgressCommonRule {
	if in == nil {
		return nil
	}
	out := new(EgressCommonRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EgressDenyRule) DeepCopyInto(out *EgressDenyRule) {
	*out = *in
	in.EgressCommonRule.DeepCopyInto(&out.EgressCommonRule)
	if in.ToPorts != nil {
		in, out := &in.ToPorts, &out.ToPorts
		*out = make(PortDenyRules, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ICMPs != nil {
		in, out := &in.ICMPs, &out.ICMPs
		*out = make(ICMPRules, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EgressDenyRule.
func (in *EgressDenyRule) DeepCopy() *EgressDenyRule {
	if in == nil {
		return nil
	}
	out := new(EgressDenyRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EgressRule) DeepCopyInto(out *EgressRule) {
	*out = *in
	in.EgressCommonRule.DeepCopyInto(&out.EgressCommonRule)
	if in.ToPorts != nil {
		in, out := &in.ToPorts, &out.ToPorts
		*out = make(PortRules, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ToFQDNs != nil {
		in, out := &in.ToFQDNs, &out.ToFQDNs
		*out = make(FQDNSelectorSlice, len(*in))
		copy(*out, *in)
	}
	if in.ICMPs != nil {
		in, out := &in.ICMPs, &out.ICMPs
		*out = make(ICMPRules, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Authentication != nil {
		in, out := &in.Authentication, &out.Authentication
		*out = new(Authentication)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EgressRule.
func (in *EgressRule) DeepCopy() *EgressRule {
	if in == nil {
		return nil
	}
	out := new(EgressRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointSelector) DeepCopyInto(out *EndpointSelector) {
	*out = *in
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.requirements != nil {
		in, out := &in.requirements, &out.requirements
		*out = new(labels.Requirements)
		if **in != nil {
			in, out := *in, *out
			*out = make([]labels.Requirement, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointSelector.
func (in *EndpointSelector) DeepCopy() *EndpointSelector {
	if in == nil {
		return nil
	}
	out := new(EndpointSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in EndpointSelectorSlice) DeepCopyInto(out *EndpointSelectorSlice) {
	{
		in := &in
		*out = make(EndpointSelectorSlice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointSelectorSlice.
func (in EndpointSelectorSlice) DeepCopy() EndpointSelectorSlice {
	if in == nil {
		return nil
	}
	out := new(EndpointSelectorSlice)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in EntitySlice) DeepCopyInto(out *EntitySlice) {
	{
		in := &in
		*out = make(EntitySlice, len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EntitySlice.
func (in EntitySlice) DeepCopy() EntitySlice {
	if in == nil {
		return nil
	}
	out := new(EntitySlice)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvoyConfig) DeepCopyInto(out *EnvoyConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyConfig.
func (in *EnvoyConfig) DeepCopy() *EnvoyConfig {
	if in == nil {
		return nil
	}
	out := new(EnvoyConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FQDNSelector) DeepCopyInto(out *FQDNSelector) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FQDNSelector.
func (in *FQDNSelector) DeepCopy() *FQDNSelector {
	if in == nil {
		return nil
	}
	out := new(FQDNSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in FQDNSelectorSlice) DeepCopyInto(out *FQDNSelectorSlice) {
	{
		in := &in
		*out = make(FQDNSelectorSlice, len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FQDNSelectorSlice.
func (in FQDNSelectorSlice) DeepCopy() FQDNSelectorSlice {
	if in == nil {
		return nil
	}
	out := new(FQDNSelectorSlice)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeaderMatch) DeepCopyInto(out *HeaderMatch) {
	*out = *in
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(Secret)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeaderMatch.
func (in *HeaderMatch) DeepCopy() *HeaderMatch {
	if in == nil {
		return nil
	}
	out := new(HeaderMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ICMPField) DeepCopyInto(out *ICMPField) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ICMPField.
func (in *ICMPField) DeepCopy() *ICMPField {
	if in == nil {
		return nil
	}
	out := new(ICMPField)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ICMPRule) DeepCopyInto(out *ICMPRule) {
	*out = *in
	if in.Fields != nil {
		in, out := &in.Fields, &out.Fields
		*out = make([]ICMPField, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ICMPRule.
func (in *ICMPRule) DeepCopy() *ICMPRule {
	if in == nil {
		return nil
	}
	out := new(ICMPRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ICMPRules) DeepCopyInto(out *ICMPRules) {
	{
		in := &in
		*out = make(ICMPRules, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ICMPRules.
func (in ICMPRules) DeepCopy() ICMPRules {
	if in == nil {
		return nil
	}
	out := new(ICMPRules)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressCommonRule) DeepCopyInto(out *IngressCommonRule) {
	*out = *in
	if in.FromEndpoints != nil {
		in, out := &in.FromEndpoints, &out.FromEndpoints
		*out = make([]EndpointSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FromRequires != nil {
		in, out := &in.FromRequires, &out.FromRequires
		*out = make([]EndpointSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FromCIDR != nil {
		in, out := &in.FromCIDR, &out.FromCIDR
		*out = make(CIDRSlice, len(*in))
		copy(*out, *in)
	}
	if in.FromCIDRSet != nil {
		in, out := &in.FromCIDRSet, &out.FromCIDRSet
		*out = make(CIDRRuleSlice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FromEntities != nil {
		in, out := &in.FromEntities, &out.FromEntities
		*out = make(EntitySlice, len(*in))
		copy(*out, *in)
	}
	if in.aggregatedSelectors != nil {
		in, out := &in.aggregatedSelectors, &out.aggregatedSelectors
		*out = make(EndpointSelectorSlice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressCommonRule.
func (in *IngressCommonRule) DeepCopy() *IngressCommonRule {
	if in == nil {
		return nil
	}
	out := new(IngressCommonRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressDenyRule) DeepCopyInto(out *IngressDenyRule) {
	*out = *in
	in.IngressCommonRule.DeepCopyInto(&out.IngressCommonRule)
	if in.ToPorts != nil {
		in, out := &in.ToPorts, &out.ToPorts
		*out = make(PortDenyRules, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ICMPs != nil {
		in, out := &in.ICMPs, &out.ICMPs
		*out = make(ICMPRules, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressDenyRule.
func (in *IngressDenyRule) DeepCopy() *IngressDenyRule {
	if in == nil {
		return nil
	}
	out := new(IngressDenyRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressRule) DeepCopyInto(out *IngressRule) {
	*out = *in
	in.IngressCommonRule.DeepCopyInto(&out.IngressCommonRule)
	if in.ToPorts != nil {
		in, out := &in.ToPorts, &out.ToPorts
		*out = make(PortRules, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ICMPs != nil {
		in, out := &in.ICMPs, &out.ICMPs
		*out = make(ICMPRules, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Authentication != nil {
		in, out := &in.Authentication, &out.Authentication
		*out = new(Authentication)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressRule.
func (in *IngressRule) DeepCopy() *IngressRule {
	if in == nil {
		return nil
	}
	out := new(IngressRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8sServiceNamespace) DeepCopyInto(out *K8sServiceNamespace) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8sServiceNamespace.
func (in *K8sServiceNamespace) DeepCopy() *K8sServiceNamespace {
	if in == nil {
		return nil
	}
	out := new(K8sServiceNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8sServiceSelectorNamespace) DeepCopyInto(out *K8sServiceSelectorNamespace) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8sServiceSelectorNamespace.
func (in *K8sServiceSelectorNamespace) DeepCopy() *K8sServiceSelectorNamespace {
	if in == nil {
		return nil
	}
	out := new(K8sServiceSelectorNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L7Rules) DeepCopyInto(out *L7Rules) {
	*out = *in
	if in.HTTP != nil {
		in, out := &in.HTTP, &out.HTTP
		*out = make([]PortRuleHTTP, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Kafka != nil {
		in, out := &in.Kafka, &out.Kafka
		*out = make([]kafka.PortRule, len(*in))
		copy(*out, *in)
	}
	if in.DNS != nil {
		in, out := &in.DNS, &out.DNS
		*out = make([]PortRuleDNS, len(*in))
		copy(*out, *in)
	}
	if in.L7 != nil {
		in, out := &in.L7, &out.L7
		*out = make([]PortRuleL7, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make(PortRuleL7, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L7Rules.
func (in *L7Rules) DeepCopy() *L7Rules {
	if in == nil {
		return nil
	}
	out := new(L7Rules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Listener) DeepCopyInto(out *Listener) {
	*out = *in
	if in.EnvoyConfig != nil {
		in, out := &in.EnvoyConfig, &out.EnvoyConfig
		*out = new(EnvoyConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Listener.
func (in *Listener) DeepCopy() *Listener {
	if in == nil {
		return nil
	}
	out := new(Listener)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortDenyRule) DeepCopyInto(out *PortDenyRule) {
	*out = *in
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]PortProtocol, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortDenyRule.
func (in *PortDenyRule) DeepCopy() *PortDenyRule {
	if in == nil {
		return nil
	}
	out := new(PortDenyRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in PortDenyRules) DeepCopyInto(out *PortDenyRules) {
	{
		in := &in
		*out = make(PortDenyRules, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortDenyRules.
func (in PortDenyRules) DeepCopy() PortDenyRules {
	if in == nil {
		return nil
	}
	out := new(PortDenyRules)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortProtocol) DeepCopyInto(out *PortProtocol) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortProtocol.
func (in *PortProtocol) DeepCopy() *PortProtocol {
	if in == nil {
		return nil
	}
	out := new(PortProtocol)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortRule) DeepCopyInto(out *PortRule) {
	*out = *in
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]PortProtocol, len(*in))
		copy(*out, *in)
	}
	if in.TerminatingTLS != nil {
		in, out := &in.TerminatingTLS, &out.TerminatingTLS
		*out = new(TLSContext)
		(*in).DeepCopyInto(*out)
	}
	if in.OriginatingTLS != nil {
		in, out := &in.OriginatingTLS, &out.OriginatingTLS
		*out = new(TLSContext)
		(*in).DeepCopyInto(*out)
	}
	if in.ServerNames != nil {
		in, out := &in.ServerNames, &out.ServerNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Listener != nil {
		in, out := &in.Listener, &out.Listener
		*out = new(Listener)
		(*in).DeepCopyInto(*out)
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = new(L7Rules)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortRule.
func (in *PortRule) DeepCopy() *PortRule {
	if in == nil {
		return nil
	}
	out := new(PortRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortRuleDNS) DeepCopyInto(out *PortRuleDNS) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortRuleDNS.
func (in *PortRuleDNS) DeepCopy() *PortRuleDNS {
	if in == nil {
		return nil
	}
	out := new(PortRuleDNS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortRuleHTTP) DeepCopyInto(out *PortRuleHTTP) {
	*out = *in
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.HeaderMatches != nil {
		in, out := &in.HeaderMatches, &out.HeaderMatches
		*out = make([]*HeaderMatch, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(HeaderMatch)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortRuleHTTP.
func (in *PortRuleHTTP) DeepCopy() *PortRuleHTTP {
	if in == nil {
		return nil
	}
	out := new(PortRuleHTTP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in PortRuleL7) DeepCopyInto(out *PortRuleL7) {
	{
		in := &in
		*out = make(PortRuleL7, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortRuleL7.
func (in PortRuleL7) DeepCopy() PortRuleL7 {
	if in == nil {
		return nil
	}
	out := new(PortRuleL7)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in PortRules) DeepCopyInto(out *PortRules) {
	{
		in := &in
		*out = make(PortRules, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortRules.
func (in PortRules) DeepCopy() PortRules {
	if in == nil {
		return nil
	}
	out := new(PortRules)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rule) DeepCopyInto(out *Rule) {
	*out = *in
	in.EndpointSelector.DeepCopyInto(&out.EndpointSelector)
	in.NodeSelector.DeepCopyInto(&out.NodeSelector)
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = make([]IngressRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IngressDeny != nil {
		in, out := &in.IngressDeny, &out.IngressDeny
		*out = make([]IngressDenyRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Egress != nil {
		in, out := &in.Egress, &out.Egress
		*out = make([]EgressRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EgressDeny != nil {
		in, out := &in.EgressDeny, &out.EgressDeny
		*out = make([]EgressDenyRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Labels = in.Labels.DeepCopy()
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rule.
func (in *Rule) DeepCopy() *Rule {
	if in == nil {
		return nil
	}
	out := new(Rule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Rules) DeepCopyInto(out *Rules) {
	{
		in := &in
		*out = make(Rules, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Rule)
				(*in).DeepCopyInto(*out)
			}
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rules.
func (in Rules) DeepCopy() Rules {
	if in == nil {
		return nil
	}
	out := new(Rules)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Secret) DeepCopyInto(out *Secret) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Secret.
func (in *Secret) DeepCopy() *Secret {
	if in == nil {
		return nil
	}
	out := new(Secret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Service) DeepCopyInto(out *Service) {
	*out = *in
	if in.K8sServiceSelector != nil {
		in, out := &in.K8sServiceSelector, &out.K8sServiceSelector
		*out = new(K8sServiceSelectorNamespace)
		(*in).DeepCopyInto(*out)
	}
	if in.K8sService != nil {
		in, out := &in.K8sService, &out.K8sService
		*out = new(K8sServiceNamespace)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Service.
func (in *Service) DeepCopy() *Service {
	if in == nil {
		return nil
	}
	out := new(Service)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSelector) DeepCopyInto(out *ServiceSelector) {
	*out = *in
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.requirements != nil {
		in, out := &in.requirements, &out.requirements
		*out = new(labels.Requirements)
		if **in != nil {
			in, out := *in, *out
			*out = make([]labels.Requirement, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSelector.
func (in *ServiceSelector) DeepCopy() *ServiceSelector {
	if in == nil {
		return nil
	}
	out := new(ServiceSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSContext) DeepCopyInto(out *TLSContext) {
	*out = *in
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(Secret)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSContext.
func (in *TLSContext) DeepCopy() *TLSContext {
	if in == nil {
		return nil
	}
	out := new(TLSContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ToGroups) DeepCopyInto(out *ToGroups) {
	*out = *in
	if in.AWS != nil {
		in, out := &in.AWS, &out.AWS
		*out = new(AWSGroup)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ToGroups.
func (in *ToGroups) DeepCopy() *ToGroups {
	if in == nil {
		return nil
	}
	out := new(ToGroups)
	in.DeepCopyInto(out)
	return out
}