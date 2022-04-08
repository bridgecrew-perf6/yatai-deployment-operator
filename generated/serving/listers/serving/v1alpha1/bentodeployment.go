/*
Copyright 2022.

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
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	v1alpha1 "github.com/bentoml/yatai-deployment-operator/api/serving/v1alpha1"
)

// BentoDeploymentLister helps list BentoDeployments.
// All objects returned here must be treated as read-only.
type BentoDeploymentLister interface {
	// List lists all BentoDeployments in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BentoDeployment, err error)
	// BentoDeployments returns an object that can list and get BentoDeployments.
	BentoDeployments(namespace string) BentoDeploymentNamespaceLister
	BentoDeploymentListerExpansion
}

// bentoDeploymentLister implements the BentoDeploymentLister interface.
type bentoDeploymentLister struct {
	indexer cache.Indexer
}

// NewBentoDeploymentLister returns a new BentoDeploymentLister.
func NewBentoDeploymentLister(indexer cache.Indexer) BentoDeploymentLister {
	return &bentoDeploymentLister{indexer: indexer}
}

// List lists all BentoDeployments in the indexer.
func (s *bentoDeploymentLister) List(selector labels.Selector) (ret []*v1alpha1.BentoDeployment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BentoDeployment))
	})
	return ret, err
}

// BentoDeployments returns an object that can list and get BentoDeployments.
func (s *bentoDeploymentLister) BentoDeployments(namespace string) BentoDeploymentNamespaceLister {
	return bentoDeploymentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BentoDeploymentNamespaceLister helps list and get BentoDeployments.
// All objects returned here must be treated as read-only.
type BentoDeploymentNamespaceLister interface {
	// List lists all BentoDeployments in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BentoDeployment, err error)
	// Get retrieves the BentoDeployment from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.BentoDeployment, error)
	BentoDeploymentNamespaceListerExpansion
}

// bentoDeploymentNamespaceLister implements the BentoDeploymentNamespaceLister
// interface.
type bentoDeploymentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BentoDeployments in the indexer for a given namespace.
func (s bentoDeploymentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.BentoDeployment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BentoDeployment))
	})
	return ret, err
}

// Get retrieves the BentoDeployment from the indexer for a given namespace and name.
func (s bentoDeploymentNamespaceLister) Get(name string) (*v1alpha1.BentoDeployment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("bentodeployment"), name)
	}
	return obj.(*v1alpha1.BentoDeployment), nil
}
