package main

import (
	"flag"
	"github.com/platform9/cnxmd/pkg/cnxmd"
	"os"
)

func main() {
	var bindAddr string
	var lstnPort int
	var destFqdn string
	var destPort int
	var destHostname string
	flag.StringVar(&bindAddr, "bind", "0.0.0.0",
		"bind address")
	flag.IntVar(&lstnPort, "port", 0,
		"listening port (default: dynamic port)")
	flag.StringVar(&destFqdn, "dest", "",
		"fqdn of CNXMD-compliant destination proxy server (required)")
	flag.IntVar(&destPort, "dport",
		443, "TCP port of proxy server")
	flag.StringVar(&destHostname, "dhost", "",
		"name of destination host (default: same as destination proxy server)")
	flag.Parse()
	if destFqdn == "" {
		flag.Usage()
		os.Exit(1)
	}
	if destHostname == "" {
		destHostname = destFqdn
	}
	kv := map[string]string {"host": destHostname}
	cnxmd.ServeClientProxy(bindAddr, lstnPort, destFqdn, destPort, kv)
}
