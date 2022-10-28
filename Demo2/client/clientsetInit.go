package client

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

var K8sClientset *kubernetes.Clientset

func init() {
	configPath := "etc/kube.conf"
	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		fmt.Println(err)
	}

	K8sClientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

}
