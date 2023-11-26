package dev_browser

import "context"

type Browser struct {
	path     string // ej /index.html, /
	port     string //8080
	domain   string // localhost
	protocol string // ej https,

	with     string //browser option
	height   string //browser option
	position string //browser option

	context.Context
	context.CancelFunc

	dev_mode bool
}
