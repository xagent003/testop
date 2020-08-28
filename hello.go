package main

import (
	"bytes"
	"fmt"
    "io"
	"io/ioutil"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	extv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
    //"github.com/ghodss/yaml"
    "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type MultusManifest struct {
	Namespace          corev1.Namespace               `json:"namespace,omitempty"`
	CustomResourceDefinition  extv1.CustomResourceDefinition `json:"customresourcedefinition,omitempty"`
	ClusterRole        rbacv1.ClusterRole             `json:"clusterrole,omitempty"`
	ClusterRoleBinding rbacv1.ClusterRoleBinding      `json:"clusterrolebinding,omitempty"`
    ServiceAccount     corev1.ServiceAccount          `json:"serviceaccount,omitempty"`
    ConfigMap          corev1.ConfigMap               `json:"configmap,omitempty"`
	DaemonSet          appsv1.DaemonSet             `json:"daemonset,omitempty"`
}

type TestPodManifest struct {
	Pod corev1.Pod `json:"pod,omitempty"`
}


func main() {
    yamlfile, err := ioutil.ReadFile("/home/xagent/go/src/test2/plugins/multus/multus2.yaml")
    fmt.Printf("raw: %s\n", string(yamlfile))
	if err != nil {
		fmt.Println("Failed to read file")
	} else {
		fmt.Println("read the file")
	}

    resource_list := []*unstructured.Unstructured{}
	decoder := yaml.NewYAMLOrJSONDecoder(bytes.NewReader(yamlfile), 4096)
    for {
        resource := unstructured.Unstructured{}
        err := decoder.Decode(&resource)
        if err == nil {
            fmt.Printf("appended +%v\n\n", resource)
            resource_list = append(resource_list, &resource)
        } else if err == io.EOF {
            break;
        } else {
            fmt.Printf("Error decoding!!!: %s\n", err)
        }
    }

	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
	}
    fmt.Printf("unstructured list: \n+%v\n\n", resource_list)

    return
}
