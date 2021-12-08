package main

import (
	"context"
	"flag"
	"log"
	"net"
	"time"
)

func main() {
	addrP := flag.String("addr", "", "peer TCP address")
	flag.Parse()
	log.Infof("The ping addr =%s\n", *addrP)
	var d net.Dialer
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	conn, err := d.DialContext(ctx, "tcp", *addrP)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	if _, err := conn.Write([]byte("Hello, World from 59 server!")); err != nil {
		log.Fatal(err)
	}
}
