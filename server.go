package main

import (
	"log"
	"net"
	"os"

	"github.com/armon/go-socks5"
	"github.com/caarlos0/env/v6"
)

type params struct {
	BindIP   string `env:"BIND_IP" envDefault:"0.0.0.0"`
	BindPort string `env:"BIND_PORT" envDefault:"1080"`
}

func main() {
	cfg := params{}
	err := env.Parse(&cfg)
	if err != nil {
		log.Printf("%+v\n", err)
	}

	socks5conf := &socks5.Config{
		BindIP: net.ParseIP(cfg.BindIP),
		Logger: log.New(os.Stdout, "", log.LstdFlags),
	}

	cator := socks5.NoAuthAuthenticator{}
	socks5conf.AuthMethods = []socks5.Authenticator{cator}

	server, err := socks5.New(socks5conf)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Start listening proxy service on %s:%s\n", cfg.BindIP, cfg.BindPort)
	if err := server.ListenAndServe("tcp", cfg.BindIP+":"+cfg.BindPort); err != nil {
		log.Fatal(err)
	}
}
