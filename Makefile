GOFILES := $(wildcard cmd/server/*.go) $(wildcard internal/**/*.go)
EXE := server.exe

.PHONY: ogen server

ogen:
	cd tools && go generate ./ogen

server: $(GOFILES) ogen
	go build -o $(EXE) ./cmd/server
