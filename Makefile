create:
	protoc --proto_path=proto proto/*.proto --go_out=gen/
	protoc --proto_path=proto proto/*.proto --go-grpc_out=gen/
	protoc -I ./proto \
        --go_out ./gen/proto --go_opt paths=source_relative \
        --go-grpc_out ./gen/proto --go-grpc_opt paths=source_relative \
        --grpc-gateway_out ./gen/proto --grpc-gateway_opt paths=source_relative \
    	proto/deposit.proto
	
clean:
	rm gen/proto/*.go


# untuk "google/api/annotations.proto"
# mkdir -p google/api    
# curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto > google/api/annotations.proto     
# curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto > google/api/http.proto
