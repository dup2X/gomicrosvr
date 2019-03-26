package discovery

import (
	"time"
)

type Balancer interface {
	Get() (ed *Endpoint, err error)
	Vote(host string, port int, score int)
}

type Endpoint struct {
	Host         string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
