package http

import "time"

type Config struct {
	Host    string
	Port    int
	Timeout time.Duration
}
