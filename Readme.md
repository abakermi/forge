# Forge - Load Testing Utility
[![codecov](https://codecov.io/github/abakermi/forge/branch/codecov-integration/graph/badge.svg?token=64IHXT3ROF)](https://codecov.io/github/abakermi/forge)
[![Check & Build](https://github.com/abakermi/forge/actions/workflows/ci.yml/badge.svg)](https://github.com/abakermi/forge/actions/workflows/ci.yml)
[![License: AGPL v3](https://img.shields.io/badge/License-AGPL_v3-blue.svg)](https://www.gnu.org/licenses/agpl-3.0)

Forge is a simple and efficient load testing utility written in Go. It allows you to simulate load on web services by sending HTTP requests or WebSocket messages at a specified rate.

## Installation

To install Forge, use `go get`:

```bash
go get github.com/abakermi/forge
```
## Usage
To use Forge, simply run the forge command followed by the URL you want to test:

```bash
forge https://example.com
```
You can also specify the concurrency level and requests per second (RPS) using flags:

```bash
forge -c 10 -rps 100 https://example.com
```
This command will simulate load with 10 concurrent clients sending 100 requests per second to `https://example.com`.
## Options

- `**-c, --concurrency**`: Number of concurrent clients (default is 1).
- `**-r, --rps**`: Requests per second (default is 1).

## Supported Schemes
Forge supports both HTTP and WebSocket protocols. It automatically detects the scheme of the provided URL and uses the appropriate client:

- For HTTP/HTTPS URLs, Forge sends HTTP GET requests.
- For WS/WSS URLs, Forge establishes WebSocket connections and sends ping messages.

## License

See [`LICENSE`](./LICENSE)

