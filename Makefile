goose-install:
	go get -u github.com/pressly/goose/v3/cmd/goose 

gen-rpc:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. --go-grpc_opt=paths=source_relative     services/inventry-service/pkg/api/main.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. --go-grpc_opt=paths=source_relative     services/order-service/pkg/api/main.proto

gen-proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative services/order-service/pkg/inventry/main.proto

proto-history:
	protoc --go_out=. --go_opt=paths=source_relative \
				--go-grpc_opt=require_unimplemented_servers=false \
				--go-grpc_out=. --go-grpc_opt=paths=source_relative \
				./pkg/history/main.proto


tests:
	go test ./tests -v

lint:
	golangci-lint run -v
