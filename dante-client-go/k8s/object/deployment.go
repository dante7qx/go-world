package k8s_test

import (
	"dante-client-go/factory"
	"fmt"
	"k8s.io/client-go/util/retry"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
)

/**
	创建 deployment
 */
func CreateUpdateDeleteDeployment(namespace string, oper string) {
	clientset := factory.BuildKubeConfig()
	deploymentsClient := clientset.AppsV1().Deployments(namespace)
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "client-go-deployment",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(3),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "clientgo-demo",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "clientgo-demo",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: "tomcat:8.5",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 8080,
								},
							},
						},
					},
				},
			},
		},
	}

	switch oper {
	case "create": {
		fmt.Println("create deployment")
		result, err := deploymentsClient.Create(deployment)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
	}
	case "update": {
		fmt.Println("update deployment")
		retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
			// Retrieve the latest version of Deployment before attempting update
			// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
			result, getErr := deploymentsClient.Get("client-go-deployment", metav1.GetOptions{})
			if getErr != nil {
				panic(fmt.Errorf("Failed to get latest version of Deployment: %v", getErr))
			}

			result.Spec.Replicas = int32Ptr(2)                           // reduce replica count
			result.Spec.Template.Spec.Containers[0].Image = "tomcat:9" 	// change tomcat version
			_, updateErr := deploymentsClient.Update(result)
			return updateErr
		})
		if retryErr != nil {
			panic(fmt.Errorf("Update failed: %v", retryErr))
		}

	}
	case "delete": {
		fmt.Println("delete deployment")
		deletePolicy := metav1.DeletePropagationForeground
		if err := deploymentsClient.Delete("client-go-deployment", &metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		}); err != nil {
			panic(err)
		}
	}
	case "list": {
		list, err := deploymentsClient.List(metav1.ListOptions{})
		if err != nil {
			panic(err)
		}
		for _, d := range list.Items {
			fmt.Printf(" * %s (%d replicas)\n", d.Name, *d.Spec.Replicas)
		}

	}

	}

}


func int32Ptr(i int32) *int32 {
	return &i
}
