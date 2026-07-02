# MegaVerse CLI

Command-line interface for interacting with MegaVerse.

## Installation

```bash
# Build from source
go build -o megaverse ./cmd

# Or install directly
go install megaverse.dev/sdk/cli/cmd@latest
```

## Commands

| Command | Description |
|---------|-------------|
| `megaverse version` | Show CLI version |
| `megaverse health` | Check API health |
| `megaverse config set <key> <value>` | Set config value |
| `megaverse config get <key>` | Get config value |
| `megaverse config list` | List all config |

## Configuration

```bash
# Set API URL
megaverse config set base-url http://localhost:8080

# Set API key
megaverse config set api-key your-api-key

# View config
megaverse config list
```

## Development

```bash
# Run
go run ./cmd

# Build
go build -o megaverse ./cmd

# Test
go test ./...
```
