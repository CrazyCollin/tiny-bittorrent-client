package peers

import (
	"encoding/binary"
	"fmt"
	"net"
	"strconv"
)

type Peer struct {
	IP   net.IP
	Port uint16
}

func (p *Peer) String() string {
	return net.JoinHostPort(p.IP.String(), strconv.Itoa(int(p.Port)))
}

//
// Unmarshal
// @Description: 解析各个peer的ip/port信息
// @param peerBytes
// @return []Peer
// @return error
//
func Unmarshal(peerBytes []byte) ([]Peer, error) {
	const peerSize = 6
	numPeers := len(peerBytes) / peerSize
	if len(peerBytes)%peerSize != 0 {
		return nil, fmt.Errorf("peers nums error")
	}
	peers := make([]Peer, numPeers)
	for i := 0; i < numPeers; i++ {
		offset := i * peerSize
		peers[i].IP = peerBytes[offset : offset+4]
		peers[i].Port = binary.BigEndian.Uint16(peerBytes[offset+4 : offset+6])
	}
	return peers, nil
}
