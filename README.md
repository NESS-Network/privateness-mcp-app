# privateness-mcp-app

MCP over HTTP/3 WebTransport — standard-compliant, browser-ready.

- Auth: **Privateness blockchain keys** (EmerSSH optional bridge to OpenSSH; not used for MCP auth)
- DNS: **EmerDNS** (RFC-1035, UDP 53/5335) via your dns-reverse-proxy
- Transports (optional): Skywire, Yggdrasil, I2P, Tor, AmneziaWG, AmneziaXRAY, IPFS/BitTorrent (EmerMagnet)
- Streams: `cmd`, `resp`, `notify`, `file`, `dht`, plus transport streams
- Metering: ephemeral per-session bandwidth/time
- Billing: **Privateness coin hours** (enabled here)

## Quick start
```bash
make build
# Dev self-signed certs auto-generated on first run into ./devcert/
make run
```

## Env
- `TLS_CERT`, `TLS_KEY` — TLS1.3 cert & key (if absent, self-signed dev cert is created)
- `LISTEN` — default `:443`
- `BASE_PATH` — default `/mcp`

## Notes
- Repos are public; configs emphasize private usage. Blockchain is the single source of truth; **MCP is the only interface**.
