# Golang Distributed File System

This project is a **high-performance**, **distributed file system** implemented in Go, designed to provide scalable, fault-tolerant, and efficient storage solutions for modern applications.
It implements the **transport layer** of a **peer-to-peer** (P2P) distributed file system using Go. Providing core functionalities for communication between nodes over TCP, message encoding/decoding, handshaking between peers, and the ability to extend to other protocols if needed.

#### Features

- Scalability: Easily scale out by adding more nodes to handle increasing amounts of data and traffic.

- Fault Tolerance: Automatic data replication and recovery mechanisms ensure high availability and durability.
- Performance: Optimized for high throughput and low latency, making it suitable for a variety of workloads.
- Simple API: Provides a straightforward API for file operations, making it easy to integrate into your applications.

### Getting Started
#### Prerequisites

    Go 1.18 or higher
    A Unix-like operating system (Linux or macOS) for development
    Basic understanding of distributed systems concepts

### Installation

##### 1. Clone the repository:

```bash

git clone --depth 1 https://github.com/AnurajBhaskar47/Distributed_File_System.git
cd Distributed_File_System
```

##### 2. Build the project:

```bash
go build -o bin/fs
```

###### Run the system:

Start the master node with workers:

```bash
./bin/fs
```

###### For testing
```bash
go test -v ./p2p/tcp_transport_test.go
```

---

##### Using Makefile for build and run
```bash
make run
```
<video src="./misc/file_build.mp4" width="640" height="400" controls></video>

### Project Structure
###### This is an overall view of the file structure of the project:
```
.
├── bin
│   └── fs
├── crypto.go
├── crypto_test.go
├── go.mod
├── go.sum
├── main.go
├── Makefile
├── misc
├── p2p
│   ├── encoding.go
│   ├── handshake.go
│   ├── message.go
│   ├── tcp_transport.go
│   ├── tcp_transport_test.go
│   └── transport.go
├── README.md
├── server.go
├── store.go
└── store_test.go
```
The following Golang files are included in the transport layer implementation:

- `encoding.go`: Handles encoding and decoding of messages.
- `handshake.go`: Manages the handshake process between peers.
- `message.go`: Defines the structure of messages exchanged between peers.
- `tcp_transport.go`: Implements TCP-based peer communication.
- `tcp_transport_test.go`: Contains unit tests for the TCP transport layer.
- `transport.go`: Defines interfaces for peer and transport communication.

---

### `encoding.go`

This file is responsible for encoding and decoding messages exchanged between peers in the network.

- **Decoder Interface**: A common interface for decoders with a `Decode` method to decode messages from the network.
- **GOBDecoder**: Uses Go’s `gob` package for binary decoding of messages.
- **DefaultDecoder**: A custom implementation that reads stream flags and payloads directly from the network, handling stream-based communication.

---

### `handshake.go`

Defines the handshake function that runs when two peers establish a connection.

- **HandshakeFunc**: A function type for handshake procedures. It accepts a `Peer` and returns an error if the handshake fails.
- **NOPHandshakeFunc**: A no-operation handshake function that always succeeds, acting as the default handshake behavior.

---

### `message.go`

This file defines the structure of the messages (`RPC`) exchanged between peers.

- **Constants**: 
  - `IncomingMessage`: Flag to mark a standard message.
  - `IncomingStream`: Flag to mark a stream-based message.
- **RPC struct**: The message structure used for communication between peers. It contains:
  - `From`: The address of the sender.
  - `Payload`: The data being transmitted.
  - `Stream`: A flag indicating whether the message is part of a stream.

---

### `tcp_transport.go`

Implements the TCP transport layer for peer-to-peer communication. It manages TCP connections, handshakes, message transmission, and stream handling.

- **TCPPeer**: Represents a remote peer using a TCP connection. It contains:
  - The `net.Conn` representing the peer’s connection.
  - An `outbound` flag to indicate if the connection was initiated by this node.
  - A wait group (`wg`) for synchronizing stream handling.
- **TCPTransport**: Manages the lifecycle of TCP connections, from accepting incoming connections to reading and decoding messages. It is responsible for:
  - Listening for new connections.
  - Handshaking with peers.
  - Reading messages and handling stream-based communication.

---

### `tcp_transport_test.go`

This file contains unit tests for the `TCPTransport` implementation using the Go `testing` package and the `testify` library.

- **TestTCPTransport**: Tests that the `TCPTransport` object initializes correctly and listens on the proper address. It also verifies the acceptance of incoming connections and proper functioning of the transport.

---

### `transport.go`

Defines the `Peer` and `Transport` interfaces, abstracting peer communication and transport mechanisms.

- **Peer Interface**: Represents a remote node in the network. It provides methods for:
  - `Send([]byte)`: Send data to the peer.
  - `CloseStream()`: Close the current data stream.
  - Accessing the underlying `net.Conn`.
- **Transport Interface**: Abstracts the transport mechanism, allowing different transport types (TCP, UDP, websockets) to be used. It provides methods for:
  - `Addr()`: Get the transport's address.
  - `Dial(string)`: Initiate a connection to another peer.
  - `ListenAndAccept()`: Listen for and accept incoming peer connections.
  - `Consume()`: Get a channel to consume incoming `RPC` messages.
  - `Close()`: Close the transport.

---

<!-- ![[./misc/file_structure.png]] -->
<!-- #### 3. Configuration

Edit the config.yaml file to configure your master and worker nodes. Refer to the Configuration Guide for detailed instructions.

##### Usage

###### Initialize the file system:

```bash
./gdfs init
```

###### Upload a file:

```bash
./gdfs put /path/to/local/file /path/in/distributed/system
```

###### Download a file:

```bash
./gdfs get /path/in/distributed/system /path/to/local/destination
```

###### List files:

```bash
    ./gdfs ls /path/in/distributed/system
``` -->

<!-- ### Documentation

For more detailed documentation, please refer to the following resources:

    [User Guide](https://www.google.com)
    API Documentation
    Configuration Guide

Contributing

We welcome contributions from the community! If you’d like to contribute to the project, please refer to our Contributing Guidelines. -->

### License

This project is licensed under the MIT License.

**Contact**
For any questions or issues, please open an issue on GitHub or contact me at bhaskar25903@gmail.com.
