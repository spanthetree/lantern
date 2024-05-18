package speedtest

import (
	"fmt"

	"github.com/showwin/speedtest-go/speedtest"
	"github.com/showwin/speedtest-go/speedtest/transport"
)

// Note: The current packet loss analyzer does not support udp over http.
// This means we cannot get packet loss through a proxy.
func main() {
	// Retrieve available servers
	var speedtestClient = speedtest.New()
	serverList, _ := speedtestClient.FetchServers()
	targets, _ := serverList.FindServer([]int{})

	// Create a packet loss analyzer, use default options
	analyzer := speedtest.NewPacketLossAnalyzer(nil)

	// Perform packet loss analysis on all available servers
	for _, server := range targets {
		err := analyzer.Run(server.Host, func(packetLoss *transport.PLoss) {
			fmt.Println(packetLoss, server.Host, server.Name)
			// fmt.Println(packetLoss.Loss())
		})
		checkError(err)
	}

	// or test all at the same time.
	packetLoss, err := analyzer.RunMulti(targets.Hosts())
	checkError(err)
	// fmt.Println(packetLoss)
	return packetloss
}
