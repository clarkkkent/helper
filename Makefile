docker-start:
	docker-compose up --build

protobuf:
	protoc user-service/proto/user/user.proto --go_out=plugins=grpc:.


