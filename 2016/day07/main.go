package main

import (
	"fmt"

	"advent/lib/util"
)

func main() {
	tls := 0
	ssl := 0

	scanner := util.InputScanner()
	for scanner.Scan() {
		ip := scanner.Text()

		if IsTLSCompliant(ip) {
			tls++
		}

		if IsSSLCompliant(ip) {
			ssl++
		}
	}

	fmt.Println("TLS-complient IPs:", tls)
	fmt.Println("SSL-complient IPs:", ssl)
}
