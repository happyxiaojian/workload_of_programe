package main

import (
	"crypto/tls"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"net"
	"time"
)

type Server struct {
	listener net.Listener
	Name string
	timeout time.Duration
}


type option func(*Server)

func NewServer(addr string, options ...option)(*Server, error){
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	srv := Server{listener:l}

	for _, option := range options {
		option(&srv)
	}

	return &srv, nil
}

func main(){
	srv, _ := NewServer(":8989") // defaults

	spew.Dump(srv)

	srv2, _ := NewServer(":8787", Timeout, Tls, Name)

	fmt.Println("==================================")
	spew.Dump(srv2)

}

func Timeout(srv *Server){
	srv.timeout = 49 * time.Second
}

func Tls(srv *Server){
	config := loadTLSConfig()
	srv.listener = tls.NewListener(srv.listener, &config)
}

func loadTLSConfig() tls.Config {
	return tls.Config{}
}

func Name(srv *Server){
	srv.Name = "landon"
}







