/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by client-gen. DO NOT EDIT.

package scheme

import (
	certificatev1alpha1 "kubeform.dev/provider-hcloud-api/apis/certificate/v1alpha1"
	firewallv1alpha1 "kubeform.dev/provider-hcloud-api/apis/firewall/v1alpha1"
	floatingv1alpha1 "kubeform.dev/provider-hcloud-api/apis/floating/v1alpha1"
	loadv1alpha1 "kubeform.dev/provider-hcloud-api/apis/load/v1alpha1"
	managedv1alpha1 "kubeform.dev/provider-hcloud-api/apis/managed/v1alpha1"
	networkv1alpha1 "kubeform.dev/provider-hcloud-api/apis/network/v1alpha1"
	rdnsv1alpha1 "kubeform.dev/provider-hcloud-api/apis/rdns/v1alpha1"
	serverv1alpha1 "kubeform.dev/provider-hcloud-api/apis/server/v1alpha1"
	snapshotv1alpha1 "kubeform.dev/provider-hcloud-api/apis/snapshot/v1alpha1"
	sshv1alpha1 "kubeform.dev/provider-hcloud-api/apis/ssh/v1alpha1"
	uploadedv1alpha1 "kubeform.dev/provider-hcloud-api/apis/uploaded/v1alpha1"
	volumev1alpha1 "kubeform.dev/provider-hcloud-api/apis/volume/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

var Scheme = runtime.NewScheme()
var Codecs = serializer.NewCodecFactory(Scheme)
var ParameterCodec = runtime.NewParameterCodec(Scheme)
var localSchemeBuilder = runtime.SchemeBuilder{
	certificatev1alpha1.AddToScheme,
	firewallv1alpha1.AddToScheme,
	floatingv1alpha1.AddToScheme,
	loadv1alpha1.AddToScheme,
	managedv1alpha1.AddToScheme,
	networkv1alpha1.AddToScheme,
	rdnsv1alpha1.AddToScheme,
	serverv1alpha1.AddToScheme,
	snapshotv1alpha1.AddToScheme,
	sshv1alpha1.AddToScheme,
	uploadedv1alpha1.AddToScheme,
	volumev1alpha1.AddToScheme,
}

// AddToScheme adds all types of this clientset into the given scheme. This allows composition
// of clientsets, like in:
//
//   import (
//     "k8s.io/client-go/kubernetes"
//     clientsetscheme "k8s.io/client-go/kubernetes/scheme"
//     aggregatorclientsetscheme "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset/scheme"
//   )
//
//   kclientset, _ := kubernetes.NewForConfig(c)
//   _ = aggregatorclientsetscheme.AddToScheme(clientsetscheme.Scheme)
//
// After this, RawExtensions in Kubernetes types will serialize kube-aggregator types
// correctly.
var AddToScheme = localSchemeBuilder.AddToScheme

func init() {
	v1.AddToGroupVersion(Scheme, schema.GroupVersion{Version: "v1"})
	utilruntime.Must(AddToScheme(Scheme))
}