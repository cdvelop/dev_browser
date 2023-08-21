package dev_browser

import (
	"log"
	"strconv"

	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
)

func (b *Browser) jsLogs() {

	// crea un mapa para registrar los mensajes de log únicos
	uniqueLogs := make(map[string]bool)

	// captura los logs de JavaScript
	chromedp.ListenTarget(b.Context, func(ev interface{}) {
		switch ev := ev.(type) {
		case *runtime.EventConsoleAPICalled:
			for _, arg := range ev.Args {
				s, err := strconv.Unquote(string(arg.Value))
				if err != nil {
					log.Println("ERROR chromedp: ", err)
					continue
				}
				// verifica si el mensaje de log ya se ha registrado
				if !uniqueLogs[s] {
					uniqueLogs[s] = true
					// fmt.Printf("LOG: %s\n", s)
				}
			}
		}
	})
}
