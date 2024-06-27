sqlc:
	sqlc generate

server:
	go run cmd/main.go	

proto:
	for service in user authentication book order; do \
		protoc --proto_path=internal/grpc/proto/$$service:internal/grpc/proto/google/type --go_out=internal/grpc/gen/$$service --go_opt=paths=source_relative \
		--go-grpc_out=internal/grpc/gen/$$service --go-grpc_opt=paths=source_relative \
		internal/grpc/proto/$$service/*.proto; \
	done

.PHONY: sqlc server proto
