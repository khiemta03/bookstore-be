postgres:
	docker run --name order-service-postgres12 -p 6004:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=order-service -d postgres:12-alpine

createdb:
	docker exec -it order-service-postgres12 createdb --username=root --owner=root order-service-database

dropdb:
	docker exec -it order-service-postgres12 dropdb order-service-database

migrateup:
	migrate -path migration -database "postgresql://root:order-service@localhost:6004/order-service-database?sslmode=disable" -verbose up

migratedown:
	migrate -path migration -database "postgresql://root:order-service@localhost:6004/order-service-database?sslmode=disable" -verbose down

sqlc:
	sqlc generate

server:
	go run cmd/main.go	

proto:
	protoc --proto_path=internal/grpc/proto/book --go_out=internal/grpc/gen/book --go_opt=paths=source_relative \
    --go-grpc_out=internal/grpc/gen/book --go-grpc_opt=paths=source_relative \
    internal/grpc/proto/book/*.proto 

	protoc --proto_path=internal/grpc/proto/order --go_out=internal/grpc/gen/order --go_opt=paths=source_relative \
    --go-grpc_out=internal/grpc/gen/order --go-grpc_opt=paths=source_relative \
	--experimental_allow_proto3_optional \
    internal/grpc/proto/order/*.proto

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server
