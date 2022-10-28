package main

import (
	"Demo2/client"
	"Demo2/deployments"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	clientset := client.K8sClientset

	deplist := deployments.NewListDplInfo(clientset)
	//获取 deployment 名字
	deplist.Name()
	打印deployment 副本数
	deplist.Replicas()

	// 根据yaml创建deployment
	depcreate := deployments.NewCreateDplForYaml(clientset)
	fmt.Println(depcreate.CreateDep("yamls/nginx.yaml"))

	fmt.Println(clientset.AppsV1().Deployments("default").Delete(context.Background(), "web1", metav1.DeleteOptions{}))

	// 更新deployment的副本数
	deployments.NewUpdateDep(clientset).UpdateReplicas("web1", 3)
	deplist.Replicas()

}
