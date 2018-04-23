package server

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var (
	shutdownErr = errors.New("server stopped")
)

type Server struct {
	inner	 http.Server
	upgrader websocket.Upgrader
}

func New(address string) *Server {
	server := &Server {
		upgrader: websocket.Upgrader{
			HandshakeTimeout: time.Second,
			ReadBufferSize: 4096,
			WriteBufferSize: 4096,
			Subprotocols: []string{ "game"},
		},
		inner: http.Server{
			Addr: address,
			Handler: nil,
		},
	}

	http.Handle("/", http.FileServer(http.Dir("www")))
	http.Handle("/ws", server)
	return server
}

func (f *Server) Run(done <-chan struct{}) error {
	go (func() {
		<-done
		f.inner.Shutdown(nil)
	})()
	return f.inner.ListenAndServe()
}

func (f *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := f.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error in websocket negotiation: %v", err)
		return
	}
	typ, bytes, err := conn.ReadMessage()
	if err != nil {
		log.Printf("Error reading from ws: %v", err)
	}
	log.Printf("Read from ws: type=%d msg=%s", typ, bytes)
	conn.WriteJSON("Ola k ase")
	conn.Close()
}
