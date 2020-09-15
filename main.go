//go:generate go run -tags generate gen.go

package main

import (
	"flag"
	"log"
	"net"

	"github.com/eyedeekay/di2prc/lib"
	"github.com/eyedeekay/goSam"
	"github.com/getlantern/go-socks5"
	"github.com/i19/autorestart"
)

var (
	port   = flag.String("sam", "127.0.0.1:7656", "SAM address to set up the web service on.")
	hport  = "127.0.0.1:4444" //flag.String("proxy", "127.0.0.1:4444", "HTTP proxy port")
	pport  = "127.0.0.1:4446"
	ribbon = flag.Bool("ribbon", false, "use a horizontal ribbon instead of a vertical panel")
)

func main() {
	flag.Parse()
	autorestart.Run(worker)
}

func worker() {

	// Create a SOCKS5 server
	sam, err := goSam.NewDefaultClient()
	if err != nil {
		panic(err)
	}

	log.Println("Client Created")

	// create a transport that uses SAM to dial TCP Connections
	conf := &socks5.Config{
		Dial:     sam.DialContext,
		Resolver: sam,
	}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}
	ln, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	//pport = ln.Addr().String()
	go run(ln, server)

	if *ribbon {
		l, v := di2prc.Launch(*port, "127.0.0.1:4444", pport, 800, 400)
		defer v.Destroy()
		defer l.Close()
	} else {
		l, v := di2prc.Launch(*port, "127.0.0.1:4444", pport, 500, 800)
		defer v.Destroy()
		defer l.Close()
	}

}

func run(ln net.Listener, server *socks5.Server) {
	if err := server.Serve(ln); err != nil {
		panic(err)
	}
}
