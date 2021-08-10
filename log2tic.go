package main

import (
	"fmt"
	"net"
	"bufio"
	"regexp"
	"github.com/go-resty/resty/v2"
)

func eval(c net.Conn) {
	cb := bufio.NewReader(c)
	for {
		str, err := cb.ReadString('\n')
		if err != nil { return }
		re := regexp.MustCompile(`.*assigned [0-9]+\.[0-9]+(\.[0-9]+\.[0-9]+).*`)
		ip := re.ReplaceAllString(str, "$1%20%20%20%20%20%20%20%20")

		re = regexp.MustCompile(`\.[0-9]+\.[0-9]+%20.*`)
		if re.Match([]byte(ip)) {
			q := fmt.Sprintf("text=%s", ip)

			client := resty.New()
			r , _ := client.R().EnableTrace().SetQueryString(q).
				Get("http://mozajk:10001/")
			fmt.Println(r)
		}
	}
}


func main() {
	ln, _ := net.Listen("tcp", ":10002")

	for {
		c, _ := ln.Accept()
		go eval(c)
	}
}
