
run: p2p.out
	./p2p.out

p2p.out: p2p.go
	go build -o ./p2p.out p2p.go

test:
	go test -v  -race 
