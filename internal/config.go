package internal

import (
	"os"
	"sync"
	"time"

	"github.com/SpaceSlow/gophkeeper/pkg/netaddress"
)

type ServerConfig struct {
	NetAddress           netaddress.NetAddress
	CertificatePath      string
	PrivateKeyPath       string
	DSN                  string
	SecretKey            string
	KeyLen               int
	PasswordIterationNum int
	TokenLifetime        time.Duration
	TimeoutShutdown      time.Duration
}

var defaultConfig = &ServerConfig{
	NetAddress:           netaddress.NetAddress{Port: 443},
	CertificatePath:      os.Getenv("CERTIFICATE"),
	PrivateKeyPath:       os.Getenv("PRIVATE_KEY"),
	DSN:                  os.Getenv("DSN"),
	SecretKey:            os.Getenv("SECRET_KEY"),
	TokenLifetime:        time.Hour,
	KeyLen:               32,
	PasswordIterationNum: 500_000,
	TimeoutShutdown:      5 * time.Second,
}

var serverConfig *ServerConfig = nil
var once sync.Once

func LoadServerConfig() *ServerConfig {
	once.Do(func() {
		serverConfig = defaultConfig
	})
	return serverConfig
}
