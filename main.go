package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

// GetLocalIPv4 获取本地ipv4地址
func GetLocalIPv4() (ip string, err error) {
	addrList, err := net.InterfaceAddrs()
	if err != nil {
		return ip, err
	}
	for _, addr := range addrList {
		// 过滤掉回环地址
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			ipv4 := ipNet.IP.To4()
			// 如果ip不符合ipv4格式，继续查找下一个
			if ipv4 == nil {
				continue
			}
			return ipv4.String(), nil
		}
	}
	return ip, errors.New("no find ip address")
}

func GetExternalIP() (string, error) {
	// 有很多类似网站提供这种服务，这是我知道且正在用的
	// 备用：https://myexternalip.com/raw （cip.cc 应该是够快了，我连手机热点的时候不太稳，其他自己查）
	response, err := http.Get("http://ip.cip.cc")
	if err != nil {
		return "", errors.New("external IP fetch failed, detail:" + err.Error())
	}

	defer response.Body.Close()
	res := ""

	// 类似的API应当返回一个纯净的IP地址
	for {
		tmp := make([]byte, 32)
		n, err := response.Body.Read(tmp)
		if err != nil {
			if err != io.EOF {
				return "", errors.New("external IP fetch failed, detail:" + err.Error())
			}
			res += string(tmp[:n])
			break
		}
		res += string(tmp[:n])
	}

	return strings.TrimSpace(res), nil
}

func main() {
	// 获取本地ip v4地址
	ip, err := GetExternalIP()
	if err != nil {
		ip = "192.168.1.1" // 默认地址
	}

	fmt.Printf("ip = %s\n", ip)
	// 输出：ip = 192.168.2.21
}
