package config

import (
	"os"
	"time"
)

type ServerConfig struct {
	NetAddress      NetAddress
	CertificatePath string
	PrivateKeyPath  string

	TimeoutShutdown time.Duration
}

var DefaultConfig = &ServerConfig{
	NetAddress:      NetAddress{Port: 443},
	CertificatePath: os.Getenv("CERTIFICATE"),
	PrivateKeyPath:  os.Getenv("PRIVATE_KEY"),
	TimeoutShutdown: 5 * time.Second,
}
