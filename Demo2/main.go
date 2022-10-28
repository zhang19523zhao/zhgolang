package main

import (
	"Demo2/client"
	"Demo2/deployments"
	"fmt"
)

func main() {
	clientset := client.K8sClientset

	deplist := deployments.NewListDplInfo(clientset)
	//获取 deployment 名字
	deplist.Name()

	depcreate := deployments.NewCreateDplForYaml(clientset)
	fmt.Println(depcreate.CreateDep("yamls/nginx.yaml"))
}
