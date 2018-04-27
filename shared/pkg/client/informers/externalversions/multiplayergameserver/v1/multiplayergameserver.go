/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	multiplayergameserver_v1 "github.com/dgkanatsios/azuregameserversscalingkubernetes/shared/pkg/apis/multiplayergameserver/v1"
	versioned "github.com/dgkanatsios/azuregameserversscalingkubernetes/shared/pkg/client/clientset/versioned"
	internalinterfaces "github.com/dgkanatsios/azuregameserversscalingkubernetes/shared/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/dgkanatsios/azuregameserversscalingkubernetes/shared/pkg/client/listers/multiplayergameserver/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// MultiplayerGameServerInformer provides access to a shared informer and lister for
// MultiplayerGameServers.
type MultiplayerGameServerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.MultiplayerGameServerLister
}

type multiplayerGameServerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewMultiplayerGameServerInformer constructs a new informer for MultiplayerGameServer type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMultiplayerGameServerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMultiplayerGameServerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredMultiplayerGameServerInformer constructs a new informer for MultiplayerGameServer type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMultiplayerGameServerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AzureV1().MultiplayerGameServers(namespace).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AzureV1().MultiplayerGameServers(namespace).Watch(options)
			},
		},
		&multiplayergameserver_v1.MultiplayerGameServer{},
		resyncPeriod,
		indexers,
	)
}

func (f *multiplayerGameServerInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMultiplayerGameServerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *multiplayerGameServerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&multiplayergameserver_v1.MultiplayerGameServer{}, f.defaultInformer)
}

func (f *multiplayerGameServerInformer) Lister() v1.MultiplayerGameServerLister {
	return v1.NewMultiplayerGameServerLister(f.Informer().GetIndexer())
}