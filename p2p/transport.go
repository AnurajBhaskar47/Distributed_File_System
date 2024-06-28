package p2p

// Peer is an interface -> represents the remote node
type Peer interface {}

// Transport handles comm. b/w nodes in the network
// form (TCP, UDP, websockets, ...)
type Transport interface {
    ListenAndAccept() error
}
