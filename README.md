# realtimex

A lightweight, transport-agnostic **real-time communication library for Go** supporting:

- WebSockets
- Server-Sent Events (SSE)
- Channel-based broadcasting
- Gin integration
- Clean event-driven architecture

Designed for building **POS systems, dashboards, live notifications, chat systems, and real-time APIs**.

---

## Features

- Unified event model for WS + SSE
- Simple broadcaster/manager API
- Channel/room support (optional extension)
- WebSocket support via Gorilla WebSocket
- SSE streaming with flushing support
- Gin HTTP adapters included
- Framework-agnostic core

---

## Installation

```bash
go get github.com/danielcomboni/realtimex