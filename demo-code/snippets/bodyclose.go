package snippets

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

// START BROKEN // OMIT
func httpRequest(ctx context.Context, hc *http.Client) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://httpbin.org/status/500", nil)
	if err != nil {
		return fmt.Errorf("http request: new request: %w", err)
	}

	resp, err := hc.Do(req)
	if err != nil {
		return fmt.Errorf("http request: execute request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http request: unexpected status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("http request: read response body: %w", err)
	}

	fmt.Println(string(data))

	return nil
}

// END BROKEN // OMIT

func httpRequestFixed(ctx context.Context, hc *http.Client) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://httpbin.org/status/500", nil)
	if err != nil {
		return fmt.Errorf("http request: new request: %w", err)
	}

	// START FIXED // OMIT
	resp, err := hc.Do(req)
	if err != nil {
		return fmt.Errorf("http request: execute request: %w", err)
	}
	defer func() { // HL
		_, _ = io.Copy(io.Discard, resp.Body) // HL
		_ = resp.Body.Close()                 // HL
	}()

	if resp.StatusCode != http.StatusOK {
		// ...
		// END FIXED // OMIT
		return fmt.Errorf("http request: unexpected status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("http request: read response body: %w", err)
	}

	fmt.Println(string(data))

	return nil
}
