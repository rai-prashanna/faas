package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

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
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		req.Host = req.URL.Host
		log.Printf("received request on proxy with url ...  %s  \n",req.RequestURI)
		requestedurl := strings.Split(req.RequestURI, "?")

		queryparam := requestedurl[1]
		functionname := strings.Trim(requestedurl[0],"/")

		log.Printf("lookup for %s service on ...\n",functionname)
		targetIP, targetport := getSocketOfContainerByLabel(functionname)
		targeturl := "http://"+targetIP+":"+targetport
		log.Printf("suceessfully found %s service \n",functionname)

		target, err := url.Parse(targeturl)
		queryurl := "/?"+queryparam
		proxyqueryurl, err := url.Parse(queryurl)

		if err != nil {
			log.Fatal(err)
		}
		proxy := httputil.NewSingleHostReverseProxy(target)
		req.URL=proxyqueryurl

		proxy.ServeHTTP(w, req)
		log.Printf("forwarded incoming request to %s  \n",functionname)

	})
	log.Println("listening incoming request on port 8080 ")
	log.Println("hit http://localhost:8080/factorialservice?num=3 or")
	log.Println("hit http://localhost:8080/digservice?url=www.wwe.com")
	err := http.ListenAndServe(":"+listenerPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}




