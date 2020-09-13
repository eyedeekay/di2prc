
DIRC_VERSION=0.0.4


build: fmt get gen
	go build --tags=netgo

run: fmt gen
	go run ./ #-ribbon true

test:
	cd import && go test

fmt: style
	gofmt -w -s *.go */*.go

clean:
	rm i2pcontrol.js

i2pcontrol.js:
	wget "https://raw.githubusercontent.com/eyedeekay/I2P-in-Private-Browsing-Mode-Firefox/i2pcontrol/i2pcontrol/i2pcontrol.js"

gen:
	go run --tags=generate gen.go

style:
	sed -i 's|#222|#1F1A24|g' www/css/site.css

get:
	rm -rf v$(DIRC_VERSION).zip www
	wget https://github.com/chr15m/dirc/archive/v$(DIRC_VERSION).zip
	unzip v$(DIRC_VERSION).zip
	mv dirc-$(DIRC_VERSION) www
