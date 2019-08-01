package app

import (
	"errors"
	"fmt"
	"math/big"
	"net"
	"net/http"
	"strconv"
	"log"


	// local packages
//	"app/config"
	// vendor packages
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"

)

func RunApp() {
	var port string = "7070"

	ip, err := externalIP()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ip)


	e := echo.New()
	e.GET("/factorial", func(c echo.Context) error {
		strnum := c.QueryParam("num")
		num, _ := strconv.ParseInt(strnum, 10, 64)
		log.Println("received request from proxy ito factorail proxy")
		var fact big.Int

		result := fact.MulRange(1, num)
		log.Println("the result ")
		log.Println(result.String())



		return c.String(http.StatusOK, "result:"+result.String()+" ")

	})

fmt.Printf("Factorail service running ... %s:%s\n", ip,port)
	e.Run(standard.New(ip + ":" + port))

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

