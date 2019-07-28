package app

import (
	"fmt"
	"math/big"
	"net/http"
	"strconv"

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
	e.GET("/factorial", func(c echo.Context) error {
		strnum := c.QueryParam("num")
		num, _ := strconv.ParseInt(strnum, 10, 64)

		var fact big.Int

		result := fact.MulRange(1, num)

		return c.String(http.StatusOK, "result:"+result.String()+" ")

	})

	fmt.Printf("Factorail service running ... %s:%s\n", host, port)
	e.Run(standard.New(host + ":" + port))

}


