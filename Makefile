pb:
		@echo "pb Start"
		cd encoding && make pb

mkseg:
	-mkdir seg

mkout:
		-mkdir out

out/%.qr.png:seg/%.seg.libv2ray.pb
		./V2RayConfigureFileUtil -t QR -i $^ -o $@

seg: mkseg
	./V2RayConfigureFileUtil -t seg -i inp.json -o seg/

srcfiles := $(shell echo seg/*.seg.libv2ray.pb)
destfiles := $(patsubst seg/%.seg.libv2ray.pb,out/%.qr.png,$(srcfiles))

dispo:
	@echo "$(srcfiles)"

convqr: mkout dispo $(destfiles)
	@echo "Done"


genqr: clean seg
		$(MAKE) convqr
		@echo "Done"

clean:
		-rm -R out seg
