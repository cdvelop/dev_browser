package dev_browser

import (
	"fmt"
	"os"
	"strings"

	"github.com/cdvelop/gotools"
)

// path: ej: /index.html, / ,/home
// port: ej: 9090
// domain ej: 192.0.0.5,app.com default localhost
func (b *Browser) captureArguments() {

	for _, opt := range os.Args {

		switch {
		case strings.Contains(opt, "path:"):
			gotools.ExtractTwoPointArgument(opt, &b.path)
			continue

		case strings.Contains(opt, "port:"):
			gotools.ExtractTwoPointArgument(opt, &b.port)
			continue

		case strings.Contains(opt, "domain:"):
			gotools.ExtractTwoPointArgument(opt, &b.domain)
			continue

		case strings.Contains(opt, "with:"):
			gotools.ExtractTwoPointArgument(opt, &b.with)
			continue

		case strings.Contains(opt, "height:"):
			gotools.ExtractTwoPointArgument(opt, &b.height)
			continue

		case strings.Contains(opt, "position:"):
			gotools.ExtractTwoPointArgument(opt, &b.position)
			continue

		case opt == "help" || opt == "?" || opt == "ayuda":

			fmt.Println("default: port:8080 path:/")
			fmt.Println("*** ej valores admitidos***")
			fmt.Println("protocol:https default http")
			fmt.Println("domain:/192.168.0.2 default localhost")
			fmt.Println("port:9090 default 8080")
			fmt.Println("path:/login,/home default /")
			fmt.Println("-----------------------")
			fmt.Println("--- Browser Options ---")
			fmt.Println("with:800")
			fmt.Println("height:600")
			fmt.Println("position:1930,0")
			fmt.Println("*-position es en caso de que tengas segundo monitor")
			gotools.ShowErrorAndExit("")
		}

	}

	if b.path == "" {
		b.path = "/"
	}

	if b.port == "" {
		b.port = "8080"
	}

	if b.domain == "" {
		b.domain = "localhost"
	}

	if b.protocol == "" {
		b.protocol = "http"
	}

}
