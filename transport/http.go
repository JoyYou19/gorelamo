package transport

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"gorelamo/config"
	"io"
	"net/http"
)

type HTTP struct {
	client *http.Client
	auth   string
}

func (h *HTTP) Post(url string, body any, out any) error {
	data, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", url, bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", h.auth)

	resp, err := h.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	raw, _ := io.ReadAll(resp.Body)

	if resp.StatusCode >= 400 {
		return parseDBError(raw)
	}

	if out != nil {
		return json.Unmarshal(raw, out)
	}
	return nil
}

func NewHTTP(cfg *config.Config) *HTTP {
	var auth string
	if cfg.Username != "" {
		auth = "Basic " + base64.StdEncoding.EncodeToString(
			[]byte(cfg.Username+":"+cfg.Password),
		)
	}

	return &HTTP{
		client: cfg.HTTPClient,
		auth:   auth,
	}
}
