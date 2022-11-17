package utils

import (
	"bytes"
	"github.ocm/kubebuilder-demo/api/v1beta1"
	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	netv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"text/template"
)

func parseTemplate(tmp string, app *v1beta1.App) []byte {
	temp, err := template.ParseFiles("controller/template/" + tmp + ".yml")
	if err != nil {
		panic(err)
	}
	b := new(bytes.Buffer)
	err = temp.Execute(b, app)
	if err != nil {
		panic(err)
	}
	return b.Bytes()
}

func NewDeployment(app *v1beta1.App) *appv1.Deployment {
	d := &appv1.Deployment{}
	err := yaml.Unmarshal(parseTemplate("deployment", app), d)
	if err != nil {
		panic(err)
	}
	return d
}

func NewService(app *v1beta1.App) *corev1.Service {
	s := &corev1.Service{}
	err := yaml.Unmarshal(parseTemplate("service", app), s)
	if err != nil {
		panic(err)
	}
	return s
}

func NewIngress(app *v1beta1.App) *netv1.Ingress {
	i := &netv1.Ingress{}
	err := yaml.Unmarshal(parseTemplate("service", app), i)
	if err != nil {
		panic(err)
	}
	return i
}
