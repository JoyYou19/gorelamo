package transport

import (
	"bytes"
	"encoding/xml"
	"io"
	"net/http"
)

func (h *HTTP) PostXML(url string, body []byte, out any) error {
	req, _ := http.NewRequest("POST", url, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/xml")
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
		return xml.Unmarshal(raw, out)
	}
	return nil
}
