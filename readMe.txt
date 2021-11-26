Test crud grpc project.
Use: PostgreSQL
 
1) compilate .proto: protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/userinfo.proto
2) run server: go run server/server.go
3) run client: go run client/client.go
