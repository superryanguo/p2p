package main

import (
	"fmt"
	"time"
	//TODO: use the zap logger
	"github.com/perlin-network/noise"
	"github.com/perlin-network/noise/kademlia"
)

func main() {
	//TODO: add a flag to pass the client here
	node, err := noise.NewNode()
	if err != nil {
		panic(err)
	}
	defer node.Close()

	node.Handle(func(ctx noise.HandlerContext) error {
		fmt.Printf("Got a message from others: '%s'\n", string(ctx.Data()))
		return nil
	})
	kn := kademlia.New()

	node.Bind(kn.Protocol())

	if err := node.Listen(); err != nil {
		panic(err)
	}

	fmt.Printf("p2p node serving@%s\n", node.Addr())
	for {
		time.Sleep(5 * time.Second)
		fmt.Printf("Node discovered %d peer(s).\n", len(kn.Discover()))
	}

}
