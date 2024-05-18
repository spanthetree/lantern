package main

import (
	"log"
	"os"

	"github.com/spanthetree/lantern/internal/consul"
	"github.com/spanthetree/lantern/internal/info"
	"github.com/spanthetree/lantern/internal/shared"

	// "github.com/spanthetree/lantern/internal/iperf"
	// "github.com/spanthetree/lantern/internal/ping"
	"github.com/spanthetree/lantern/internal/prometheus"
	// "github.com/spanthetree/lantern/internal/traceroute"
	// "github.com/spanthetree/lantern/internal/speedtest"
)

func main() {
	localAddr, err := info.GetIP()
	if err != nil {
		log.Fatalf("failed to get primary IP: %v", err)
	}
	fqdn, err := info.GetFQDN()
	if err != nil {
		log.Fatalf("failed to get FQDN: %v", err)
	}
	idBytes, err := os.ReadFile("/etc/machine-id")
	if err != nil {
		log.Fatalf("cannot read file: %v", err)
	}
	id := string(idBytes)

	LocalServiceData := shared.LocalServiceData{
		ID:         id,
		Name:       "lantern",
		FQDN:       fqdn,
		Address:    localAddr,
		Tags:       []string{"ping", "traceroute", "iperf"},
		DataCenter: "sol",
	}

	// Register service in consul
	err = consul.RegisterService(LocalServiceData)
	if err != nil {
		log.Fatalf("Failed to register with consul: %v", err)
	}

	// Start prometheus module
	prometheus.Start()

}
