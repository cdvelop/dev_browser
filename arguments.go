package dev_browser

import (
	"fmt"
	"os"

	. "github.com/cdvelop/output"
	"github.com/cdvelop/strings"
)

// path: ej: /index.html, / ,/home

func (b *Browser) captureArguments() {

	for _, opt := range os.Args {

		switch {

		case opt == "dev":
			b.dev_mode = true

		case strings.Contains(opt, "path:") == 1:
			strings.ExtractTwoPointArgument(opt, &b.path)
			continue

		case strings.Contains(opt, "with:") == 1:
			strings.ExtractTwoPointArgument(opt, &b.with)
			continue

		case strings.Contains(opt, "height:") == 1:
			strings.ExtractTwoPointArgument(opt, &b.height)
			continue

		case strings.Contains(opt, "position:") == 1:
			strings.ExtractTwoPointArgument(opt, &b.position)
			continue

		case opt == "help" || opt == "?" || opt == "ayuda":

			fmt.Println(`
			
			- default: port:8080  or get env var = "APP_PORT"

			- protocol default http 
			    if port contain number 44 protocol is https.

			--- arguments path:

			    path:/login, /home default /
			
			--- arguments Browser ---

			with:800
			height:600
			position:1930,0
			*-position if you have second monitor
			
			`)

			ShowErrorAndExit("")
		}

	}

	if b.path == "" {
		b.path = "/"
	}

}
