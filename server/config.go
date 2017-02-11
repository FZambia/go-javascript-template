package server

type Config struct {
	// Address to serve HTTP server on.
	Address string
	// Web enables serving web interface.
	Web bool
	// WebPath allows to set custom directory to serve web app from.
	WebPath string
}
