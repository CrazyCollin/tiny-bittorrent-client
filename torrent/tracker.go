package torrent

import (
	"CrazyCollin/personalProjects/tiny-bittorrent-client/peers"
	"github.com/jackpal/bencode-go"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

//
// bencodeTrackerResp
// @Description: 从tracker中返回的信息
//
type bencodeTrackerResp struct {
	Interval int    `bencode:"interval"`
	Peers    string `bencode:"peers"`
}

//
// buildTrackerURL
// @Description:向tracker声明的url发送get请求，建立get url
// @receiver t
// @param peerID
// @param port
// @return string
// @return error
//
func (t *TorrentFile) buildTrackerURL(peerID [20]byte, port uint16) (string, error) {
	base, err := url.Parse(t.Announce)
	if err != nil {
		return "", err
	}
	params := url.Values{
		"info_hash":  []string{string(t.InfoHash[:])},
		"peer_id":    []string{string(peerID[:])},
		"port":       []string{strconv.Itoa(int(port))},
		"uploaded":   []string{"0"},
		"downloaded": []string{"0"},
		"compact":    []string{"1"},
		"left":       []string{strconv.Itoa(t.Length)},
	}
	base.RawQuery = params.Encode()
	return base.String(), nil
}

//
// requestPeers
// @Description: 发出请求，获取peer结果并包装
// @receiver t
// @param peerID
// @param port
// @return []peers.Peer
// @return error
//
func (t *TorrentFile) requestPeers(peerID [20]byte, port uint16) ([]peers.Peer, error) {
	url, err := t.buildTrackerURL(peerID, port)
	if err != nil {
		return nil, err
	}
	client := http.Client{Timeout: 15 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	trackerResp := bencodeTrackerResp{}
	err = bencode.Unmarshal(resp.Body, &trackerResp)
	if err != nil {
		return nil, err
	}
	return peers.Unmarshal([]byte(trackerResp.Peers))
}
