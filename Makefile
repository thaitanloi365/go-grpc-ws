build-service:
	protoc --go_out=plugins=grpc:. ws/message.proto
