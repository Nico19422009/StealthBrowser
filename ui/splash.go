package ui

import (
    "embed"
    "time"
    "github.com/webview/webview"
)

//go:embed splash.html
var splashHTML embed.FS

func ShowSplash(duration time.Duration) {
    htmlData, _ := splashHTML.ReadFile("splash.html")
    w := webview.New(false)
    defer w.Destroy()
    w.SetTitle("StealthBrowser Loading...")
    w.SetSize(500, 300, webview.HintNone)
    w.Navigate("data:text/html," + string(htmlData))
    go func() {
        time.Sleep(duration)
        w.Terminate()
    }()
    w.Run()
}

func ShowError(msg string) {
    w := webview.New(false)
    defer w.Destroy()
    w.SetTitle("❌ Error")
    w.SetSize(500, 200, webview.HintNone)
    w.Navigate("data:text/html,<h2 style='color:red;font-family:sans-serif'>" + msg + "</h2>")
    w.Run()
}
