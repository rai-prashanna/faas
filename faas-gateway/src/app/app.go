package app

import (
	"fmt"

	// local packages

	// vendor packages
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types/container"

	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
	"github.com/docker/go-connections/nat"
)

func RunApp() {
	fmt.Printf("123 + 456 = %d\n", 2313)
	main()
}

func main() {

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	hostBinding := nat.PortBinding{
		HostIP:   "127.0.0.0",
		HostPort: "8000",
	}
	containerPort, err := nat.NewPort("tcp", "80")
	if err != nil {
		panic("Unable to get the port")
	}

	portBinding := nat.PortMap{containerPort: []nat.PortBinding{hostBinding}}

	cont, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "nginx",
	}, &container.HostConfig{
			PortBindings: portBinding,
		}, nil, "")
	if err != nil {
		panic(err)
	}


	cli.ContainerStart(context.Background(), cont.ID, types.ContainerStartOptions{})
	fmt.Printf("Container %s is started", cont.ID)


}