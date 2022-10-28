package deployments

import (
	"context"
	"fmt"
	appv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ListDplInfo struct {
	Dlist *appv1.DeploymentList
}

func NewListDplInfo(clientset *kubernetes.Clientset) *ListDplInfo {
	ctx := context.Background()
	listopt := metav1.ListOptions{}
	deploymentList, err := clientset.AppsV1().Deployments("default").List(ctx, listopt)
	if err != nil {
		fmt.Println(err)
	}
	return &ListDplInfo{
		deploymentList,
	}
}

func (this *ListDplInfo) Name() {
	// 遍历获取deployment 名字
	for _, deployment := range this.Dlist.Items {
		fmt.Println(deployment.Name)
	}
}
