module github.com/cdvelop/dev_browser

go 1.20

require (
	github.com/cdvelop/strings v0.0.7
	github.com/chromedp/chromedp v0.9.3
)

require github.com/lxn/win v0.0.0-20210218163916-a377121e959e // indirect

require (
	github.com/cdvelop/output v0.0.16
	github.com/chromedp/cdproto v0.0.0-20231114014204-3e458d5176f9
	github.com/chromedp/sysutil v1.0.0 // indirect
	github.com/fstanis/screenresolution v0.0.0-20190527020317-869904d15333
	github.com/gobwas/httphead v0.1.0 // indirect
	github.com/gobwas/pool v0.2.1 // indirect
	github.com/gobwas/ws v1.3.1 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	golang.org/x/sys v0.14.0 // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/output => ../output

replace github.com/cdvelop/timetools => ../timetools

replace github.com/cdvelop/input => ../input

replace github.com/cdvelop/gotools => ../gotools

replace github.com/cdvelop/strings => ../strings
