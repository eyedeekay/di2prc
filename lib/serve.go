package di2prc

import (
	"bytes"
	//	"fmt"
	"github.com/eyedeekay/samsocks/sammy"
	"github.com/justinas/nosurf"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
	//	"crypto/tls"
	"time"
)

const timeout = time.Second * 5

var (
	connIDs  = make(chan uint64)
	connDone = make(chan uint64)
)

type fileServer struct {
	http.Server
	*fs
}

func (serv fileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Path) < 2 {
		r.URL.Path = "/index.html"
	}
	fi, err := serv.fs.Open(r.URL.Path)
	if err != nil {
		fi, err = serv.fs.Open("/index.html")
		if err != nil {
			return
		}
	}
	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, fi)
	if err != nil {
		return
	}
	fi.Close()
	if strings.Contains(r.URL.Path, ".css") {
		w.Header().Add("Content-Type", "text/css")
	}
	if strings.Contains(r.URL.Path, ".js") {
		w.Header().Add("Content-Type", "text/javascript")
	}
	if strings.Contains(r.URL.Path, ".svg") {
		w.Header().Add("Content-Type", "image/svg+xml")
	}
	w.Write(buf.Bytes())
}

func FileServer(files *fs) *fileServer {
	var fis fileServer
	fis.fs = files
	return &fis
}

var Options_DChat_Short = []string{"inbound.length=1", "outbound.length=1", "inbound.lengthVariance=0", "outbound.lengthVariance=0", "inbound.quantity=3", "outbound.quantity=3", "inbound.backupQuantity=2", "outbound.backupQuantity=2", "i2cp.closeOnIdle=false", "i2cp.reduceOnIdle=false", "i2cp.leaseSetEncType=4,0"}

func Listen(yoursam, certfile, keyfile string) net.Listener {
	opts := make(map[string]string)
	opts[`sam`] = yoursam
	opts[`servertun`] = "dirc2p" + sammy.RandStringBytes()
	opts[`keypath`] = opts[`servertun`]+".i2pkeys"
	ln, err := sammy.Sammy(opts)
	if err != nil {
		log.Fatal(err)
	}
	/*srv := http.Server{}
	srv.Addr = ln.Addr().(i2pkeys.I2PAddr).Base32()
	if srv.TLSConfig == nil {
		srv.TLSConfig = &tls.Config{
			ServerName: ln.Addr().(i2pkeys.I2PAddr).Base32(),
		}
	}

	if srv.TLSConfig.NextProtos == nil {
		srv.TLSConfig.NextProtos = []string{"http/1.1"}
	}

	//	var err error
	srv.TLSConfig.Certificates = make([]tls.Certificate, 1)
	srv.TLSConfig.Certificates[0], err = tls.LoadX509KeyPair(certfile, keyfile)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(fmt.Sprintf("http://%s", ln.Addr().(i2pkeys.I2PAddr).Base32()))*/
	//	go http.ServeTLS(ln, nosurf.New(FileServer(FS)), certfile, keyfile) //, )
	go http.Serve(ln, nosurf.New(FileServer(FS))) //, )
	return ln
}
