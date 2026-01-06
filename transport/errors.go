package transport

import (
	"encoding/base64"
	"encoding/json"
	"github.com/JoyYou19/gorelamo/errors"
	"strings"
)

func parseDBError(raw []byte) error {
	if len(raw) == 0 {
		return errors.ErrUnknown
	}

	// 1. Try JSON first
	var jsonErr map[string]any
	if err := json.Unmarshal(raw, &jsonErr); err == nil {
		// Common patterns: { "error": "..."} or { "message": "..." }
		if msg, ok := jsonErr["error"].(string); ok {
			return &errors.DBError{
				Message: msg,
				Raw:     raw,
			}
		}
		if msg, ok := jsonErr["message"].(string); ok {
			return &errors.DBError{
				Message: msg,
				Raw:     raw,
			}
		}
	}

	// 2. Try base64 decode (your license error case)
	if decoded, err := tryBase64(raw); err == nil {
		return &errors.DBError{
			Message: decoded,
			Raw:     raw,
		}
	}

	// 3. Fallback: plain text
	return &errors.DBError{
		Message: strings.TrimSpace(string(raw)),
		Raw:     raw,
	}
}

func tryBase64(raw []byte) (string, error) {
	s := strings.TrimSpace(string(raw))

	// Cheap heuristic: base64 strings are long and don't contain spaces
	if strings.ContainsAny(s, " \n\t") {
		return "", errors.ErrUnknown
	}

	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(decoded)), nil
}
