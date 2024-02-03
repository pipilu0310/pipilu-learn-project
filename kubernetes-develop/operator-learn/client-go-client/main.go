package main

import (
	"context"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// RESTClient，可以直接创建RESTClient，操作任意类型的资源
	//// config
	//config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedConfigPathFlag)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// client
	//restClient, err := rest.RESTClientFor(config)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// get data
	//pod := v1.Pod{}
	//err = restClient.Get().Namespace("default").Resource("pods").Name("test").Do(context.TODO()).Into(&pod)
	//if err != nil {
	//	println(err)
	//} else {
	//	println(pod.Name)
	//}

	// clientSet 可以直接创建clientSet，但是需要手动指定API版本和资源类型
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedConfigPathFlag)
	if err != nil {
		panic(err)
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	podObj, err := clientSet.CoreV1().Pods("default").Get(context.TODO(), "test", v1.GetOptions{})
	if err != nil {
		panic(err)
	} else {
		println(podObj.Name)
	}
}
