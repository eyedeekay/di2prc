//go:generate go run -tags generate gen.go

package main

import (
	"flag"
	"github.com/eyedeekay/di2prc/lib"
)

var (
	port   = flag.String("sam", "127.0.0.1:7656", "Port to run the web interface on, default is randomly assigned.")
	pport  = flag.String("proxy", "127.0.0.1:4444", "Port to use to proxy requests to i2p-control")
	ribbon = flag.Bool("ribbon", false, "use a horizontal ribbon instead of a vertical panel")
)

func main() {
	flag.Parse()
	if *ribbon {
		l, v := di2prc.Launch(*port, *pport, 800, 400)
		defer v.Destroy()
		defer l.Close()
	} else {
		l, v := di2prc.Launch(*port, *pport, 500, 800)
		defer v.Destroy()
		defer l.Close()
	}

}
