
run: p2p.out
	./p2p.out

p2p.out: p2p.go log/log.go
	go build -gcflags "-N -l" -o ./p2p.out p2p.go

test:
	go test -v  -race 
