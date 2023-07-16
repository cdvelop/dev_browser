package dev_browser

import (
	"log"

	"github.com/chromedp/chromedp"
)

func (b *Browser) Reload() {
	// fmt.Println("Recargando Navegador")
	err := chromedp.Run(b.Context, chromedp.Reload())
	if err != nil {
		log.Println("Error al recargar Pagina ", err)
	}

}
