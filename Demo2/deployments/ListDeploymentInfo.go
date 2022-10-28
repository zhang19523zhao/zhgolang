package deployments

import (
	"context"
	"fmt"
	appv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ListDplInfo struct {
	clientset *kubernetes.Clientset
}

func NewListDplInfo(clientset *kubernetes.Clientset) *ListDplInfo {
	return &ListDplInfo{
		clientset: clientset,
	}
}

func (this *ListDplInfo) GetDep() *appv1.DeploymentList {
	ctx := context.Background()
	listopt := metav1.ListOptions{}
	deploymentList, err := this.clientset.AppsV1().Deployments("default").List(ctx, listopt)
	if err != nil {
		fmt.Println(err)
	}
	return deploymentList
}

// 打印出deployment 名字
func (this *ListDplInfo) Name() {

	depList := this.GetDep()
	// 遍历获取deployment 名字
	for _, deployment := range depList.Items {
		fmt.Println(deployment.Name)
	}
}

// 打印出 deployment 副本数
func (this *ListDplInfo) Replicas() {
	depList := this.GetDep()
	// 遍历获取deployment 名字
	for _, deployment := range depList.Items {
		fmt.Println(deployment.Name,
			deployment.Status.Replicas,            // 总副本数
			deployment.Status.AvailableReplicas,   // 可用副本数
			deployment.Status.UnavailableReplicas, // 不可用副本数
		)
	}
}
