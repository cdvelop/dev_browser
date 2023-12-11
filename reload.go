package dev_browser

import (
	"github.com/chromedp/chromedp"
)

func (b *Browser) Reload() (err string) {
	if b.Context != nil {
		// fmt.Println("Recargando Navegador")
		er := chromedp.Run(b.Context, chromedp.Reload())
		if er != nil {
			return "Reload error al recargar Pagina " + er.Error()
		}
	}
	return
}
