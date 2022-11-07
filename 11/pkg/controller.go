package pkg

import (
	informer "k8s.io/client-go/informers/core/v1"
	netInformer "k8s.io/client-go/informers/networking/v1"
	"k8s.io/client-go/kubernetes"
	coreLister "k8s.io/client-go/listers/core/v1"
	netLister "k8s.io/client-go/listers/networking/v1"
	"k8s.io/client-go/tools/cache"
)

type Controller struct {
	client        kubernetes.Interface
	ingressLister netLister.IngressLister
	serviceLister coreLister.ServiceLister
}

func (c *Controller) addService() func(obj interface{}) {

}

func (c *Controller) updateService() func(oldObj interface{}, newObj interface{}) {

}

func (c *Controller) deleteIngress() func(obj interface{}) {

}

func (c *Controller) Run(stopCh chan struct{}) {
	<-stopCh
}

func NewController(client kubernetes.Interface, ingressInformer netInformer.IngressInformer, serviceInformer informer.ServiceInformer) Controller {
	c := Controller{
		client:        client,
		ingressLister: ingressInformer.Lister(),
		serviceLister: serviceInformer.Lister(),
	}

	serviceInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: c.addService(),

		UpdateFunc: c.updateService(),
	})

	ingressInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		DeleteFunc: c.deleteIngress(),
	})

	return c
}
