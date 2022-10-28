package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func main() {
	//k8s 集群访问文件路径
	configPath := "etc/kube.conf"

	//在 kubeconfig 中使用当前上下文环境，config 获取支持 url 和 path 方式，通过 BuildConfigFromFlags() 函数获取 restclient.Config 对象，用来下边根据该 config 对象创建 client 集合
	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		fmt.Println(err)
	}

	//根据获取的 config 来创建一个 clientset 对象。通过调用 NewForConfig 函数创建 clientset 对象。
	//NewForConfig 函数具体实现就是初始化 clientset 中的每个 client，基本涵盖了 k8s 内各种类型
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	// 获取default命名空间下所有的pod
	podList, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	for _, pod := range podList.Items {
		// 打印pod的名字
		fmt.Println(pod.Name)
	}

}
