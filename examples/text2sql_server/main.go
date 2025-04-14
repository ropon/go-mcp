package main

import (
	"fmt"
	"log"

	"github.com/ropon/go-mcp/protocol"
	"github.com/ropon/go-mcp/server"
	"github.com/ropon/go-mcp/transport"
)

func main() {
	s, err := server.NewServer(
		transport.NewStdioServerTransport(),
		server.WithServerInfo(protocol.Implementation{
			Name:    "text2sql_server",
			Version: "1.0.0",
		}),
	)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}
	// register tool
	s.RegisterTool(getText2sqlTool(), text2sqlHandler)
	// start mcp server
	if err := s.Run(); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
