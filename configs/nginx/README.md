# MegaVerse Nginx Configuration

## Overview

Nginx configuration for reverse proxy and load balancing.

## Configuration

```nginx
upstream api_gateway {
    least_conn;
    server api-gateway-1:8080;
    server api-gateway-2:8080;
    server api-gateway-3:8080;
}

server {
    listen 443 ssl http2;
    server_name api.megaverse.dev;

    ssl_certificate /etc/nginx/ssl/cert.pem;
    ssl_certificate_key /etc/nginx/ssl/key.pem;

    location / {
        proxy_pass http://api_gateway;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }

    location /ws {
        proxy_pass http://api_gateway;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
    }
}
```

## Features

- SSL termination
- Load balancing
- WebSocket support
- Rate limiting
- Gzip compression
- Security headers
