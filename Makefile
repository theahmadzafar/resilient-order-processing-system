
gen-rpc:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. --go-grpc_opt=paths=source_relative     services/inventry-service/pkg/api/main.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. --go-grpc_opt=paths=source_relative     services/order-service/pkg/api/main.proto


tests:
	go test ./tests -v

lint:
	golangci-lint run -v
docker-up:
	docker compose build --no-cache --progress=plain
	docker compose up -d
	

docker-dryrun:
	docker compose \
	-f compose.yaml \
	--progress=plain \
	--project-directory . \
	up --detach \
	--dry-run \

docker-down: 
	docker compose \
	-f compose.yaml \
	--progress=plain \
	down 

docker-stats: 
	docker compose \
	-f compose.yaml \
	--progress=plain \
	stats 
