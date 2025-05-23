gen :
	 PATH=$$PATH:$$HOME/go/bin \
	protoc --proto_path=proto \
		--go_out=pb --go_opt=paths=source_relative \
		--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		proto/*.proto

clean:
	rm pb/*.go

run:
	go run main.go