gen:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:proto
clean:
	rm proto/*.go

