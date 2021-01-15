package rest

import (
	"time"
)

type ClientOptions struct {
	Timeout time.Duration
}

type ClientOption func(*ClientOptions)

func Timeout(t time.Duration) ClientOption {
	return func(o *ClientOptions) {
		o.Timeout = t
	}
}
