postgres:
	docker run --name user-service-postgres12 -p 6001:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=user-service -d postgres:12-alpine

createdb:
	docker exec -it user-service-postgres12 createdb --username=root --owner=root user-service-database

dropdb:
	docker exec -it user-service-postgres12 dropdb user-service-database

migrateup:
	migrate -path migration -database "postgresql://root:user-service@localhost:6001/user-service-database?sslmode=disable" -verbose up

migratedown:
	migrate -path migration -database "postgresql://root:user-service@localhost:6001/user-service-database?sslmode=disable" -verbose down

sqlc:
	sqlc generate

server:
	go run cmd/main.go	

proto:
	protoc --proto_path=internal/grpc/user/proto --go_out=internal/grpc/user/pb --go_opt=paths=source_relative \
    --go-grpc_out=internal/grpc/user/pb --go-grpc_opt=paths=source_relative \
    internal/grpc/user/proto/*.proto

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server
