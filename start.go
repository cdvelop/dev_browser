package dev_browser

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/chromedp/chromedp"
)

func (b *Browser) DevBrowserSTART(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("*** START DEV BROWSER ***")

	b.CreateBrowserContext()
	// defer cancel()
	var protocol = "http"

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	} else {
		if strings.Contains(port, "44") {
			protocol = "https"
		}
	}

	// Navega a una página web

	url := protocol + `://localhost:` + port + b.path

	err := chromedp.Run(b.Context, b.sendkeys(url)) // chromedp.Navigate(url),

	if err != nil {
		log.Fatal("Error al navegar "+b.path+" ", err)
	}

	// Espera hasta que la página esté completamente cargada
	var loaded bool

	err = chromedp.Run(b.Context, chromedp.ActionFunc(func(ctx context.Context) error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				var readyState string
				err := chromedp.Run(ctx, chromedp.EvaluateAsDevTools(`document.readyState`, &readyState))
				if err != nil {
					return err
				}

				if readyState == "complete" {
					loaded = true

					return nil
				}
			}
		}
	}))
	if err != nil {
		log.Fatal(err)
	}

	// Verifica si la página se ha cargado correctamente
	if !loaded {
		log.Fatal("La página no se ha cargado correctamente")
	}

}

func (b Browser) sendkeys(host string) chromedp.Tasks {

	return chromedp.Tasks{
		chromedp.Navigate(host),
	}
}
