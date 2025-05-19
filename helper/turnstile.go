package helper

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

func VerifyTurnstile(token, ip string) bool {
	body := map[string]string{
		"secret":   os.Getenv("TURNSTILE_SECRET_KEY"),
		"response": token,
		"remoteip": ip,
	}
	jsonBody, _ := json.Marshal(body)

	resp, err := http.Post("https://challenges.cloudflare.com/turnstile/v0/siteverify",
		"application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	var result struct {
		Success bool `json:"success"`
	}
	json.NewDecoder(resp.Body).Decode(&result)

	return result.Success
}
