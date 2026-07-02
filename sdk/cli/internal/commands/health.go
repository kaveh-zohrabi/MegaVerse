package commands

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func RunHealth(baseURL string) error {
	client := &http.Client{Timeout: 5 * time.Second}
	
	resp, err := client.Get(baseURL + "/health")
	if err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	return nil
}
