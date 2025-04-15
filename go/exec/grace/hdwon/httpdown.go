package hdown

import "time"

const (
	defaultStopTimeout = time.Minute
	defaultKillTimeout = time.Minute
)

type Server interface {
	Wait() error
	Stop() error
}

type HTTP struct {
	StopTimeout time.Duration
	KillTimeout time.Duration
	Stats       stats.Client
}
