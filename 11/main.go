package main

import (
	"github.com/huanmin/client-go-demo/11/pkg"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func main() {

	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		clusterConfig, err := rest.InClusterConfig()
		if err != nil {
			log.Fatalln("can't get config")
		}
		config = clusterConfig
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln("can't create client")
	}

	factory := informers.NewSharedInformerFactory(clientset, 0)
	servicesInformer := factory.Core().V1().Services()
	ingressesInformer := factory.Networking().V1().Ingresses()

	controller := pkg.NewController(clientset, ingressesInformer, servicesInformer)

	stopCh := make(chan struct{})
	factory.Start(stopCh)
	factory.WaitForCacheSync(stopCh)

	controller.Run(stopCh)
}
