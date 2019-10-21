package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.
// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file

// ArgoCDSpec defines the desired state of ArgoCD
// +k8s:openapi-gen=true
type ArgoCDSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	ApplicationController ApplicationControllerStruct `json:"applicationController"`
	Server                ServerStruct                `json:"server"`
	RepoServer            RepoServerStruct            `json:"repoServer"`
	DexServer             DexServerStruct             `json:"dexServer"`
	Ingress               IngressStruct               `json:"ingress"`
}

type IngressStruct struct {
	// Enable ingress
	Enabled bool `json:"enabled"`
	// Annotations for ingress object
	Annotations map[string]string `json:"anntotations,omitempty"`
	Path        string            `json:"path"`
	// Ingress additional hosts
	AdditionalHosts []string `json:"additionalHosts,omitempty"`
	// Ingress TLS configuration (check!)
	Tls string `json:"tls,omitempty"`
}
type DexServerStruct struct {
	// Container port for Dex Server HTTP
	ContainerPortHttp int `json:"containerPortHttp"`
	// Container port for Dex Server GRPC
	ContainerPortGrpc int `json:"containerPortGrpc"`
	// Service port for Dex Server HTTP
	ServicePortHttp int `json:"servicePortHttp"`
	// Service port for Dex Server GRPC
	ServicePortGrpc int `json:"servicePortGrpc"`
	// Describe a container image
	Image ImageStruct `json:"image"`
	// Describe a container image
	InitImage ImageStruct `json:"initImage"`
	// Additional volume mounts
	VolumeMounts []corev1.VolumeMount `json:"volumeMounts,omitempty"`
	// Additional volumes
	Volumes []corev1.VolumeSource `json:"volumes,omitempty"`
	// Annotations for the server deployment
	Annotations map[string]string `json:"anntotations,omitempty"`
}

type RepoServerStruct struct {
	// Container port for repo server
	ContainerPort int `json:"containerPort"`
	// Service port for repo server
	ServicePort int `json:"servicePort"`
	// Describe a container image
	Image ImageStruct `json:"image"`
	// Additional volume mounts
	// +optional
	VolumeMounts []corev1.VolumeMount `json:"volumeMounts,omitempty"`
	// Additional volumes
	// +optional
	Volumes []corev1.VolumeSource `json:"volumes,omitempty"`
}
type ApplicationControllerStruct struct {
	// Container port for application controller server and metrics
	ContainerPort int `json:"containerPort"`
	// Service port for applicaiton controller server
	ServicePort int `json:"servicePort"`
	// Describe a container image
	Image ImageStruct `json:"image"`
	// Additional volume mounts
	// +optional
	VolumeMounts []corev1.VolumeMount `json:"volumeMounts,omitempty"`
	// Additional volumes
	// +optional
	Volumes []corev1.VolumeSource `json:"volumes,omitempty"`
}

type ServerStruct struct {
	// Container port for server
	ContainerPort int `json:"containerPort"`
	// Container port for server metrics
	MetricsPort int `json:"metricsPort"`
	// HTTP Container port for server
	ServicePortHttp int `json:"servicePortHttp"`
	// HTTPS Container port for server
	ServicePortHttps int `json:"servicePortHttps"`
	// Annotations for server service
	ServiceAnnotations map[string]string `json:"serviceAnnotations,omitempty"`
	// Describe a container image
	Image ImageStruct `json:"image"`
	// Service Type string describes ingress methods for a service
	ServiceType corev1.ServiceType `json:"serviceType"`
	// Add additional arguments
	ExtraArgs []string `json:"extraArgs,omitempty"`
	// Additional volume mounts
	VolumeMounts []corev1.VolumeMount `json:"volumeMounts,omitempty"`
	// Additional volumes
	Volumes []corev1.VolumeSource `json:"volumes,omitempty"`
	// Annotations for the server deployment
	Annotations map[string]string `json:"anntotations,omitempty"`
}

type ImageStruct struct {
	// Docker image pull policy
	PullPolicy corev1.PullPolicy `json:"pullPolicy"`
	// Docker image repo
	Repository string `json:"repository"`
	// Docker image tag
	Tag string `json:"tag"`
}

// ArgoCDStatus defines the observed state of ArgoCD
// +k8s:openapi-gen=true
type ArgoCDStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ArgoCD is the Schema for the argocds API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type ArgoCD struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ArgoCDSpec   `json:"spec,omitempty"`
	Status ArgoCDStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ArgoCDList contains a list of ArgoCD
type ArgoCDList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ArgoCD `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ArgoCD{}, &ArgoCDList{})
}
