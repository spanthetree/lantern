package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/spanthetree/lantern/internal/info"
	// "github.com/spanthetree/lantern/internal/consul"
	// "github.com/spanthetree/lantern/internal/iperf"
	// "github.com/spanthetree/lantern/internal/ping"
	// "github.com/spanthetree/lantern/internal/prometheus"
	// "github.com/spanthetree/lantern/internal/traceroute"
	// "github.com/spanthetree/lantern/internal/speedtest"
)

type LocalServiceData struct {
	ID         string
	Name       string
	Address    net.IP
	Tags       []string
	DataCenter string
}

func main() {
	primaryIP, err := info.GetIP()
	if err != nil {
		log.Fatalf("failed to get primary IP: %v", err)
	}
	fqdn, err := info.GetFQDN()
	if err != nil {
		log.Fatalf("failed to get FQDN: %v", err)
	}
	id, err := os.ReadFile("/etc/machine-id")
	if err != nil {
		log.Fatalf("cannot read file: %v", err)
	}
	idStr := string(id)

	localServiceData := LocalServiceData{
		ID:         idStr,
		Name:       fqdn,
		Address:    primaryIP,
		Tags:       []string{"ping", "traceroute", "iperf"},
		DataCenter: "sol",
	}

	fmt.Printf("%+v\n", localServiceData)
}
