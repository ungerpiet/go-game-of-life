# Conway's Game of Life

[Live-Demo ansehen](https://ungerpiet.github.io/go-game-of-life/)

Eine interaktive Umsetzung von Conways Game of Life in Go mit Ebitengine. Spiele direkt im Browser oder starte die Anwendung lokal.

## Features

- Klassische Zell-Automation auf einem Raster mit verbundenen Rändern
- Simulation pausieren, zurücksetzen und mit Zufallsmustern neu starten
- Zellen im Pausemodus interaktiv setzen und löschen

## Quickstart

Lokal starten:

```sh
go run .
```

WebAssembly bauen:

```sh
GOOS=js GOARCH=wasm go build -o main.wasm .
```
