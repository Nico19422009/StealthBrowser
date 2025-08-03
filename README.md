# StealthBrowser 🕵️‍♂️

**StealthBrowser** is a modern internet browser designed for **anonymity**, **privacy**, and **freedom of information**. It provides a clean and minimalistic interface while routing all traffic through the **Tor network** to avoid surveillance, tracking, and censorship.

---

## 🔒 Features

- 🧅 Connects to the Tor network on startup
- 🕵️ Anonymous browsing out of the box
- 🧼 No telemetry, tracking or fingerprinting
- ⚡ Lightweight and fast (built in Go with native WebView)
- 🖥️ Works on Windows, Linux and macOS
- 💻 Splash screen and stealth loading
- ⭐ Tabs, bookmarks, history

---

## 🚀 Getting Started (Windows)

> ⚠️ You **must download the official [Tor Expert Bundle](https://www.torproject.org/download/tor/)**

1. Download the [Tor Expert Bundle](https://www.torproject.org/download/tor/)
2. Extract the archive
3. Copy `tor.exe` into the `tor/` subfolder inside the StealthBrowser directory  
   → Final path: `StealthBrowser/tor/tor.exe`
4. Run `StealthBrowser.exe` or build from source

✅ Now you're ready to browse `.onion` and clearnet sites anonymously!

---

## 🧠 Why StealthBrowser?

In today's world, freedom of information is under attack. StealthBrowser is built to fight back by providing a tool that:

- Protects user identity
- Bypasses regional restrictions
- Blocks trackers and censorship

---

## 🛠️ Build From Source (Linux/Mac/Windows)

```bash
go get github.com/webview/webview
go get github.com/cretz/bine/tor

go run main.go
