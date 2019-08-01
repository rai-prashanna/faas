package app

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"

	// local packages
	// vendor packages
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"

)

func RunApp() {
	var port string = "6060"

	ip, err := externalIP()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ip)

	e := echo.New()
	e.GET("/dig", func(c echo.Context) error {
		url := c.QueryParam("url")
		ips:=dig(url)
		return c.JSON(http.StatusCreated, ips)
	})

	fmt.Printf("dig service running ... %s:%s\n", ip, port)

	e.Run(standard.New(ip + ":" + port))

}


func dig(url string) []net.IP{
	ips, err := net.LookupIP(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err)
		os.Exit(1)
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

