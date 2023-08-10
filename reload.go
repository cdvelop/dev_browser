package dev_browser

import (
	"fmt"

	"github.com/chromedp/chromedp"
)

func (b *Browser) Reload() error {
	// fmt.Println("Recargando Navegador")
	err := chromedp.Run(b.Context, chromedp.Reload())
	if err != nil {
		return fmt.Errorf("error al recargar Pagina %v", err)
	}

	return nil
}
