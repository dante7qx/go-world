package k8s_test

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"dante-client-go/factory"
	"fmt"
	"time"
)

/**
	列出指定 namespace 的 pod
 */
func ListPod(namespace string) {
	clientset := factory.BuildKubeConfig()
	for {
		pods, err := clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
		for index, pod := range pods.Items {
			fmt.Println(index, pod.Name)
		}

		/*
		_, err = clientset.CoreV1().Pods("default").Get("example-xxxxx", metav1.GetOptions{})
		if errors.IsNotFound(err) {
			fmt.Printf("Pod not found\n")
		} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
			fmt.Printf("Error getting pod %v\n", statusError.ErrStatus.Message)
		} else if err != nil {
			panic(err.Error())
		} else {
			fmt.Printf("Found pod\n")
		}
		*/

		time.Sleep(10 * time.Second)
	}
}
