pbTarget = LibV2RayPackedConfig.pb.go QRCode.pb.go

%.pb.go:%.proto
	protoc --go_out=. $<

pb: $(pbTarget)
	@echo "Done"
