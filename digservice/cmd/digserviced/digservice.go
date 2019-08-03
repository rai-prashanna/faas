package main

import (
	"errors"
	"net"
	"net/http"
	"log"

	// local packages
	// vendor packages
	"github.com/labstack/echo"

)

func main() {
	var port string = "6060"

	ip, err := externalIP()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("successfully found ip of dig-service host")
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		url := c.QueryParam("url")
		log.Println("received request from proxy into dig-service")
		ips:=dig(url)
		return c.JSON(http.StatusCreated, ips)
	})
	log.Printf("dig service is up and running ... %s:%s\n", ip, port)
	e.Start(ip + ":" + port)
}


func dig(url string) []net.IP{
	ips, err := net.LookupIP(url)
	if err != nil {
		log.Fatalf("Could not get IPs: %v\n",err)
	}
	return ips
}


func externalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}

