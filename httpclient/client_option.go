package httpclient

import (
	"net"
	"time"
)

// ClientOption implements optional pattern to configure Client object
type ClientOption func(*Client)

// WithMaxConns set maximum number of connections per each host which may be established
//
// DefaultMaxConnsPerHost is used if not set
func WithMaxConns(conns int) ClientOption {
	return func(c *Client) {
		c.lib.MaxConnsPerHost = conns
	}
}

// WithIdleKeepAliveDuration ...
func WithIdleKeepAliveDuration(duration time.Duration) ClientOption {
	return func(c *Client) {
		c.lib.MaxIdleConnDuration = duration
	}
}

// WithDialTimeout ...
func WithDialTimeout(duration time.Duration) ClientOption {
	return func(c *Client) {
		c.lib.Dial = func(addr string) (net.Conn, error) {
			conn, err := net.DialTimeout("tcp", addr, duration)
			if err != nil {
				return conn, err
			}
			return conn, nil
		}
	}
}

// WithMaxIdemponentCallAttempts retry number
func WithMaxIdemponentCallAttempts(idemponent int) ClientOption {
	return func(c *Client) {
		c.lib.MaxIdemponentCallAttempts = idemponent
	}
}
