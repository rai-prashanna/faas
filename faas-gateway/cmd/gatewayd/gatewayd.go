package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
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



func getSocketOfContainerByLabel(faasName string) (string,string,error){
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
		log.Fatal(err)
	}

	if len(containers) > 0 {
		first_container :=containers[0]
		labels := first_container.Labels
		return first_container.ID[:12],labels["faas.port"],nil
	} else {
		log.Println("There are no containers running")
		log.Println("you need to implement logic to launch container")
	}
	return "","",errors.New("SERVICE NOT FOUND ")
}

//  ^function\/:\./$
var baseurlpattern = regexp.MustCompile(`^\/function\/:{1}`)
func ProxyHandlar() {
	var listenerPort string = "80"
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		req.Host = req.URL.Host
		log.Printf("received request on proxy with url ...  %s  \n",req.RequestURI)
		if baseurlpattern.MatchString(req.RequestURI) {
			split := baseurlpattern.Split(req.RequestURI, -1)
			fmt.Println(split[1])


			requestedurl := strings.Split(split[1], "?")

			queryparam := requestedurl[1]
			functionname := strings.Trim(requestedurl[0],"/")

			log.Printf("lookup for %s service on ...\n",functionname)
			targetIP, targetport,err := getSocketOfContainerByLabel(functionname)
			if err != nil {
				log.Fatal(err)
			}
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
		}else {
			defaultMsg()
		}
	})
	defaultMsg()
	err := http.ListenAndServe(":"+listenerPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func defaultMsg()  {
	log.Println("listening incoming request on port 8080 ")
	log.Println("hit http://localhost:8080/function/:factorialservice?num=3 or")
	log.Println("hit http://localhost:8080/function/:digservice?url=www.wwe.com")
}

