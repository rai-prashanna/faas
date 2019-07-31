package app

import (
	"fmt"
	// local packages

	// vendor packages
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"

	"golang.org/x/net/context"
)

func RunApp() {
	fmt.Printf("123 + 456 = %d\n", 2313)
	main()
}

func main() {

	//ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	//hostBinding := nat.PortBinding{
	//	HostIP:   "127.0.0.0",
	//	HostPort: "8000",
	//}
	//
	//containerPort, err := nat.NewPort("tcp", "80")
	//if err != nil {
	//	panic("Unable to get the port")
	//}
	//
	//portBinding := nat.PortMap{containerPort: []nat.PortBinding{hostBinding}}
	//
	//cont, err := cli.ContainerCreate(ctx, &container.Config{
	//	Image: "nginx",
	//}, &container.HostConfig{
	//		PortBindings: portBinding,
	//	}, nil, "")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Container %s is started", cont.ID)
	//containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	//{"label":{"faas.name":"factorialservice"}}
//	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{map[string][]string{"label": {"faas.name:factorialservice"}}})
//	containers, err := cli.ListContainers(types.ListContainersOptions{All: true, Filters: map[string][]string{"label": {"faas.name:factorialservice"}}})
	filters := filters.NewArgs()
	filters.Add("label", "faas.name=factorialservice")
	filters.Add("label", "faas.port=7070")

	//"faas.name": "factorialservice"
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
		for _, container := range containers {
			fmt.Printf("%s %s\n", container.ID[:10], container.ID)
			r, err := cli.ContainerInspect(context.Background(), container.ID)
			//d := json.NewDecoder(r)
			if err != nil {
				panic(err)
			}
			//pnetworks := r.NetworkSettings   ContainerJSONBase.HostConfig.Binds
			fmt.Printf("%s %s\n", r.NetworkSettings)
			//fmt.Printf("%s %s\n", d)

		}
	} else {
		fmt.Println("There are no containers running")
	}


}

