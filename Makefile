
run:
	go1.18 run compare.go

build:
	go1.18 build -o base58 base58.go

all:
	for i in `ls -rt *go`; do echo "--"; echo "go run $$i" ; go run $$i; done

solana:
	go1.18 run solana.go

flow:
	go1.18 run flow-get.go

sse:
	go1.18 run gin-sse.go
	#go1.18 run server-send-events.go
	
import:
	go mod tidy