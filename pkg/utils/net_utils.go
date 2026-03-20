package utils

import (
	"errors"
	"net"
)

func GetLocalIp() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", errors.New("cannot find local IP address")
}

func CreateServiceUnitName(serviceName string) string {
	localIp, err := GetLocalIp()
	if err != nil {
		return "unknown"
	}
	return serviceName + "/" + localIp
}
