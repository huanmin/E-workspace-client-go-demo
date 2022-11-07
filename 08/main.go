package main

import (
	"fmt"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/workqueue"
)

func main() {
	//创建一个config
	//创建一个client
	//获取 informor
	//添加事件处理器
	//开始执行

	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	factory := informers.NewSharedInformerFactory(clientset, 0)
	informer := factory.Core().V1().Pods().Informer()

	rateLimitingQueue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			fmt.Println("Add Event")
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err != nil {
				panic(err)
			}
			rateLimitingQueue.AddRateLimited(key)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			fmt.Println("Update Event")
			key, err := cache.MetaNamespaceKeyFunc(newObj)
			if err != nil {
				panic(err)
			}
			rateLimitingQueue.AddRateLimited(key)
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Println("Delete Event")
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err != nil {
				panic(err)
			}
			rateLimitingQueue.AddRateLimited(key)
		},
	})

	stopCh := make(chan struct{})
	factory.Start(stopCh)
	factory.WaitForCacheSync(stopCh)
	<-stopCh
}
