package app

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	// local packages
	// vendor packages
	"golang.org/x/net/context"
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
)

func RunApp() {
	fmt.Println("started ......\n")
	ProxyHandlar()
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
		return first_container.ID[:12],labels["faas.port"]
	} else {
		fmt.Println("There are no containers running")
		fmt.Println("you need to implement logic to launch container")
	}
	return "",""
}



func ProxyHandlar() {
	http.HandleFunc("/factorial", func(w http.ResponseWriter, req *http.Request) {
		req.Host = req.URL.Host // if you remove this line the request will fail... I want to debug why.
		factorialtargetIP, factorialtargetport := getSocketOfContainerByLabel("factorialservice")
		factorialtargeturl := "http://"+factorialtargetIP+":"+factorialtargetport
		log.Println("forwarding to factorial proxy")
		log.Println(factorialtargeturl)

		factorialtarget, err := url.Parse(factorialtargeturl)
		if err != nil {
			log.Fatal(err)
		}
		factorialproxy := httputil.NewSingleHostReverseProxy(factorialtarget)
		factorialproxy.ServeHTTP(w, req)


	})
	http.HandleFunc("/dig", func(w http.ResponseWriter, req *http.Request) {
		req.Host = req.URL.Host // if you remove this line the request will fail... I want to debug why.
		digtargetIP, digtargetport := getSocketOfContainerByLabel("factorialservice")
		digtargeturl := "http://"+digtargetIP+":"+digtargetport
		digtarget, err := url.Parse(digtargeturl)
		if err != nil {
			log.Fatal(err)
		}
		digproxy := httputil.NewSingleHostReverseProxy(digtarget)
		digproxy.ServeHTTP(w, req)
	})
	err := http.ListenAndServe(":80", nil)

	if err != nil {
		panic(err)
	}
}



