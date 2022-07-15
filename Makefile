
run:
	@#go run .
	for i in `ls -rt *go`; do echo "--"; echo "go run $$i" ; go run $$i; done

import:
	go mod tidy