package browser

import (
    "embed"
    "log"
    "github.com/webview/webview"
)

//go:embed static/gui.html
var staticFiles embed.FS

func Start(proxy string) {
    htmlData, err := staticFiles.ReadFile("static/gui.html")
    if err != nil {
        log.Fatal("GUI HTML not found:", err)
    }

    w := webview.New(true)
    defer w.Destroy()
    w.SetTitle("StealthBrowser 🕵️")
    w.SetSize(1024, 768, webview.HintNone)

    w.Bind("navigateTo", func(url string) {
        w.Eval(`document.getElementById("urlbar").value = "` + url + `";`)
        w.Eval(`document.getElementById("tab-frame-" + currentTab).src = "` + url + `";`)
        w.Eval(`addToHistory("` + url + `");`)
    })

    w.Navigate("data:text/html," + string(htmlData))
    w.Run()
}
