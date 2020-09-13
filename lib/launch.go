//go:generate go run -tags generate gen.go

package di2prc

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/eyedeekay/sam3/i2pkeys"
	"github.com/webview/webview"
)

func Launch(port string, pport string, width, height int) (ln net.Listener, w webview.WebView) {
	ln = Listen(port, "", "")
	//	defer ln.Close()
	debug := true
	os.Setenv("http_proxy", "http://"+pport)
	os.Setenv("https_proxy", "http://"+pport)
	os.Setenv("ftp_proxy", "http://"+pport)
	os.Setenv("all_proxy", "http://"+pport)
	os.Setenv("HTTP_PROXY", "http://"+pport)
	os.Setenv("HTTPS_PROXY", "http://"+pport)
	os.Setenv("FTP_PROXY", "http://"+pport)
	os.Setenv("ALL_PROXY", "http://"+pport)
	time.Sleep(5 * time.Second)
	w = webview.New(debug)
	//	defer w.Destroy()
	/*localStorage["dirc-wt-config"] = JSON.stringify({rtcConfig: {tracker: {iceServers: [{urls: 'ppnxqa3o6ldzjaurbm4vrbutwsdlmaar5hhamga6jxvmstkeo4uq.b32.i2p' }, {url: 'YOUR-OTHER-SERVER'}, ...] }}});*/
	w.Init(`
localStorage['dirc-wt-config'] = JSON.stringify({
  rtcConfig: {
    tracker: {
      iceServers: [{
        urls: 'ppnxqa3o6ldzjaurbm4vrbutwsdlmaar5hhamga6jxvmstkeo4uq.b32.i2p'
      }],
    }
  }
});
announceList = [
  ['http://yru3sbhbksao6uoaes4n56jtnmqa3k2i5mv67c7lb2x7eqcfp2la.b32.i2p'],
  ['wss://yru3sbhbksao6uoaes4n56jtnmqa3k2i5mv67c7lb2x7eqcfp2la.b32.i2p'],
  ['wss://tracker.btorrent.xyz'],
  ['wss://tracker.openwebtorrent.com']
]
global.WEBTORRENT_ANNOUNCE = announceList
	`)
	w.SetTitle("di2prc")
	w.SetSize(width, height, webview.HintFixed)
	w.Navigate(fmt.Sprintf("http://%s", ln.Addr().(i2pkeys.I2PAddr).Base32()))
	w.Run()
	return
}
