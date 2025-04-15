package ghttp

import (
	"crypto/tls"
	"exec/go/exec/grace/gnet"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	logger     *log.Logger
	didInherit = os.Getenv("LISTEN_FDS") != ""
	ppid       = os.Getppid()
)

type option func(*app)

type app struct {
	servers []*http.Server
	//http    *httpdown.http
	net       *gnet.Net
	listeners []net.Listener
	//sds  []httpdown.Server
	preStartProcess func() error
	errors          chan error
}

func newApp(servers []*http.Server) *app {
	return &app{
		servers:         servers,
		net:             &gnet.Net{},
		listeners:       make([]net.Listener, 0, len(servers)),
		preStartProcess: func() error { return nil },
		errors:          make(chan error, 1+(len(servers)*2)),
	}
}

func (a *app) listen() error {
	for _, s := range a.servers {
		l, err := a.net.Listen("tcp", s.Addr)
		if err != nil {
			return err
		}
		if s.TLSConfig != nil {
			l = tls.NewListener(l, s.TLSConfig)
		}
		a.listeners = append(a.listeners, l)
	}
	return nil
}

func (a *app) singleHandler(wg *sync.WaitGroup) {
	ch := make(chan os.Signal, 10)
	//signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR2)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	for {
		sig := <-ch
		switch sig {
		//case syscall.SIGINT, syscall.SIGTERM:
		case syscall.SIGTERM:
			signal.Stop(ch)
			return
		//case syscall.SIGUSR2:
		case syscall.SIGINT:
			err := a.preStartProcess()
			if err != nil {
				a.errors <- err
			}
			if _, err := a.net.StartProcess(); err != nil {
				a.errors <- err
			}
		}
	}
}
func (a *app) serve() {
	// for i, s := range a.servers {

	// }
}

func (a *app) Wait() {
	var wg sync.WaitGroup
	//wg.Add(len())
}
