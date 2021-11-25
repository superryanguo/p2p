package main

import (
	"context"
	"flag"
	"time"

	"github.com/superryanguo/p2p/log"

	"github.com/perlin-network/noise"
	"github.com/perlin-network/noise/kademlia"
)

func main() {
	addrP := flag.String("addr", "", "peer TCP address")
	flag.Parse()
	log.Debugf("The ping addr =%s\n", *addrP)
	node, err := noise.NewNode()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer node.Close()

	node.Handle(func(ctx noise.HandlerContext) error {
		log.Debugf("Got a message from others: '%s'\n", string(ctx.Data()))
		return nil
	})
	kn := kademlia.New()

	node.Bind(kn.Protocol())

	if err := node.Listen(); err != nil {
		log.Fatal(err.Error())
	}

	go log.LogServer(":8097")
	defer log.Sync()

	log.Infof("p2p node serving@%s\n", node.Addr())
	//ad := "localhost:43633" ad := "127.0.0.1:43633"

	if *addrP != "" {
		if _, err := node.Ping(context.TODO(), *addrP); err != nil {
			log.Fatal(err.Error())
		}
	}

	for {
		time.Sleep(5 * time.Second)
		if len(kn.Discover()) != 0 {
			log.Infof("Node discovered %d peer(s).\n", len(kn.Discover()))
		}
	}

}
