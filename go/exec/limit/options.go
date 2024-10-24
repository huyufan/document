package limit

import "net"

type Option func(*Options)

type Options struct {
	IPv4Mask           net.IPMask
	IPv6Mask           net.IPMask
	TrustForwardHeader bool
	ClientIPHeader     string
}

func WithIPv4Mask(mask net.IPMask) Option {
	return func(o *Options) {
		o.IPv4Mask = mask
	}
}

func WithIPv6Mask(mask net.IPMask) Option {
	return func(o *Options) {
		o.IPv6Mask = mask
	}
}

func WithTrustForwardHeader(enable bool) Option {
	return func(o *Options) {
		o.TrustForwardHeader = enable
	}
}

func WithClientIPHeader(header string) Option {
	return func(o *Options) {
		o.ClientIPHeader = header
	}
}
