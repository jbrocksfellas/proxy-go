package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/things-go/go-socks5"
)

func main() {
	server := socks5.NewServer(
		socks5.WithLogger(socks5.NewLogger(log.New(os.Stdout, "socks5: ", log.LstdFlags))),
		socks5.WithAssociateHandle(func(ctx context.Context, writer io.Writer, request *socks5.Request) error {
			fmt.Print("no of bytes", request.Request.Bytes())
			return nil
		}),
	)

	if err := server.ListenAndServe("tcp", "localhost:8000"); err != nil {
		panic(err)
	}

}
