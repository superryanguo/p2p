package main

import (
	"context"
	"flag"
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
	addrP := flag.String("addr", "", "peer TCP address")
	flag.Parse()
	fmt.Printf("The ping addr =%s\n", *addrP)

	alice, err := noise.NewNode()
	check(err)

	// Gracefully release resources for Alice and Bob at the end of the example.

	defer alice.Close()

	// When Bob gets a message from Alice, print it out and respond to Alice with 'Hi Alice!'

	check(alice.Listen())

	// Have Alice send Bob a request with the message 'Hi Bob!'

	res, err := alice.Request(context.TODO(), *addrP, []byte("Hi Alice->Bob!"))
	check(err)

	// Print out the response Bob got from Alice.

	fmt.Printf("Got a message from Bob: '%s'\n", string(res))

}
