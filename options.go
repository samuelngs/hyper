package hyper

import (
	"crypto/rand"
	"fmt"

	"github.com/samuelngs/hyper/cache"
	"github.com/samuelngs/hyper/engine"
	"github.com/samuelngs/hyper/message"
	"github.com/samuelngs/hyper/router"
	"github.com/samuelngs/hyper/sync"
)

// Option func
type Option func(*Options)

// Options is the hyper server options
type Options struct {

	// hyper server unique id
	ID string

	// server bind to address [host:port]
	Addr string

	// HTTP protocol 1.1 / 2.0
	Protocol engine.Protocol

	// sync engine
	Sync sync.Service

	// Cache engine
	Cache cache.Service

	// Router
	Router router.Service

	// Message broker
	Message message.Service

	// before and after funcs
	BeforeStart []func() error
	AfterStop   []func() error

	// AllowedOrigins is a list of origins a cross-domain request can be executed from.
	// If the special "*" value is present in the list, all origins will be allowed.
	// An origin may contain a wildcard (*) to replace 0 or more characters
	// (i.e.: http://*.domain.com). Usage of wildcards implies a small performance penality.
	// Only one wildcard can be used per origin.
	// Default value is ["*"]
	AllowedOrigins []string

	// AllowOriginFunc is a custom function to validate the origin. It take the origin
	// as argument and returns true if allowed or false otherwise. If this option is
	// set, the content of AllowedOrigins is ignored.
	AllowOriginFunc func(origin string) bool

	// AllowedMethods is a list of methods the client is allowed to use with
	// cross-domain requests. Default value is simple methods (GET and POST)
	AllowedMethods []string

	// AllowedHeaders is list of non simple headers the client is allowed to use with
	// cross-domain requests.
	// If the special "*" value is present in the list, all headers will be allowed.
	// Default value is [] but "Origin" is always appended to the list.
	AllowedHeaders []string

	// ExposedHeaders indicates which headers are safe to expose to the API of a CORS
	// API specification
	ExposedHeaders []string

	// AllowCredentials indicates whether the request can include user credentials like
	// cookies, HTTP authentication or client side SSL certificates.
	AllowCredentials bool

	// MaxAge indicates how long (in seconds) the results of a preflight request
	// can be cached
	MaxAge int

	// OptionsPassthrough instructs preflight to let other potential next handlers to
	// process the OPTIONS method. Turn this on if your application handles OPTIONS.
	OptionsPassthrough bool
}

func newID() string {
	b := new([16]byte)
	rand.Read(b[:])
	b[8] = (b[8] | 0x40) & 0x7F
	b[6] = (b[6] & 0xF) | (4 << 4)
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func newOptions(opts ...Option) Options {
	opt := Options{
		ID:       newID(),
		Addr:     ":0",
		Protocol: engine.HTTP,
	}
	for _, o := range opts {
		o(&opt)
	}
	if opt.Cache == nil {
		opt.Cache = cache.New(
			cache.ID(opt.ID),
		)
	}
	if opt.Message == nil {
		opt.Message = message.New(
			message.ID(opt.ID),
		)
	}
	if opt.Sync == nil {
		opt.Sync = sync.New(
			sync.ID(opt.ID),
			sync.Cache(opt.Cache),
			sync.Message(opt.Message),
		)
	}
	if opt.Router == nil {
		opt.Router = router.New(
			router.ID(opt.ID),
		)
	}
	return opt
}

// ID to change server reference id
func ID(s string) Option {
	return func(o *Options) {
		o.ID = s
	}
}

// Addr to change server bind address
func Addr(s string) Option {
	return func(o *Options) {
		o.Addr = s
	}
}

// HTTP to use 1.1 network protocol
func HTTP() Option {
	return func(o *Options) {
		o.Protocol = engine.HTTP
	}
}

// HTTP2 to use 2.0 network protocol
func HTTP2() Option {
	return func(o *Options) {
		o.Protocol = engine.HTTP2
	}
}

// Sync to set custom sync engine
func Sync(s sync.Service) Option {
	return func(o *Options) {
		o.Sync = s
	}
}

// Cache to set custom cache engine
func Cache(c cache.Service) Option {
	return func(o *Options) {
		o.Cache = c
	}
}

// Router to set custom router
func Router(r router.Service) Option {
	return func(o *Options) {
		o.Router = r
	}
}

// Message to set custom message broker
func Message(m message.Service) Option {
	return func(o *Options) {
		o.Message = m
	}
}

// BeforeStart to add before start action to hyper server
func BeforeStart(f func() error) Option {
	return func(o *Options) {
		o.BeforeStart = append(o.BeforeStart, f)
	}
}

// AfterStop to add after stop action to hyper server
func AfterStop(f func() error) Option {
	return func(o *Options) {
		o.AfterStop = append(o.AfterStop, f)
	}
}

func AllowedOrigins(a []string) Option {
	return func(o *Options) {
		o.AllowedOrigins = append(o.AllowedOrigins, a...)
	}
}

func AllowOriginFunc(f func(string) bool) Option {
	return func(o *Options) {
		o.AllowOriginFunc = f
	}
}

func AllowedMethods(a []string) Option {
	return func(o *Options) {
		o.AllowedMethods = append(o.AllowedMethods, a...)
	}
}

func AllowedHeaders(a []string) Option {
	return func(o *Options) {
		o.AllowedHeaders = append(o.AllowedMethods, a...)
	}
}

func ExposedHeaders(a []string) Option {
	return func(o *Options) {
		o.ExposedHeaders = append(o.AllowedMethods, a...)
	}
}

func AllowCredentials(b bool) Option {
	return func(o *Options) {
		o.AllowCredentials = b
	}
}

func MaxAge(i int) Option {
	return func(o *Options) {
		o.MaxAge = i
	}
}

func OptionsPassthrough(b bool) Option {
	return func(o *Options) {
		o.OptionsPassthrough = b
	}
}
