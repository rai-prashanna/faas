package app

import (
	"fmt"
	"math/big"
	"net"
	"net/http"
	"strconv"
	"os"


	// local packages
	"helpers"
	"app/config"

	// vendor packages

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func RunApp() {
	result := helpers.Sum("123", "456")
	fmt.Printf("123 + 456 = %d\n", result)

	cfg := config.Load()
	port, _ := cfg["port"].(string)
	host, _ := cfg["host"].(string)

	e := echo.New()
	e.GET("/hello", func(c echo.Context) error {
		//return c.String(http.StatusOK, "Hello, World!")
		// Get team and member from the query string
		strnum := c.QueryParam("num")
		num, _ := strconv.ParseInt(strnum, 10, 64)

		var fact big.Int

		result:= fact.MulRange(1,num)

		return c.String(http.StatusOK, "result:" + result.String() + " ")

	})

	e.GET("/", func(c echo.Context) error {
		//return c.String(http.StatusOK, "Hello, World!")
		// Get team and member from the query string
		
		
		return c.String(http.StatusOK, "hello my friend ")

	})
	e.GET("/dig", func(c echo.Context) error {
			url := c.QueryParam("url")
			ips:=dig(url)
		s := []string{}
		for _, ip := range ips {
			fmt.Printf("google.com. IN A %s\n", ip.String())
			s = append(s, fmt.Sprintf("%s", ip.String()))		}
		return c.JSON(http.StatusCreated, ips)
	})
	fmt.Printf("Start running ... %s:%s\n", host, port)

	e.Run(standard.New(host + ":" + port))
}

func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:" + team + ", member:" + member)
}

func dig(url string) []net.IP{
	ips, err := net.LookupIP(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err)
		os.Exit(1)
	}
	//for _, ip := range ips {
	//	fmt.Printf("google.com. IN A %s\n", ip.String())
	//
	//}
	return ips
}
