/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package main

import (
	jsoniter "github.com/json-iterator/go"
	"k8s.io/apimachinery/pkg/runtime/schema"
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
	"kubeform.dev/provider-hcloud-controller/controllers"
)

type Data struct {
	JsonIt       jsoniter.API
	ResourceType string
}

var (
	allJsonIt = map[schema.GroupVersionResource]Data{
		{
			Group:    "certificate.hcloud.kubeform.com",
			Version:  "v1alpha1",
			Resource: "certificates",
		}: {
			JsonIt:       controllers.GetJSONItr(certificatev1alpha1.GetEncoder(), certificatev1alpha1.GetDecoder()),
			ResourceType: "hcloud_certificate",
		},
		{
			Group:    "firewall.hcloud.kubeform.com",
			Version:  "v1alpha1",
			Resource: "firewalls",
		}: {
			JsonIt:       controllers.GetJSONItr(firewallv1alpha1.GetEncoder(), firewallv1alpha1.GetDecoder()),
			ResourceType: "hcloud_firewall",
		},
		{
			Group:    "floating.hcloud.kubeform.com",
			Version:  "v1alpha1",
			Resource: "ips",
		}: {
			JsonIt:       controllers.GetJSONItr(floatingv1alpha1.GetEncoder(), floatingv1alpha1.GetDecoder()),
			ResourceType: "hcloud_floating_ip",
		},
		{
			Group:    "floating.hcloud.kubeform.com",
			Version:  "v1alpha1",
			Resource: "ipassignments",
		}: {
			JsonIt:       controllers.GetJSONItr(floatingv1alpha1.GetEncoder(), floatingv1alpha1.GetDecoder()),
			ResourceType: "hcloud_floating_ip_assignment",
		},
		{
			Group:    "load.hcloud.kubeform.com",
			Version:  "v1alpha1",
			Resource: "balancers",
		}: {
			JsonIt:       controllers.GetJSONItr(loadv1alpha1.GetEncoder(), loadv1alpha1.GetDecoder()),
			ResourceType: "hcloud_load_balancer",
		},
		{
			Group:    "load.hcloud.kubeform.com",
			Version:  "v1alpha1",
			Resource: "balancernetworks",
		}: {
			JsonIt:       controllers.GetJSONItr(loadv1alpha1.GetEncoder(), loadv1alpha1.GetDecoder()),
			ResourceType: "hcloud_load_balancer_network",
		},
		{
			Group:    "load.hcloud.kubeform.com",
			Version:  "v1alpha1",
			Resource: "balancerservices",
		}: {
			JsonIt:       controllers.GetJSONItr(loadv1alpha1.GetEncoder(), loadv1alpha1.GetDecoder()),
			ResourceType: "hcloud_load_balancer_service",
		},
		{
			Group:    "load.hcloud.kubeform.com",
			Version:  "v1alpha1",
			Resource: "balancertargets",
		}: {
			JsonIt:       controllers.GetJSONItr(loadv1alpha1.GetEncoder(), loadv1alpha1.GetDecoder()),
			ResourceType: "hcloud_load_balancer_target",
		},
		{
			Group:    "managed.hcloud.kubeform.com",
			Version:  "v1alpha1",
			Resource: "certificates",
		}: {
			JsonIt:       controllers.GetJSONItr(managedv1alpha1.GetEncoder(), managedv1alpha1.GetDecoder()),
			ResourceType: "hcloud_managed_certificate",
		},
		{
			Group:    "network.hcloud.kubeform.com",
			Version:  "v1alpha1",
			Resource: "networks",
		}: {
			JsonIt:       controllers.GetJSONItr(networkv1alpha1.GetEncoder(), networkv1alpha1.GetDecoder()),
			ResourceType: "hcloud_network",
		},
		{
			Group:    "network.hcloud.kubeform.com",
			Version:  "v1alpha1",
			Resource: "routes",
		}: {
			JsonIt:       controllers.GetJSONItr(networkv1alpha1.GetEncoder(), networkv1alpha1.GetDecoder()),
			ResourceType: "hcloud_network_route",
		},
		{
			Group:    "network.hcloud.kubeform.com",
			Version:  "v1alpha1",
			Resource: "subnets",
		}: {
			JsonIt:       controllers.GetJSONItr(networkv1alpha1.GetEncoder(), networkv1alpha1.GetDecoder()),
			ResourceType: "hcloud_network_subnet",
		},
		{
			Group:    "rdns.hcloud.kubeform.com",
			Version:  "v1alpha1",
			Resource: "rdns",
		}: {
			JsonIt:       controllers.GetJSONItr(rdnsv1alpha1.GetEncoder(), rdnsv1alpha1.GetDecoder()),
			ResourceType: "hcloud_rdns",
		},
		{
			Group:    "server.hcloud.kubeform.com",
			Version:  "v1alpha1",
			Resource: "servers",
		}: {
			JsonIt:       controllers.GetJSONItr(serverv1alpha1.GetEncoder(), serverv1alpha1.GetDecoder()),
			ResourceType: "hcloud_server",
		},
		{
			Group:    "server.hcloud.kubeform.com",
			Version:  "v1alpha1",
			Resource: "networks",
		}: {
			JsonIt:       controllers.GetJSONItr(serverv1alpha1.GetEncoder(), serverv1alpha1.GetDecoder()),
			ResourceType: "hcloud_server_network",
		},
		{
			Group:    "snapshot.hcloud.kubeform.com",
			Version:  "v1alpha1",
			Resource: "snapshots",
		}: {
			JsonIt:       controllers.GetJSONItr(snapshotv1alpha1.GetEncoder(), snapshotv1alpha1.GetDecoder()),
			ResourceType: "hcloud_snapshot",
		},
		{
			Group:    "ssh.hcloud.kubeform.com",
			Version:  "v1alpha1",
			Resource: "keys",
		}: {
			JsonIt:       controllers.GetJSONItr(sshv1alpha1.GetEncoder(), sshv1alpha1.GetDecoder()),
			ResourceType: "hcloud_ssh_key",
		},
		{
			Group:    "uploaded.hcloud.kubeform.com",
			Version:  "v1alpha1",
			Resource: "certificates",
		}: {
			JsonIt:       controllers.GetJSONItr(uploadedv1alpha1.GetEncoder(), uploadedv1alpha1.GetDecoder()),
			ResourceType: "hcloud_uploaded_certificate",
		},
		{
			Group:    "volume.hcloud.kubeform.com",
			Version:  "v1alpha1",
			Resource: "volumes",
		}: {
			JsonIt:       controllers.GetJSONItr(volumev1alpha1.GetEncoder(), volumev1alpha1.GetDecoder()),
			ResourceType: "hcloud_volume",
		},
		{
			Group:    "volume.hcloud.kubeform.com",
			Version:  "v1alpha1",
			Resource: "attachments",
		}: {
			JsonIt:       controllers.GetJSONItr(volumev1alpha1.GetEncoder(), volumev1alpha1.GetDecoder()),
			ResourceType: "hcloud_volume_attachment",
		},
	}
)

func getJsonItAndResType(gvr schema.GroupVersionResource) Data {
	return allJsonIt[gvr]
}
