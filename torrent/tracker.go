package torrent

import "CrazyCollin/personalProjects/tiny-bittorrent-client/peers"

//
// bencodeTrackerResp
// @Description: 从tracker中返回的信息
//
type bencodeTrackerResp struct {
	Interval int    `bencode:"interval"`
	Peers    string `bencode:"peers"`
}

func (t *TorrentFile) buildTrackerURL(peerID [20]byte, port uint16) (string, error) {

	return "", nil
}

func (t *TorrentFile) requestPeers(peerID [20]byte, port uint16) ([]peers.Peer, error) {
	return nil, nil
}
