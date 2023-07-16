package dev_browser

import (
	"context"
	"strconv"

	"github.com/cdvelop/gotools"
	"github.com/chromedp/chromedp"
)

func (b *Browser) CreateBrowserContext() {
	var position = "0,0" // ej: "1930,0"

	width, height, err := GetScreenSize()
	if err != nil {
		gotools.ShowErrorAndExit(err.Error())
	}

	if b.with != "" {
		width, err = strconv.Atoi(b.with)
		if err != nil {
			gotools.ShowErrorAndExit(err.Error())
		}
	}

	if b.height != "" {
		height, err = strconv.Atoi(b.height)
		if err != nil {
			gotools.ShowErrorAndExit(err.Error())
		}
	}

	if b.position != "" {
		position = b.position
	}

	// fmt.Printf("tamaño monitor: [%d] x [%d] position: [%v]\n", width, height, position)

	opts := append(

		// select all the elements after the third element
		chromedp.DefaultExecAllocatorOptions[:],
		// chromedp.NoDefaultBrowserCheck,
		chromedp.Flag("headless", false), // Desactivar el modo headless

		// chromedp.NoFirstRun,
		// chromedp.NoDefaultBrowserCheck,
		chromedp.Flag("auto-open-devtools-for-tabs", true),
		//quitar mensaje: Chrome is being controlled by automated test software

		// chromedp.Flag("--webview-log-js-console-messages", true),
		chromedp.WindowSize(width, height),
		chromedp.Flag("window-position", position),
		// chromedp.WindowSize(1530, 870),
		// chromedp.Flag("window-position", "1540,0"),
		chromedp.Flag("use-fake-ui-for-media-stream", true),
		// chromedp.Flag("exclude-switches", "enable-automation"),
		// chromedp.Flag("disable-blink-features", "AutomationControlled"),
		// chromedp.NoFirstRun,
		// chromedp.NoDefaultBrowserCheck,
		// chromedp.Flag("disable-infobars", true),
		// chromedp.Flag("enable-automation", true),
		// chromedp.Flag("disable-infobars", true),
		// chromedp.Flag("exclude-switches", "disable-infobars"),
	)

	parentCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)

	b.Context, b.CancelFunc = chromedp.NewContext(parentCtx)
}
