package server

import (
	"net/http"

	"github.com/rakyll/statik/fs"
)

func Run(c *Config) error {
	if c.Web {
		var appFS http.FileSystem
		if c.WebPath != "" {
			appFS = http.Dir(c.WebPath)
		} else {
			appFS, _ = fs.New()
		}
		http.Handle("/", http.FileServer(appFS))
	}
	return http.ListenAndServe(c.Address, nil)
}
