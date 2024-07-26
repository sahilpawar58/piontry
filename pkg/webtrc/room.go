package webrtc

import (
	"time"
	"sync"
	"log"
	"github.com/gofiber/websocket"
	"github.com/pion/webrtc/v3"
)

func RoomConn(c *websocket.Conn, p *Peers){
	var config webrtc.Configuration

	peerConnection,err := webrtc.NewPeerConnection(config)
	if err != nil {
		log.Print(err)
		return
	}
	newPeer := PeerConnectionState{
		PeerConnection: peerChatConn,
		websocket: %ThreadSafeWriter{},
		Conn: c,
		Mutex: sync.Mutex{},
	}
}