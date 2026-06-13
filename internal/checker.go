package internal

import (
	"fmt"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

type Ip struct {
	Address string `json:"address"`
	IsBusy  bool   `json:"busy"`
}

func isBusy(ip *Ip, wg *sync.WaitGroup) {
	defer wg.Done()

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("ping", "-n", "1", "-w", "500", ip.Address)
	} else {
		cmd = exec.Command("ping", "-c", "1", "-w", "1", ip.Address)
	}

	err := cmd.Run()
	ip.IsBusy = (err == nil)
}

func VerifyIps(ips ...Ip) []Ip {
	var wg sync.WaitGroup
	for i, _ := range ips {
		wg.Add(1)
		go isBusy(&ips[i], &wg)
	}

	wg.Wait()
	return ips
}

func CreateListIps(ipPattern string) []Ip {
	var lista []Ip
	ipParts := strings.Split(ipPattern, ".")

	if len(ipParts) != 4 {
		return lista
	}

	primeiroBloco, err := strconv.Atoi(ipParts[0])

	if err != nil || primeiroBloco == 127 {
		return lista
	}

	if ipParts[2] == "X" || ipParts[2] == "x" {
		subnetStart := fmt.Sprintf("%s.%s.", ipParts[0], ipParts[1])
		subnetEnd := fmt.Sprintf(".%s", ipParts[3])
		for i := 1; i <= 254; i++ {
			var tempIp Ip = Ip{Address: fmt.Sprintf("%s%d%s", subnetStart, i, subnetEnd), IsBusy: false}
			lista = append(lista, tempIp)
		}
	}

	if ipParts[3] == "X" || ipParts[3] == "x" {
		subnet := fmt.Sprintf("%s.%s.%s.", ipParts[0], ipParts[1], ipParts[2])
		for i := 1; i <= 254; i++ {
			var tempIp Ip = Ip{Address: fmt.Sprintf("%s%d", subnet, i), IsBusy: false}
			lista = append(lista, tempIp)
		}
	}

	if (ipParts[2] == "X" || ipParts[2] == "x") && (ipParts[3] == "X" || ipParts[3] == "x") {
		subnet := fmt.Sprintf("%s.%s.", ipParts[0], ipParts[1])
		for i := 1; i <= 254; i++ {
			subnet := fmt.Sprintf("%s%d.", subnet, i)
			for j := 1; j <= 254; j++ {
				var tempIp Ip = Ip{Address: fmt.Sprintf("%s%d", subnet, j), IsBusy: false}
				lista = append(lista, tempIp)
			}
		}
	}

	return lista
}

func Check(ip string) []Ip {
	lista := CreateListIps(ip)
	return VerifyIps(lista...)
}
