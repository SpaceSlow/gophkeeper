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
	secretKey            string
	keyLen               int
	passwordIterationNum int
	tokenLifetime        time.Duration
	TimeoutShutdown      time.Duration
}

func (c *ServerConfig) SecretKey() string {
	return c.secretKey
}

func (c *ServerConfig) KeyLen() int {
	return c.keyLen
}

func (c *ServerConfig) PasswordIterationNum() int {
	return c.passwordIterationNum
}

func (c *ServerConfig) TokenLifetime() time.Duration {
	return c.tokenLifetime
}

var defaultConfig = &ServerConfig{
	NetAddress:           netaddress.NetAddress{Port: 443},
	CertificatePath:      os.Getenv("CERTIFICATE"),
	PrivateKeyPath:       os.Getenv("PRIVATE_KEY"),
	DSN:                  os.Getenv("DSN"),
	secretKey:            os.Getenv("SECRET_KEY"),
	tokenLifetime:        time.Hour,
	keyLen:               32,
	passwordIterationNum: 500_000,
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
