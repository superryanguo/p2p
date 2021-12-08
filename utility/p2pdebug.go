package main

import (
	"context"
	"flag"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"runtime/trace"

	"github.com/superryanguo/p2p/log"

	"github.com/perlin-network/noise"
	"github.com/perlin-network/noise/kademlia"
)

func main() {
	go func() {
		http.ListenAndServe(":12003", nil)
	}()
	f, err := os.Create("./trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	addrP := flag.String("addr", "", "peer TCP address")
	flag.Parse()
	log.Debugf("The ping addr =%s\n", *addrP)
	node, err := noise.NewNode(noise.WithNodeBindPort(12005))
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

	c := make(chan os.Signal, 1)
	done := make(chan struct{}, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		sig := <-c
		log.Infof("Got singal: %v\n", sig)
		done <- struct{}{}
	}()

	log.Infof("p2p node serving@%s\n", node.Addr())
	//ad := "localhost:43633" ad := "127.0.0.1:43633"

	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	if *addrP != "" {
		if _, err := node.Ping(ctx, *addrP); err != nil {
			log.Fatal(err.Error())
		}
	}

	for {
		if len(kn.Discover()) != 0 {
			log.Infof("Node discovered %d peer(s).\n", len(kn.Discover()))
		}

		select {
		case <-done:
			log.Info("Exiting by ctrl+c and so on...")
			return
		default:
			time.Sleep(5 * time.Second)
		}
	}

}
