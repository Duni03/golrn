package main

import (
	"context"
	"flag"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kconfig := flag.String("kubeconfig", "~/.kube/config", "location for kubeconfig")
	config, err := clientcmd.BuildConfigFromFlags("", *kconfig)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T\n", config)
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println(err)
	}
	pods, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Println(err)
	}
	for _, p := range pods.Items {
		fmt.Printf("pod name: %s\n", p.Name)
	}
	nodes, err := clientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Println(err)
	}
	for _, n := range nodes.Items {
		fmt.Printf("NOdes name: %s\n", n.Name)
	}
}

/*
 20  go mod init main
 21  go build main.go
 22  go get k8s.io/client-go@kubernetes
 23  go get k8s.io/client-go@kubernetes-1.29.0
 24  go build main.go
 25  go get k8s.io/client-go/tools/clientcmd
 26  go build main.go
 27  go get k8s.io/client-go/openapi@v0.29.0
 28  go build main.go
 29  go build main.go
 30  go mod init main
 31  go mod tidy
 32  go build main.go
 33  go build main.go
 34  ./main
*/
