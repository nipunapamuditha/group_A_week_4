package main

import (
	"fmt"
	"net"
	"time"
)

func CheckHostAvailability(host string) string {
	result := ""

	conn, err := net.DialTimeout("tcp", host+":80", 2*time.Second)
	if err != nil {
		result += fmt.Sprintf("%s is not available\n", host)
	} else {
		result += fmt.Sprintf("%s is available\n", host)
		conn.Close()
	}

	return result
}
