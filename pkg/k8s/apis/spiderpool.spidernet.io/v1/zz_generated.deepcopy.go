//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAllocationDetail) DeepCopyInto(out *IPAllocationDetail) {
	*out = *in
	if in.IPv4 != nil {
		in, out := &in.IPv4, &out.IPv4
		*out = new(string)
		**out = **in
	}
	if in.IPv6 != nil {
		in, out := &in.IPv6, &out.IPv6
		*out = new(string)
		**out = **in
	}
	if in.IPv4Pool != nil {
		in, out := &in.IPv4Pool, &out.IPv4Pool
		*out = new(string)
		**out = **in
	}
	if in.IPv6Pool != nil {
		in, out := &in.IPv6Pool, &out.IPv6Pool
		*out = new(string)
		**out = **in
	}
	if in.Vlan != nil {
		in, out := &in.Vlan, &out.Vlan
		*out = new(int64)
		**out = **in
	}
	if in.IPv4Gateway != nil {
		in, out := &in.IPv4Gateway, &out.IPv4Gateway
		*out = new(string)
		**out = **in
	}
	if in.IPv6Gateway != nil {
		in, out := &in.IPv6Gateway, &out.IPv6Gateway
		*out = new(string)
		**out = **in
	}
	if in.CleanGateway != nil {
		in, out := &in.CleanGateway, &out.CleanGateway
		*out = new(bool)
		**out = **in
	}
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]Route, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAllocationDetail.
func (in *IPAllocationDetail) DeepCopy() *IPAllocationDetail {
	if in == nil {
		return nil
	}
	out := new(IPAllocationDetail)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPPoolSpec) DeepCopyInto(out *IPPoolSpec) {
	*out = *in
	if in.IPVersion != nil {
		in, out := &in.IPVersion, &out.IPVersion
		*out = new(int64)
		**out = **in
	}
	if in.IPs != nil {
		in, out := &in.IPs, &out.IPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Disable != nil {
		in, out := &in.Disable, &out.Disable
		*out = new(bool)
		**out = **in
	}
	if in.ExcludeIPs != nil {
		in, out := &in.ExcludeIPs, &out.ExcludeIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Gateway != nil {
		in, out := &in.Gateway, &out.Gateway
		*out = new(string)
		**out = **in
	}
	if in.Vlan != nil {
		in, out := &in.Vlan, &out.Vlan
		*out = new(int64)
		**out = **in
	}
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]Route, len(*in))
		copy(*out, *in)
	}
	if in.PodAffinity != nil {
		in, out := &in.PodAffinity, &out.PodAffinity
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.NamespaceAffinity != nil {
		in, out := &in.NamespaceAffinity, &out.NamespaceAffinity
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeAffinity != nil {
		in, out := &in.NodeAffinity, &out.NodeAffinity
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPPoolSpec.
func (in *IPPoolSpec) DeepCopy() *IPPoolSpec {
	if in == nil {
		return nil
	}
	out := new(IPPoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPPoolStatus) DeepCopyInto(out *IPPoolStatus) {
	*out = *in
	if in.AllocatedIPs != nil {
		in, out := &in.AllocatedIPs, &out.AllocatedIPs
		*out = make(PoolIPAllocations, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TotalIPCount != nil {
		in, out := &in.TotalIPCount, &out.TotalIPCount
		*out = new(int64)
		**out = **in
	}
	if in.AllocatedIPCount != nil {
		in, out := &in.AllocatedIPCount, &out.AllocatedIPCount
		*out = new(int64)
		**out = **in
	}
	if in.AutoDesiredIPCount != nil {
		in, out := &in.AutoDesiredIPCount, &out.AutoDesiredIPCount
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPPoolStatus.
func (in *IPPoolStatus) DeepCopy() *IPPoolStatus {
	if in == nil {
		return nil
	}
	out := new(IPPoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodIPAllocation) DeepCopyInto(out *PodIPAllocation) {
	*out = *in
	if in.Node != nil {
		in, out := &in.Node, &out.Node
		*out = new(string)
		**out = **in
	}
	if in.IPs != nil {
		in, out := &in.IPs, &out.IPs
		*out = make([]IPAllocationDetail, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CreationTime != nil {
		in, out := &in.CreationTime, &out.CreationTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodIPAllocation.
func (in *PodIPAllocation) DeepCopy() *PodIPAllocation {
	if in == nil {
		return nil
	}
	out := new(PodIPAllocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolIPAllocation) DeepCopyInto(out *PoolIPAllocation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolIPAllocation.
func (in *PoolIPAllocation) DeepCopy() *PoolIPAllocation {
	if in == nil {
		return nil
	}
	out := new(PoolIPAllocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in PoolIPAllocations) DeepCopyInto(out *PoolIPAllocations) {
	{
		in := &in
		*out = make(PoolIPAllocations, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolIPAllocations.
func (in PoolIPAllocations) DeepCopy() PoolIPAllocations {
	if in == nil {
		return nil
	}
	out := new(PoolIPAllocations)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolIPPreAllocation) DeepCopyInto(out *PoolIPPreAllocation) {
	*out = *in
	if in.IPs != nil {
		in, out := &in.IPs, &out.IPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolIPPreAllocation.
func (in *PoolIPPreAllocation) DeepCopy() *PoolIPPreAllocation {
	if in == nil {
		return nil
	}
	out := new(PoolIPPreAllocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in PoolIPPreAllocations) DeepCopyInto(out *PoolIPPreAllocations) {
	{
		in := &in
		*out = make(PoolIPPreAllocations, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolIPPreAllocations.
func (in PoolIPPreAllocations) DeepCopy() PoolIPPreAllocations {
	if in == nil {
		return nil
	}
	out := new(PoolIPPreAllocations)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReservedIPSpec) DeepCopyInto(out *ReservedIPSpec) {
	*out = *in
	if in.IPVersion != nil {
		in, out := &in.IPVersion, &out.IPVersion
		*out = new(int64)
		**out = **in
	}
	if in.IPs != nil {
		in, out := &in.IPs, &out.IPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReservedIPSpec.
func (in *ReservedIPSpec) DeepCopy() *ReservedIPSpec {
	if in == nil {
		return nil
	}
	out := new(ReservedIPSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route) DeepCopyInto(out *Route) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route.
func (in *Route) DeepCopy() *Route {
	if in == nil {
		return nil
	}
	out := new(Route)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpiderEndpoint) DeepCopyInto(out *SpiderEndpoint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpiderEndpoint.
func (in *SpiderEndpoint) DeepCopy() *SpiderEndpoint {
	if in == nil {
		return nil
	}
	out := new(SpiderEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpiderEndpoint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpiderEndpointList) DeepCopyInto(out *SpiderEndpointList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SpiderEndpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpiderEndpointList.
func (in *SpiderEndpointList) DeepCopy() *SpiderEndpointList {
	if in == nil {
		return nil
	}
	out := new(SpiderEndpointList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpiderEndpointList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpiderIPPool) DeepCopyInto(out *SpiderIPPool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpiderIPPool.
func (in *SpiderIPPool) DeepCopy() *SpiderIPPool {
	if in == nil {
		return nil
	}
	out := new(SpiderIPPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpiderIPPool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpiderIPPoolList) DeepCopyInto(out *SpiderIPPoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SpiderIPPool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpiderIPPoolList.
func (in *SpiderIPPoolList) DeepCopy() *SpiderIPPoolList {
	if in == nil {
		return nil
	}
	out := new(SpiderIPPoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpiderIPPoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpiderReservedIP) DeepCopyInto(out *SpiderReservedIP) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpiderReservedIP.
func (in *SpiderReservedIP) DeepCopy() *SpiderReservedIP {
	if in == nil {
		return nil
	}
	out := new(SpiderReservedIP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpiderReservedIP) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpiderReservedIPList) DeepCopyInto(out *SpiderReservedIPList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SpiderReservedIP, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpiderReservedIPList.
func (in *SpiderReservedIPList) DeepCopy() *SpiderReservedIPList {
	if in == nil {
		return nil
	}
	out := new(SpiderReservedIPList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpiderReservedIPList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpiderSubnet) DeepCopyInto(out *SpiderSubnet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpiderSubnet.
func (in *SpiderSubnet) DeepCopy() *SpiderSubnet {
	if in == nil {
		return nil
	}
	out := new(SpiderSubnet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpiderSubnet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpiderSubnetList) DeepCopyInto(out *SpiderSubnetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SpiderSubnet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpiderSubnetList.
func (in *SpiderSubnetList) DeepCopy() *SpiderSubnetList {
	if in == nil {
		return nil
	}
	out := new(SpiderSubnetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpiderSubnetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetSpec) DeepCopyInto(out *SubnetSpec) {
	*out = *in
	if in.IPVersion != nil {
		in, out := &in.IPVersion, &out.IPVersion
		*out = new(int64)
		**out = **in
	}
	if in.IPs != nil {
		in, out := &in.IPs, &out.IPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExcludeIPs != nil {
		in, out := &in.ExcludeIPs, &out.ExcludeIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Gateway != nil {
		in, out := &in.Gateway, &out.Gateway
		*out = new(string)
		**out = **in
	}
	if in.Vlan != nil {
		in, out := &in.Vlan, &out.Vlan
		*out = new(int64)
		**out = **in
	}
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]Route, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetSpec.
func (in *SubnetSpec) DeepCopy() *SubnetSpec {
	if in == nil {
		return nil
	}
	out := new(SubnetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetStatus) DeepCopyInto(out *SubnetStatus) {
	*out = *in
	if in.ControlledIPPools != nil {
		in, out := &in.ControlledIPPools, &out.ControlledIPPools
		*out = make(PoolIPPreAllocations, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.TotalIPCount != nil {
		in, out := &in.TotalIPCount, &out.TotalIPCount
		*out = new(int64)
		**out = **in
	}
	if in.AllocatedIPCount != nil {
		in, out := &in.AllocatedIPCount, &out.AllocatedIPCount
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetStatus.
func (in *SubnetStatus) DeepCopy() *SubnetStatus {
	if in == nil {
		return nil
	}
	out := new(SubnetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadEndpointStatus) DeepCopyInto(out *WorkloadEndpointStatus) {
	*out = *in
	if in.Current != nil {
		in, out := &in.Current, &out.Current
		*out = new(PodIPAllocation)
		(*in).DeepCopyInto(*out)
	}
	if in.History != nil {
		in, out := &in.History, &out.History
		*out = make([]PodIPAllocation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadEndpointStatus.
func (in *WorkloadEndpointStatus) DeepCopy() *WorkloadEndpointStatus {
	if in == nil {
		return nil
	}
	out := new(WorkloadEndpointStatus)
	in.DeepCopyInto(out)
	return out
}
