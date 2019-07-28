package helpers

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"github.com/sparrc/go-ping"
	"time"
)

func Sum(s1 string, s2 string) int {
	i, _ := strconv.Atoi(s1)
	j, _ := strconv.Atoi(s2)

	return i + j
}

func Pingme() {
	pinger, err := ping.NewPinger("www.google.com")
	if err != nil {
		panic(err)
	}
	pinger.Count = 3
	pinger.Run() // blocks until finished
	stats := pinger.Statistics() // get send/receive/rtt stats
	fmt.Printf("\n", stats)

}