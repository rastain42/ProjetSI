package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var (
    listen = flag.Bool("l", false, "Listen")
    
    host   = flag.String("h", "192.168.64.1", "Host")

    port = flag.Int("p", 0, "Port")
)

func startServer() {

    addr := fmt.Sprintf("%s:%d", *host, *port)
    listener, err := net.Listen("tcp", addr)

    if err != nil {

        panic(err)
    }

    log.Printf("Listening for connections on %s", listener.Addr().String())

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Printf("Error accepting connection from client: %s", err)
        } else {

            fmt.Println("New Connections from", conn.RemoteAddr())

            go processClient(conn)
            go processSender(conn)

        }
    }
}

func processClient(conn net.Conn) {
    _, err := io.Copy(os.Stdout, conn)
    if err != nil {
        fmt.Println(err)
    }
    conn.Close()
}

func processSender(conn net.Conn) {
    _, err := io.Copy(conn, os.Stdin)
    if err != nil {
        fmt.Println(err)
    }
    conn.Close()
}

func main() {
    flag.Parse()
    if *listen {
        startServer()
        return
    }
}