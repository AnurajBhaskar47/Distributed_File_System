package p2p

import (
	// "bytes"
	"fmt"
	"net"
	"sync"
)

// TCPPeer represents the remote node over a TCP established connecion
type TCPPeer struct {
    // conn is the underlying connection of the peer
    conn net.Conn

    // if we dial and retrieve a conn => outbound == true
    // if we accpet and retrieve a conn => outbound == false
    outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
    return &TCPPeer{
        conn: conn,
        outbound: outbound,
    }
}

type TCPTransportOpts struct {
    ListenAddr    string
    HandshakeFunc HandshakeFunc
    Decoder       Decoder
}

type TCPTransport struct {
    TCPTransportOpts
    listenAddress string
    listener    net.Listener
    shakeHands  HandshakeFunc
    decoder     Decoder

    mu sync.RWMutex
    peers map[net.Addr]Peer
}


func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
    return &TCPTransport{
        TCPTransportOpts: opts,
    }
}

func (t *TCPTransport) ListenAndAccept() error {
    var err error
    
    t.listener, err = net.Listen("tcp", t.listenAddress)
    if err != nil {
        return err
    }

    go t.startAcceptLoop()

    return nil

}

type Temp struct{}

func (t *TCPTransport) startAcceptLoop() {
    for {
        conn, err := t.listener.Accept()
        if err != nil {
            fmt.Printf("TCP accept error")
        }
        go t.handleConn(conn)
    }
}


func (t *TCPTransport) handleConn(conn net.Conn) {
    peer := NewTCPPeer(conn, true)

    if err := t.HandshakeFunc(peer); err != nil {
        conn.Close()
        fmt.Printf("TCP handshake error: %s\n", err)
        return
    }

    // Read Loop
    // msg := &Message{}
    buf := make([]byte, 2000)
    for {
        n, err := conn.Read(buf)
        if err != nil {
            fmt.Printf("TCP error: %s\n", err)
        }
        // if err := t.Decoder.Decode(conn, msg); err != nil {
        //     fmt.Printf("TCP error: %s\n", err)
        //     continue
        // }

        fmt.Printf("message: %+v\n", buf[:n])
    }
}
