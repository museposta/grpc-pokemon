

go mod init github.com/museposta/grpc-pokemon
go get -u google.golang.org/grpc
go get -u google.golang.org/protobuf/protoc-gen-go

mkdir pokemon

protoc --go-grpc_out=. pokemon/pokemon.proto 
protoc --go_out=. pokemon/pokemon.proto

go get go.mongodb.org/mongo-driver/mongo

go get github.com/joho/godotenv


mkdir server/server.go

mkdir client/client.go

go run server/server.go

