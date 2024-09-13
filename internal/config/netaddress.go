package config

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrIncorrectNetAddress = errors.New("need address in a form host:port")
	ErrIncorrectPort       = errors.New("error occurred when parsing an incorrect port. The port requires a decimal number in the range 1-65535")
)

type NetAddress struct {
	Host string
	Port int
}

func (a NetAddress) String() string {
	if a.Host == "" && a.Port == 0 {
		return ""
	}
	return a.Host + ":" + strconv.Itoa(a.Port)
}

func (a *NetAddress) Set(s string) error {
	hp := strings.Split(s, ":")
	if len(hp) != 2 {
		return ErrIncorrectNetAddress
	}
	port, err := strconv.Atoi(hp[1])
	if err != nil || port < 1 || port > 65535 {
		return ErrIncorrectPort
	}
	a.Host = hp[0]
	a.Port = port
	return nil
}
