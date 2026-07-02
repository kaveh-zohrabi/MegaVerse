# MegaVerse CLI

Command-line interface for interacting with MegaVerse.

## Installation

```bash
# Homebrew
brew install megaverse/tap/megaverse

# Go
go install megaverse.dev/cli@latest

# Direct download
curl -fsSL https://get.megaverse.dev/cli | sh
```

## Commands

```bash
# Authentication
megaverse auth login
megaverse auth logout
megaverse auth status

# Projects
megaverse project list
megaverse project create
megaverse project delete

# Services
megaverse service list
megaverse service deploy
megaverse service status

# Database
megaverse db migrate
megaverse db seed
megaverse db backup

# AI
megaverse ai train
megaverse ai evaluate
megaverse ai deploy

# Config
megaverse config set
megaverse config get
megaverse config list
```

## Configuration

Config file: `~/.megaverse/config.yaml`

```yaml
api_key: your-api-key
base_url: https://api.megaverse.dev
output: json
```

## Documentation

See [docs.megaverse.dev/cli](https://docs.megaverse.dev/cli)
