package app

import (
	"fmt"
	// local packages

	// vendor packages
	"golang.org/x/net/context"
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"

)

func RunApp() {
	fmt.Println("started ......\n")
	clientID,port:=getSocketOfContainerByLabel("factorialservice")
	fmt.Println("Key:", clientID, "Value:", port)

	fmt.Printf("ended ......\n", 2313)
}



func getSocketOfContainerByLabel(faasName string) (string,string){
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	filters := filters.NewArgs()
	filters.Add("label", "faas.name="+faasName)
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{
		Size:    true,
		All:     true,
		Since:   "container",
		Filters: filters,
	})
	if err != nil {
		panic(err)
	}

	if len(containers) > 0 {
		first_container :=containers[0]
		labels := first_container.Labels
		// fmt.Println(containers[0].ID.s)
	//	fmt.Println(first_container.ID[:10],labels["faas.port"])

		return first_container.ID[:10],labels["faas.port"]
	} else {
		fmt.Println("There are no containers running")
		fmt.Println("you need to implement logic to launch container")
	}
	return "",""
}