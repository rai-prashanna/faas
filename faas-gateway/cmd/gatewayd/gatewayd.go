package main

import (
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

func main() {
	log.Println("Starting FAAS-GATEWAY Service")
	ProxyHandlar()
	log.Println("Stoping FAAS-GATEWAY Service")
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
		log.Println("There are no containers running")
		log.Println("you need to implement logic to launch container")
	}
	return "",""
}



func ProxyHandlar() {
	var listenerPort string = "80"
	http.HandleFunc("/factorial", func(w http.ResponseWriter, req *http.Request) {
		req.Host = req.URL.Host
		factorialtargetIP, factorialtargetport := getSocketOfContainerByLabel("factorialservice")
		factorialtargeturl := "http://"+factorialtargetIP+":"+factorialtargetport
		log.Println("suceessfully found factorial-service ")
		factorialtarget, err := url.Parse(factorialtargeturl)
		if err != nil {
			log.Fatal(err)
		}
		factorialproxy := httputil.NewSingleHostReverseProxy(factorialtarget)
		factorialproxy.ServeHTTP(w, req)
		log.Println("forwarded incoming request to factorial-service")

	})
	http.HandleFunc("/dig", func(w http.ResponseWriter, req *http.Request) {
		req.Host = req.URL.Host
		digtargetIP, digtargetport := getSocketOfContainerByLabel("digservice")
		digtargeturl := "http://"+digtargetIP+":"+digtargetport
		log.Println("suceessfully found dig-service ")
		digtarget, err := url.Parse(digtargeturl)
		if err != nil {
			log.Fatal(err)
		}
		digproxy := httputil.NewSingleHostReverseProxy(digtarget)
		digproxy.ServeHTTP(w, req)
		log.Println("forwarded incoming request to dig-service")
	})
	log.Println("listening incoming request on port 8080 ")
	log.Println("hit http://localhost:8080/factorial?num=3 or")
	log.Println("hit http://localhost:8080/dig?url=www.wwe.com")
	err := http.ListenAndServe(":"+listenerPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}



