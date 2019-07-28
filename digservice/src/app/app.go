package app

import (
	"fmt"
	"net"
	"net/http"
	"os"

	// local packages
	"app/config"
	// vendor packages
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"

)

func RunApp() {
	cfg := config.Load()
	port, _ := cfg["port"].(string)
	host, _ := cfg["host"].(string)

	e := echo.New()
	e.GET("/dig", func(c echo.Context) error {
		url := c.QueryParam("url")
		ips:=dig(url)
		return c.JSON(http.StatusCreated, ips)
	})

	fmt.Printf("dig service running ... %s:%s\n", host, port)

	e.Run(standard.New(host + ":" + port))

}


func dig(url string) []net.IP{
	ips, err := net.LookupIP(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err)
		os.Exit(1)
	}
	return ips
}