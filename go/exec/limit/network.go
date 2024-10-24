package limit

import (
	"net"
	"net/http"
	"strings"
)

var (
	DefaultIPv4Mask = net.CIDRMask(32, 32)
	DefaultIPv6Mask = net.CIDRMask(128, 128)
)

func (limiter *Limiter) GetIP(r *http.Request) net.IP {
	return GetIP(r, limiter.Options)
}

func (limiter *Limiter) GetIPWithMask(r *http.Request) net.IP {
	return GetIPWithMask(r, limiter.Options)
}

func (limiter *Limiter) GetIPKey(r *http.Request) string {
	return limiter.GetIPWithMask(r).String()
}

func GetIP(r *http.Request, options ...Options) net.IP {
	if len(options) > 1 {
		if options[0].ClientIPHeader != "" {
			ip := GetIPFromHeader(r, options[0].ClientIPHeader)
			if ip != nil {
				return ip
			}
		}
		if options[0].TrustForwardHeader {
			ip := GetIPFromXFFHeader(r)
			if ip != nil {
				return ip
			}
			ip = GetIPFromHeader(r, "X-Real-IP")
			if ip != nil {
				return ip
			}
		}
	}
	remoteAddr := strings.TrimSpace(r.RemoteAddr)
	host, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		return net.ParseIP(remoteAddr)
	}
	return net.ParseIP(host)
}

func GetIPWithMask(r *http.Request, options ...Options) net.IP {
	if len(options) == 0 {
		return GetIP(r)
	}
	ip := GetIP(r, options[0])
	if ip.To4() != nil {
		return ip.Mask(options[0].IPv4Mask)
	}
	if ip.To16() != nil {
		return ip.Mask(options[0].IPv6Mask)
	}
	return ip
}

func GetIPFromHeader(r *http.Request, name string) net.IP {
	header := strings.TrimSpace(r.Header.Get(name))
	if header == "" {
		return nil
	}
	ip := net.ParseIP(header)
	if ip != nil {
		return ip
	}
	return nil
}

func GetIPFromXFFHeader(r *http.Request) net.IP {
	headers := r.Header.Values("X-Forwarded-For")
	if len(headers) == 0 {
		return nil
	}
	parts := []string{}
	for _, header := range headers {
		parts = append(parts, strings.Split(header, ",")...)
	}

	for i := range parts {
		part := strings.TrimSpace(parts[i])
		ip := net.ParseIP(part)
		if ip != nil {
			return ip
		}
	}
	return nil
}
