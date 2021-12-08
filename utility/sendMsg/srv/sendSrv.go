package main

import (
	"fmt"

	"github.com/perlin-network/noise"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// This example demonstrates how to send/handle RPC requests across peers, how to listen for incoming
// peers, how to check if a message received is a request or not, how to reply to a RPC request, and
// how to cleanup node instances after you are done using them.
func main() {
	// Let there be nodes Alice and Bob.

	//bob, err := noise.NewNode()
	bob, err := noise.NewNode(noise.WithNodeBindPort(50000))
	check(err)

	// Gracefully release resources for Alice and Bob at the end of the example.

	defer bob.Close()

	// When Bob gets a message from Alice, print it out and respond to Alice with 'Hi Alice!'

	bob.Handle(func(ctx noise.HandlerContext) error {
		if !ctx.IsRequest() {
			return nil
		}

		fmt.Printf("Got a message from Alice: '%s'\n", string(ctx.Data()))

		return ctx.Send([]byte("Hi this is 26 server!"))
	})

	// Have Alice and Bob start listening for new peers.

	check(bob.Listen())

	fmt.Printf("beign the bob servering@%v\n", bob.Addr())
	for {

	}
}
