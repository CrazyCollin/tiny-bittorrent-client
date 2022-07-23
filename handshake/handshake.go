package handshake

type HandShake struct {
	Pstr     string
	InfoHash [20]byte
	PeerID   [20]byte
}

func (h *HandShake) Serialize() []byte {
	return nil
}
