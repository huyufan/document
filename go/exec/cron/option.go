package cron

import "time"

type Option func(*Cron)

func WithLocation(loc *time.Location) Option {
	return func(c *Cron) {
		c.location = loc
	}
}

func WithSeconds() Option {
	return WithParser(N)
}

func WithParser(p ScheduleParser) Option {
	return func(c *Cron) {
		c.parser = p
	}
}
