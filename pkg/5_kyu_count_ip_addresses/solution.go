package count_ip_addresses

import (
	"strconv"
	"strings"
)

func solution(start, end string) int {
	return IpToInt(end) - IpToInt(start)
}

func IpToInt(ip string) int {
	parts := strings.Split(ip, ".")

	p0, _ := strconv.Atoi(parts[0])
	p1, _ := strconv.Atoi(parts[1])
	p2, _ := strconv.Atoi(parts[2])
	p3, _ := strconv.Atoi(parts[3])

	return p0<<24 + p1<<16 + p2<<8 + p3
}
