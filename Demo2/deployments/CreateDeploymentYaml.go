package deployments

import (
	"context"
	"encoding/json"
	appv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"os"
	"sigs.k8s.io/yaml"
)

type CreateDplForYaml struct {
	clientset *kubernetes.Clientset
}

func (this *CreateDplForYaml) CreateDep(yamlPath string) error {
	ctx := context.Background()
	Dep := &appv1.Deployment{}

	b, err := os.ReadFile(yamlPath)
	if err != nil {
		return err
	}

	// 将yaml 转化成 josn
	depJson, err := yaml.YAMLToJSON(b)
	if err != nil {
		return err
	}

	// 将json 绑定到Dep上
	if err := json.Unmarshal(depJson, Dep); err != nil {
		return err
	}

	if _, err := this.clientset.AppsV1().Deployments("default").Create(ctx, Dep, metav1.CreateOptions{}); err != nil {
		return err
	}

	return nil
}

func NewCreateDplForYaml(clientset *kubernetes.Clientset) *CreateDplForYaml {
	return &CreateDplForYaml{
		clientset: clientset,
	}
}
