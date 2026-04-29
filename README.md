# Caching Proxy

A lightweight HTTP caching proxy server built with Go.  
It forwards incoming requests to an origin server, caches responses, and serves repeated requests directly from cache to improve performance and reduce backend load.

This project was built as part of the [roadmap.sh Backend Projects](https://roadmap.sh/projects/caching-server).

---

## Features

- Forward HTTP requests to any origin server
- Cache responses for repeated requests
- Faster response time for cached data
- Reduce load on upstream services
- Configurable proxy port
- Clear cache manually
- Simple CLI usage
- Lightweight and fast (written in Go)

---

## How It Works

1. Client sends request to the proxy server  
2. Proxy forwards request to origin server  
3. Response is cached in memory  
4. Future identical requests are served from cache  
5. Cache can be cleared manually when needed

---



```bash
git clone https://github.com/HosseinForouzan/caching-proxy.git
cd caching-proxy
