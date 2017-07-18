pbSource = LibV2RayPackedConfig.proto QRCode.proto

#%.pb.go:%.proto
#	protoc --go_out=. $^

pb: clean
	protoc --go_out=. $(pbSource)
	@echo "Done"

clean:
	-rm *.pb.go
