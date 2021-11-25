
run: p2p.out
	./p2p.out

p2p.out: p2p.go log/log.go
	go build -o ./p2p.out p2p.go

test:
	go test -v  -race 
