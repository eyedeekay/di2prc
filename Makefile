
DIRC_VERSION=0.0.4


build:
	go build --tags=netgo

asset: get fmt gen

run: fmt gen
	go run ./ #-ribbon true

test:
	cd import && go test

fmt: style
	gofmt -w -s *.go */*.go
	fixjsstyle lib/init.js
	cp lib/init.js www/js/init.js
	cp index.html www/index.html

clean:
	rm i2pcontrol.js

i2pcontrol.js:
	wget "https://raw.githubusercontent.com/eyedeekay/I2P-in-Private-Browsing-Mode-Firefox/i2pcontrol/i2pcontrol/i2pcontrol.js"

gen:
	go run --tags=generate gen.go

style:
	sed -i 's|#222|#1F1A24|g' www/css/site.css

get:
	rm -rf v$(DIRC_VERSION).zip www dirc
	wget https://github.com/chr15m/dirc/archive/v$(DIRC_VERSION).zip
	#git clone https://github.com/eyedeekay/dirc -b update
	unzip v$(DIRC_VERSION).zip
	#mv dirc/build www
	mv dirc-$(DIRC_VERSION) www
