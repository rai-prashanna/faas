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
		log.Println("the url is ...")

		log.Print(req.RequestURI)
		resourceurl := strings.Split(req.RequestURI, "?")

		queryparam := resourceurl[1]
		functionname := strings.Trim(resourceurl[0],"/")

		targetIP, targetport := getSocketOfContainerByLabel(functionname)
		targeturl := "http://"+targetIP+":"+targetport
		log.Println(targeturl)

		log.Println("suceessfully found service ")
		log.Println(targeturl)

		target, err := url.Parse(targeturl)
		queryurl := "/?"+queryparam
		proxyqueryurl, err := url.Parse(queryurl)

		if err != nil {
			log.Fatal(err)
		}
		proxy := httputil.NewSingleHostReverseProxy(target)
		req.URL=proxyqueryurl

		proxy.ServeHTTP(w, req)
		log.Println("forwarded incoming request to service")

	})
	log.Println("listening incoming request on port 8080 ")
	log.Println("hit http://localhost:8080/factorialservice?num=3 or")
	log.Println("hit http://localhost:8080/digservice?url=www.wwe.com")
	err := http.ListenAndServe(":"+listenerPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}




