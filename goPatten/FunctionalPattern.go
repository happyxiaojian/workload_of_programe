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

//======================= 观察者模式 =======================

type Observer interface {
	Update(subject *Subject)
}

type Subject struct {
	Observers []Observer
	context string
}

func NewSubject(cs string)(s *Subject){
	return &Subject{
		Observers: make([]Observer, 0),
	}
}

func (s *Subject)Attach(o Observer){
	s.Observers = append(s.Observers, o)
}

func (s *Subject)Remove(o Observer)bool{
	for i, j := range s.Observers {
		if j == o {
			s.Observers = append(s.Observers[:i], s.Observers[i+1:]...)
			return true
		}
	}
	return false
}

func (s *Subject)Notify(){
	for _, i := range s.Observers {
		i.Update(s)
	}
}

func (s *Subject)UpdateContext(c string){
	s.context = c
	s.Notify()
}

type Reader struct {
	name string
}

