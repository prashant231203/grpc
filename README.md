# gRPC Demo Project

This project demonstrates the implementation of different types of gRPC (Google Remote Procedure Call) communication patterns in Go. gRPC is a modern, open-source remote procedure call framework that can run anywhere.

## Overview

The project implements four different types of gRPC communication:

1. **Unary RPC**: Simple request/response pattern
   - Client sends a single request
   - Server responds with a single response
   - Like traditional HTTP REST

2. **Server Streaming**: Server streams multiple responses
   - Client sends one request
   - Server responds with a stream of messages
   - Used for scenarios like real-time updates

3. **Client Streaming**: Client streams multiple requests
   - Client sends multiple messages
   - Server responds with a single response
   - Useful for uploading data

4. **Bi-directional Streaming**: Both sides stream
   - Client and server send multiple messages
   - Full-duplex communication
   - Perfect for chat applications or real-time bidirectional data exchange

## Project Structure 

├── client/
│ ├── main.go # Client entry point and connection setup
│ ├── server_stream.go # Handles receiving server streams
│ ├── client_stream.go # Manages client-side streaming
│ └── bi_stream.go # Implements bidirectional streaming
├── Server/
│ ├── main.go # Server initialization and setup
│ ├── server_stream.go # Server streaming implementation
│ ├── client_stream.go # Handles client streams
│ └── bi_stream.go # Bidirectional stream handling
├── proto/
│ ├── greet.proto # Protocol buffer service definitions
│ ├── greet.pb.go # Generated protocol buffer code
│ └── greet_grpc.pb.go # Generated gRPC code
├── go.mod # Go module definition
├── go.sum # Go module checksums
└── README.md # Project documentation

## Prerequisites

Before running this project, ensure you have:

1. **Go** (version 1.22 or later)
   ```bash
   go version
   ```

2. **Protocol Buffers Compiler**
   ```bash
   # For Ubuntu/Debian
   apt install -y protobuf-compiler
   # For macOS
   brew install protobuf
   ```

3. **Go Plugins for Protocol Buffers**
   ```bash
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/prashant231203/grpc.git
   cd grpc
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   go mod vendor
   ```

3. Generate gRPC code:
   ```bash
   protoc --go_out=. --go-grpc_out=. proto/greet.proto
   ```

## Running the Application

1. Start the server:
   ```bash
   cd Server
   go run *.go
   # Server will start on port 5001
   ```

2. In a new terminal, run the client:
   ```bash
   cd client
   go run *.go
   ```

## Implementation Details

### Protocol Buffer Definition

```protobuf
service GreetService {
    // Unary RPC
    rpc SayHello (NoParam) returns (HelloResponse);
    
    // Server Streaming
    rpc ServerStreaming (NameList) returns (stream HelloResponse);
    
    // Client Streaming
    rpc ClientStreaming (stream HelloRequest) returns (MessageList);
    
    // Bi-directional Streaming
    rpc BiDirectionalStreaming (stream HelloRequest) returns (stream HelloResponse);
}
```

### Key Features

1. **Connection Management**
   - Secure connection handling
   - Proper error handling
   - Connection cleanup

2. **Stream Processing**
   - Efficient stream handling
   - Proper error handling for streams
   - Graceful stream closure

3. **Concurrent Processing**
   - Goroutines for stream handling
   - Channel usage for synchronization
   - Context management

### Error Handling

The application implements comprehensive error handling for:
- Connection failures
- Stream errors
- Context cancellations
- Invalid messages
- Timeouts

## Testing

To test different streaming patterns:

1. **Server Streaming**
   ```go
   // Uncomment in client/main.go
   callSayHelloServerStreaming(client, names)
   ```

2. **Client Streaming**
   ```go
   // Uncomment in client/main.go
   callSayHelloClientStreaming(client, names)
   ```

3. **Bi-directional Streaming**
   ```go
   // Uncomment in client/main.go
   callHelloBiDirectionalStreaming(client, names)
   ```

## Common Issues and Solutions

1. **Connection Refused**
   - Ensure server is running
   - Check port availability
   - Verify firewall settings

2. **Protocol Buffer Mismatch**
   - Regenerate protocol buffers
   - Ensure client and server use same proto definition

3. **Stream Errors**
   - Check network stability
   - Verify proper stream closure
   - Monitor for timeouts

## Contributing

1. Fork the repository
2. Create your feature branch
   ```bash
   git checkout -b feature/AmazingFeature
   ```
3. Commit your changes
4. Push to the branch
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- gRPC team for the excellent framework
- Protocol Buffers team
- Go team for the amazing language


Project Link: [https://github.com/prashant231203/grpc](https://github.com/prashant231203/grpc) 