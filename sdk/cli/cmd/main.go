package main

import (
	"fmt"
	"os"

	"megaverse.dev/sdk/cli/internal/commands"
	"megaverse.dev/sdk/cli/internal/config"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	args := os.Args[1:]
	command := args[0]

	switch command {
	case "version", "v":
		commands.RunVersion()

	case "health":
		cfg, _ := config.Load()
		if err := commands.RunHealth(cfg.BaseURL); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

	case "config":
		if len(args) < 2 {
			fmt.Println("Usage: megaverse config <set|get|list>")
			os.Exit(1)
		}
		subcmd := args[1]
		switch subcmd {
		case "set":
			if err := commands.RunConfigSet(args[2:]); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
			fmt.Println("Config updated")
		case "get":
			if err := commands.RunConfigGet(args[2:]); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
		case "list":
			if err := commands.RunConfigList(); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
		default:
			fmt.Printf("Unknown config command: %s\n", subcmd)
			os.Exit(1)
		}

	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println(`MegaVerse CLI

Usage:
  megaverse <command>

Commands:
  version     Show CLI version
  health      Check API health
  config      Manage configuration
    set       Set config value
    get       Get config value
    list      List all config

Examples:
  megaverse version
  megaverse health
  megaverse config set base-url http://localhost:8080
  megaverse config list`)
}
