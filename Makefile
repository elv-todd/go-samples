
run:
	@#go run .
	for i in `ls -rt *go`; do echo "--"; echo "go run $$i" ; go run $$i; done

flow:
	go1.18 run flow-get.go

sse:
	go1.18 run gin-sse.go
	#go1.18 run server-send-events.go
	
import:
	go mod tidy