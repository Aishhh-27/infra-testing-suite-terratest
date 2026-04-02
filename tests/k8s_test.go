package tests

import (
	"context"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraformAndK8s(t *testing.T) {
	opts := &terraform.Options{
		TerraformDir: "../terraform",
	}

	defer terraform.Destroy(t, opts)

	terraform.InitAndApply(t, opts)

	config, err := clientcmd.BuildConfigFromFlags("", "/home/aishwarya/.kube/config")
	if err != nil {
		t.Fatalf("Failed to load kubeconfig: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		t.Fatalf("Failed to create Kubernetes client: %v", err)
	}

	ctx := context.TODO()

	_, err = clientset.CoreV1().Namespaces().Get(
		ctx,
		"test-namespace",
		metav1.GetOptions{},
	)

	if err != nil {
		t.Fatalf("Namespace not found: %v", err)
	}

	t.Log("Namespace exists successfully")
}
