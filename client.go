package redkacli

import (
	"github.com/nalgeon/redka"
	"log/slog"
)

type Client struct {
	db *redka.DB
}

type Options struct {
	path       string
	driverName string
	pragma     map[string]string
	logger     *slog.Logger
}

func NewOptions() *Options {
	return &Options{
		path:       "",
		driverName: "sqlite3",
		pragma:     make(map[string]string),
	}
}

func (opts *Options) SetDriverName(driverName string) {
	opts.driverName = driverName
}

func (opts *Options) SetLogger(logger *slog.Logger) {
	opts.logger = logger
}

func (opts *Options) SetPragma(pragma map[string]string) {
	opts.pragma = pragma
}

func (opts *Options) AddPragma(key, value string) {
	opts.pragma[key] = value
}

func New(o *Options) (*Client, error) {
	db, err := redka.Open(o.path, &redka.Options{
		DriverName: o.driverName,
		Pragma:     o.pragma,
		Logger:     o.logger,
	})
	if err != nil {
		return nil, err
	}
	return &Client{db: db}, nil
}
