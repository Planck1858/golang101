# see: https://github.com/protocolbuffers/protobuf/releases
deps:
		curl -LO "https://github.com/protocolbuffers/protobuf/releases/download/v26.1/protoc-26.1-linux-x86_64.zip"
		unzip protoc-26.1-linux-x86_64.zip -d "${HOME}/.local"
		rm protoc-26.1-linux-x86_64.zip

		go get \
			google.golang.org/protobuf/cmd/protoc-gen-go \
			google.golang.org/grpc/cmd/protoc-gen-go-grpc \
			github.com/envoyproxy/protoc-gen-validate \
			github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
			github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
			
			# github.com/googleapis/gapic-generator-go

		go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
		go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
		go install github.com/envoyproxy/protoc-gen-validate@latest
		go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
		go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

		# rm -rf gapic-generator-go
		# git clone https://github.com/googleapis/gapic-generator-go.git
		# go install ./gapic-generator-go/cmd/protoc-gen-go_gapic/
		# rm -rf gapic-generator-go

generate:
		protoc \
		--proto_path=./proto/api \
		--proto_path=./proto/github.com/googleapis/googleapis \
		--proto_path=./proto/github.com/validate \
		--proto_path=./proto/github.com/grpc-ecosystem/grpc-gateway \
		--proto_path=./proto/github.com \
		--proto_path=${HOME}/.local/include \
		--go_out=./gen/go --go_opt=paths=source_relative \
		--go-grpc_out=./gen/go --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=allow_delete_body=true:./gen/go --grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=allow_merge=true,allow_delete_body=true,merge_file_name=api:./gen/go \
		--validate_out=lang=go:./gen/go --validate_opt paths=source_relative \
		./proto/api/api.proto

