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

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}

		}
	}
	return "", errors.New("Can not find the client ip address!")
}

func CreateServiceUnitName(serviceName string) string {
	localIp, err := GetLocalIp()
	if err != nil {
		return "nil"
	}
	return serviceName + ":" + localIp

}
