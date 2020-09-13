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
];
global.WEBTORRENT_ANNOUNCE = announceList;
