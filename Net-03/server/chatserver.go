package main

import (
	"bufio"
	"io"
	"log"
	"net"
)

//ChatServer class
type ChatServer struct {
	Protocol string
	Port     string
	Users    []net.Conn
}

//NewChatServer return a server object
func NewChatServer(protocol string, port string) *ChatServer {
	cs := new(ChatServer)
	cs.Protocol = protocol
	cs.Port = port
	cs.Users = make([]net.Conn, 0)
	return cs
}

//Start chat server
func (cs ChatServer) Start() {
	listener, err := net.Listen(cs.Protocol, cs.Port)
	if err != nil {
		log.Fatalf("unable to start server: %s", err)
	}
	defer listener.Close()

	log.Printf("Chat server started on port " + cs.Port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept connection: %s", err)
			continue
		}

		io.WriteString(conn, "Bienvenido al chat de GDG Marbella!\n")

		cs.Users = append(cs.Users, conn)

		go func() {
			for {
				msg, err := bufio.NewReader(conn).ReadString('\n')
				if err != nil {
					log.Println(err)
					continue
				}

				for _, user := range cs.Users {
					if user.RemoteAddr().String() != conn.RemoteAddr().String() {
						io.WriteString(user, msg)
					}
				}
			}
		}()

	}
}
