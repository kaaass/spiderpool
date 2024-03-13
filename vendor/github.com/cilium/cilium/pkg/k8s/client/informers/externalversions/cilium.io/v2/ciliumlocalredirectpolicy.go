// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by informer-gen. DO NOT EDIT.

package v2

import (
	"context"
	time "time"

	ciliumiov2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	versioned "github.com/cilium/cilium/pkg/k8s/client/clientset/versioned"
	internalinterfaces "github.com/cilium/cilium/pkg/k8s/client/informers/externalversions/internalinterfaces"
	v2 "github.com/cilium/cilium/pkg/k8s/client/listers/cilium.io/v2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CiliumLocalRedirectPolicyInformer provides access to a shared informer and lister for
// CiliumLocalRedirectPolicies.
type CiliumLocalRedirectPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v2.CiliumLocalRedirectPolicyLister
}

type ciliumLocalRedirectPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCiliumLocalRedirectPolicyInformer constructs a new informer for CiliumLocalRedirectPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCiliumLocalRedirectPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCiliumLocalRedirectPolicyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCiliumLocalRedirectPolicyInformer constructs a new informer for CiliumLocalRedirectPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCiliumLocalRedirectPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CiliumV2().CiliumLocalRedirectPolicies(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CiliumV2().CiliumLocalRedirectPolicies(namespace).Watch(context.TODO(), options)
			},
		},
		&ciliumiov2.CiliumLocalRedirectPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *ciliumLocalRedirectPolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCiliumLocalRedirectPolicyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *ciliumLocalRedirectPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&ciliumiov2.CiliumLocalRedirectPolicy{}, f.defaultInformer)
}

func (f *ciliumLocalRedirectPolicyInformer) Lister() v2.CiliumLocalRedirectPolicyLister {
	return v2.NewCiliumLocalRedirectPolicyLister(f.Informer().GetIndexer())
}