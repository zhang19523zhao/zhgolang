package deployments

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type DeleteDep struct {
	clientset *kubernetes.Clientset
}

func NewDeleteDep(clientset *kubernetes.Clientset) *DeleteDep {
	return &DeleteDep{
		clientset: clientset,
	}
}

// 删除Deployment
func (this *DeleteDep) Delete(dploymentName string) error {
	return this.clientset.AppsV1().Deployments("default").Delete(context.Background(), dploymentName, metav1.DeleteOptions{})
}
