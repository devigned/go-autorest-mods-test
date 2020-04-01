package main

import (
	"fmt"

	"github.com/docker/distribution"
	"sigs.k8s.io/cluster-api-provider-azure/api/v1alpha3"
)

func main() {
	globalScope := distribution.GlobalScope
	fmt.Printf("gs: %+v", globalScope)

	foo := v1alpha3.AzureMachine{}
	fmt.Printf("vm: %+v", foo)
}