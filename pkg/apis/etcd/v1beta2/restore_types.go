// Copyright 2017 The etcd-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1beta2

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EtcdRestoreList is a list of EtcdRestore.
type EtcdRestoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []EtcdRestore `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status

// EtcdRestore represents a Kubernetes EtcdRestore Custom Resource.
// The EtcdRestore CR name will be used as the name of the new restored cluster.
type EtcdRestore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              RestoreSpec   `json:"spec"`
	Status            RestoreStatus `json:"status,omitempty"`
}

// RestoreSpec defines how to restore an etcd cluster from existing backup.
type RestoreSpec struct {
	// BackupStorageType is the type of the backup storage which is used as RestoreSource.
	BackupStorageType BackupStorageType `json:"backupStorageType"`
	// RestoreSource tells the where to get the backup and restore from.
	RestoreSource `json:",inline"`
	// EtcdCluster references an EtcdCluster resource whose metadata and spec
	// will be used to create the new restored EtcdCluster CR.
	// This reference EtcdCluster CR and all its resources will be deleted before the
	// restored EtcdCluster CR is created.
	EtcdCluster EtcdClusterRef `json:"etcdCluster"`
}

// EtcdCluster references an EtcdCluster resource whose metadata and spec
// will be used to create the new restored EtcdCluster CR.
// This reference EtcdCluster CR and all its resources will be deleted before the
// restored EtcdCluster CR is created.
type EtcdClusterRef struct {
	// Name is the EtcdCluster resource name.
	// This reference EtcdCluster must be present in the same namespace as the restore-operator
	Name string `json:"name"`
}

type RestoreSource struct {
	// S3 tells where on S3 the backup is saved and how to fetch the backup.
	S3 *S3RestoreSource `json:"s3,omitempty"`

	// ABS tells where on ABS the backup is saved and how to fetch the backup.
	ABS *ABSRestoreSource `json:"abs,omitempty"`

	// GCS tells where on GCS the backup is saved and how to fetch the backup.
	GCS *GCSRestoreSource `json:"gcs,omitempty"`

	// OSS tells where on OSS the backup is saved and how to fetch the backup.
	OSS *OSSRestoreSource `json:"oss,omitempty"`
}

type S3RestoreSource struct {
	// Path is the full s3 path where the backup is saved.
	// The format of the path must be: "<s3-bucket-name>/<path-to-backup-file>"
	// e.g: "mybucket/etcd.backup"
	Path string `json:"path"`

	// The name of the secret object that stores the AWS credential and config files.
	// The file name of the credential MUST be 'credentials'.
	// The file name of the config MUST be 'config'.
	// The profile to use in both files will be 'default'.
	//
	// AWSSecret overwrites the default etcd operator wide AWS credential and config.
	AWSSecret string `json:"awsSecret"`

	// Endpoint if blank points to aws. If specified, can point to s3 compatible object
	// stores.
	Endpoint string `json:"endpoint"`

	// ForcePathStyle forces to use path style over the default subdomain style.
	// This is useful when you have an s3 compatible endpoint that doesn't support
	// subdomain buckets.
	ForcePathStyle bool `json:"forcePathStyle"`
}

type ABSRestoreSource struct {
	// Path is the full abs path where the backup is saved.
	// The format of the path must be: "<abs-container-name>/<path-to-backup-file>"
	// e.g: "myabscontainer/etcd.backup"
	Path string `json:"path"`

	// The name of the secret object that stores the Azure Blob Storage credential.
	ABSSecret string `json:"absSecret"`
}

type GCSRestoreSource struct {
	// Path is the full GCS path where the backup is saved.
	// The format of the path must be: "<gcs-bucket-name>/<path-to-backup-file>"
	// e.g: "mygcsbucket/etcd.backup"
	Path string `json:"path"`

	// The name of the secret object that stores the Google storage credential
	// containing at most ONE of the following:
	// An access token with file name of 'access-token'.
	// JSON credentials with file name of 'credentials.json'.
	//
	// If omitted, client will use the default application credentials.
	GCPSecret string `json:"gcpSecret,omitempty"`
}

type OSSRestoreSource struct {
	// Path is the full abs path where the backup is saved.
	// The format of the path must be: "<oss-bucket-name>/<path-to-backup-file>"
	// e.g: "myossbucket/etcd.backup"
	Path string `json:"path"`

	// The name of the secret object that stores the credential which will be used
	// to access Alibaba Cloud OSS.
	//
	// The secret must contain the following keys/fields:
	//     accessKeyID
	//     accessKeySecret
	//
	// The format of secret:
	//
	//   apiVersion: v1
	//   kind: Secret
	//   metadata:
	//     name: <my-credential-name>
	//   type: Opaque
	//   data:
	//     accessKeyID: <base64 of my-access-key-id>
	//     accessKeySecret: <base64 of my-access-key-secret>
	//
	OSSSecret string `json:"ossSecret"`

	// Endpoint is the OSS service endpoint on alibaba cloud, defaults to
	// "http://oss-cn-hangzhou.aliyuncs.com".
	//
	// Details about regions and endpoints, see:
	//  https://www.alibabacloud.com/help/doc-detail/31837.htm
	Endpoint string `json:"endpoint,omitempty"`
}

// RestoreStatus reports the status of this restore operation.
type RestoreStatus struct {
	// Succeeded indicates if the backup has Succeeded.
	Succeeded bool `json:"succeeded"`
	// Reason indicates the reason for any backup related failures.
	Reason string `json:"reason,omitempty"`
}
