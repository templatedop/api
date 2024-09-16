package api

import (
	"time"
)

type options struct {
	timeout time.Duration
}

func getOptions(opt ...Option) *options {
	def := &options{
		timeout: 45 * time.Second,
	}

	for _, o := range opt {
		o.Apply(def)
	}

	return def
}

type Option interface {
	Apply(*options)
}
