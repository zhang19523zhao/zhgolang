package deployments

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type UpdateDep struct {
	clientset *kubernetes.Clientset
}

func NewUpdateDep(clientset *kubernetes.Clientset) *UpdateDep {
	return &UpdateDep{
		clientset: clientset,
	}
}

// 这里修改Replicas 副本数为例
func (this *UpdateDep) UpdateReplicas(depName string, num int32) error {
	ctx := context.Background()
	dep, err := this.clientset.AppsV1().Deployments("default").Get(ctx, depName, metav1.GetOptions{})
	if err != nil {
		return err
	}
	// 修改副本数
	dep.Spec.Replicas = &num
	_, err = this.clientset.AppsV1().Deployments("default").Update(ctx, dep, metav1.UpdateOptions{})

	return err
}
