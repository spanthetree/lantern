package consul

import (
	"fmt"
	"log"

	"github.com/hashicorp/consul/api"
	"github.com/spanthetree/lantern/internal/shared"
)

// RegisterService registers the service with Consul under the parent service "lantern"
func RegisterService(localServiceData shared.LocalServiceData) error {
	// import default consul api config
	config := api.DefaultConfig()

	// set the consul server address
	config.Address = "consul-server-1.consul-server.tools.k8s.obfuscated.co:8500"

	// Create a new consul client
	client, err := api.NewClient(config)
	if err != nil {
		return err
	}

	// Check if already registered
	services, err := client.Agent().Services()
	if err != nil {
		return err
	}
	if _, ok := services[localServiceData.ID]; ok {
		log.Println("service is already registered with Consul")
		return nil
	}

	fmt.Println(localServiceData)

	// Define registration for consul
	registration := &api.AgentServiceRegistration{
		ID:      localServiceData.ID,
		Name:    "lantern", // Ensures all services are registered under the parent service "lantern"
		Address: localServiceData.Address,
		Tags:    localServiceData.Tags,
		Meta: map[string]string{
			"fqdn":       localServiceData.FQDN,
			"datacenter": localServiceData.DataCenter,
		},
		Check: &api.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://%s:8181/health", localServiceData.Address),
			Interval: "10s",
			Timeout:  "1s",
			// DeregisterCriticalServiceAfter: "1m",
		},
	}

	if err := client.Agent().ServiceRegister(registration); err != nil {
		return fmt.Errorf("failed to add service to consul %w", err)
	}

	// Log success
	log.Println("successfully added service")
	return nil
}
