GXX=go
BIN=shortener-svr
GXX_SRC=$(wildcard **/*.go)
PROTO=protoc
PROTO_DEF=./pb/shortener.proto
PROTO_TS_OUT=./pb/ts
PROTO_GO_OUT=./pb/go

.PHONY: grpc-go grpc-ts clean

$(BIN): $(GXX_SRC)
	$(GXX) build -o $(BIN) .

grpc-go: $(PROTO_DEF)
	mkdir -p $(PROTO_GO_OUT)
	$(PROTO) --go_out=$(PROTO_GO_OUT) --go_opt=paths=source_relative \
           --go-grpc_out=$(PROTO_GO_OUT) --go-grpc_opt=paths=source_relative \
		       $(PROTO_DEF)

grpc-ts: $(PROTO_DEF)
	mkdir -p $(PROTO_TS_OUT)
	protoc --js_out=import_style=commonjs,binary:$(PROTO_TS_OUT) $(PROTO_DEF)
	protoc --grpc-web_out=import_style=typescript,mode=grpcwebtext:$(PROTO_TS_OUT) $(PROTO_DEF)

clean:
	rm -rf $(BIN)
