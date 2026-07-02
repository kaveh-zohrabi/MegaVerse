package commands

import (
	"fmt"
	"megaverse.dev/sdk/cli/internal/config"
)

func RunConfigSet(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: megaverse config set <key> <value>")
	}

	cfg, err := config.Load()
	if err != nil {
		return err
	}

	switch args[0] {
	case "api-key":
		cfg.APIKey = args[1]
	case "base-url":
		cfg.BaseURL = args[1]
	default:
		return fmt.Errorf("unknown config key: %s", args[0])
	}

	return config.Save(cfg)
}

func RunConfigGet(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: megaverse config get <key>")
	}

	cfg, err := config.Load()
	if err != nil {
		return err
	}

	switch args[0] {
	case "api-key":
		fmt.Println(cfg.APIKey)
	case "base-url":
		fmt.Println(cfg.BaseURL)
	default:
		return fmt.Errorf("unknown config key: %s", args[0])
	}

	return nil
}

func RunConfigList() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	fmt.Printf("api-key:  %s\n", cfg.APIKey)
	fmt.Printf("base-url: %s\n", cfg.BaseURL)
	return nil
}
