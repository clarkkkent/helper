docker-start:
	docker-compose up --build

protobuf:
	protoc user-service/proto/user/user.proto --go_out=plugins=grpc:.

create-user:
	docker exec -it helper_user-service_1 bash



