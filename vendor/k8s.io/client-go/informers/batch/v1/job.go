/*
<<<<<<< HEAD
Copyright 2018 The Kubernetes Authors.
=======
Copyright 2017 The Kubernetes Authors.
>>>>>>> Initial dep workover

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

// This file was automatically generated by informer-gen

package v1

import (
	batch_v1 "k8s.io/api/batch/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	kubernetes "k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/listers/batch/v1"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// JobInformer provides access to a shared informer and lister for
// Jobs.
type JobInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.JobLister
}

type jobInformer struct {
<<<<<<< HEAD
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
=======
	factory internalinterfaces.SharedInformerFactory
>>>>>>> Initial dep workover
}

// NewJobInformer constructs a new informer for Job type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewJobInformer(client kubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
<<<<<<< HEAD
	return NewFilteredJobInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredJobInformer constructs a new informer for Job type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredJobInformer(client kubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BatchV1().Jobs(namespace).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
=======
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				return client.BatchV1().Jobs(namespace).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
>>>>>>> Initial dep workover
				return client.BatchV1().Jobs(namespace).Watch(options)
			},
		},
		&batch_v1.Job{},
		resyncPeriod,
		indexers,
	)
}

<<<<<<< HEAD
func (f *jobInformer) defaultInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredJobInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *jobInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&batch_v1.Job{}, f.defaultInformer)
=======
func defaultJobInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewJobInformer(client, meta_v1.NamespaceAll, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *jobInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&batch_v1.Job{}, defaultJobInformer)
>>>>>>> Initial dep workover
}

func (f *jobInformer) Lister() v1.JobLister {
	return v1.NewJobLister(f.Informer().GetIndexer())
}