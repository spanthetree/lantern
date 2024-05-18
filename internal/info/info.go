package info

import (
	"errors"
	"log"
	"net"
	"os"
)

// get primary IP of the local machine
func GetIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP, err
}

func GetFQDN() (string, error) {
	name, err := os.Hostname()
	if err != nil {
		return "", err
	}
	addrs, err := net.LookupHost(name)
	if err != nil {
		return "", err
	}
	if len(addrs) > 0 {
		return name, nil
	}
	return name, errors.New("no FQDN found")
}
