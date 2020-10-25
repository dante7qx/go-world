package main

import k8s "dante-client-go/k8s/object"

func main() {
	//k8s.ListPod("dante-ingress")
	k8s.CreateUpdateDeleteDeployment("dante", "list")
}


