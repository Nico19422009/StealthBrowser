package main

import (
    "fmt"
    "time"

    "stealth/tor_conn"
    "stealth/browser"
    "stealth/ui"
)

func main() {
    ui.ShowSplash(3 * time.Second)

    fmt.Println("🔄 Connecting to Tor...")
    tor, proxyAddr, err := tor_conn.StartTor()
    if err != nil {
        ui.ShowError("Tor connection failed.\nMake sure tor.exe is present.")
        return
    }
    defer tor.Close()

    fmt.Println("✅ Tor ready at:", proxyAddr)
    browser.Start(proxyAddr)
}
