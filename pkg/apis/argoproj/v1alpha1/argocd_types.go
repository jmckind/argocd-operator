package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.
// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file

// ArgoCDSpec defines the desired state of ArgoCD
// +k8s:openapi-gen=true
// All omitempty in the moment
type ArgoCDSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	ApplicationController ApplicationControllerStruct `json:"applicationController,omitempty"`
	Server                ServerStruct                `json:"server,omitempty"`
	RepoServer            RepoServerStruct            `json:"repoServer,omitempty"`
	DexServer             DexServerStruct             `json:"dexServer,omitempty"`
	Ingress               IngressStruct               `json:"ingress,omitempty"`
	Redis                 RedisStruct                 `json:"redis,omitempty"`
	Rbac                  RbacStruct                  `json:"rbac,omitempty"`
	Certificate           CertificateStruct           `json:"certificate,omitempty"`
	ClusterAdminAccess    ClusterAdminAccessStruct    `json:"clusterAdminAccess,omitempty"`
	Config                ConfigStruct                `json:"config,omitempty"`
}

type ConfigStruct struct {
	// Creates the argocd-secret secret, set to false to manage externally
	CreateSecret       bool   `json:"createSecret"`
	ResourceExclusions string `json:"resourceExclusions"`
	// Configuration for remote Git repositories for Applications
	// +optional
	Repositories string `json:"repositories,omitempty"`
	// Configuration for external Helm charts
	// +optional
	HelmRepositories string `json:"helmRepositories,omitempty"`
	// Configuration for external auth and URL
	// +optional
	DexConfig string `json:"dexConfig,omitempty"`
	// External URL for ArgoCD
	// +optional
	Url string `json:"url,omitempty"`
	// Configuration for OpenID connect
	// +optional
	OidcConfig string `json:"oidcConfig,omitempty"`
	// ResourceCustomizations can be used to create custom health checks for resources
	// +optional
	ResourceCustomizations string `json:"resourceCustomizations,omitempty"`
	// List of custom config management plugins
	// +optional
	ConfigManagementPlugins string `json:"configManagementPlugins,omitempty"`
	// Custom instance label key
	// +optional
	InstanceLabelKey      string `json:"instanceLabelKey,omitempty"`
	EnableAnonymousAccess bool   `json:"enableAnonymousAccess"`
	// +optional
	Webhook WebhookStruct `json:"webhook,omitempty"`
}

type WebhookStruct struct {
	// GitHub incoming webhook secret
	// +optional
	GithubSecret string `json:"githubSecret,omitempty"`
	// GitLab incoming webhook secret
	// +optional
	GitlabSecret string `json:"gitlabSecret,omitempty"`
	// BitBucket incoming webhook secret
	// +optional
	BitbucketSecret string `json:"bitbucketSecret,omitempty"`
}

type ClusterAdminAccessStruct struct {
	// Standard Argo CD installation with cluster-admin access. Set this true if you plan to use Argo CD to deploy applications in the same cluster that Argo CD runs in (i.e. kubernetes.svc.default). Will still be able to deploy to external clusters with inputted credentials.
	Enabled bool `json:"enabled"`
}

type CertificateStruct struct {
	// Enable certificate (requires cert-manager)
	Enabled bool `json:"enabled"`
	// +optional
	Issuer IssuerStruct `json:"issuer,omitempty"`
}

type IssuerStruct struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
}

type RbacStruct struct {
	// RBAC policy in CSV
	// +optional
	PolicyCsv string `json:"policyCvs,omitempty"`
	// The default role Argo CD will fall back to, when authorizing API requests, ie: role:readonly
	// +optional
	PolicyDefault string `json:"policyDefault,omitempty"`
	// Scopes controls which OIDC scopes to examine during rbac enforcement (in addition to sub scope). ie: [groups]
	// +optional
	Scopes string `json:"scopes,omitempty"`
}

type RedisStruct struct {
	// Container port for Redis
	ContainerPort int `json:"containerPort"`
	// Service port for Redis
	ServicePort int `json:"servicePort"`
	// Describe a container image
	Image ImageStruct `json:"image"`
}

type IngressStruct struct {
	// Enable ingress
	Enabled bool `json:"enabled"`
	// Annotations for ingress object
	// +optional
	Annotations map[string]string `json:"anntotations,omitempty"`
	Path        string            `json:"path"`
	// Ingress additional hosts
	// +optional
	AdditionalHosts []string `json:"additionalHosts,omitempty"`
	// Ingress TLS configuration
	// +optional
	Tls map[string]string `json:"tls,omitempty"`
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
	// +optional
	VolumeMounts []corev1.VolumeMount `json:"volumeMounts,omitempty"`
	// Additional volumes
	// +optional
	Volumes []corev1.VolumeSource `json:"volumes,omitempty"`
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
	// +optional
	ServiceAnnotations map[string]string `json:"serviceAnnotations,omitempty"`
	// Describe a container image
	Image ImageStruct `json:"image"`
	// Service Type string describes ingress methods for a service
	ServiceType corev1.ServiceType `json:"serviceType"`
	// Add additional arguments
	// +optional
	ExtraArgs []string `json:"extraArgs,omitempty"`
	// Additional volume mounts
	// +optional
	VolumeMounts []corev1.VolumeMount `json:"volumeMounts,omitempty"`
	// Additional volumes
	// +optional
	Volumes []corev1.VolumeSource `json:"volumes,omitempty"`
	// Annotations for the server deployment
	// +optional
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
