package utopiapaylib

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"golang.org/x/crypto/curve25519"
)

func closeRequest(resp *http.Response) {
	if resp != nil {
		resp.Body.Close()
	}
}

func sendRequest(requestType, url string, requestData []byte) ([]byte, error) {
	req, err := http.NewRequest(requestType, url, bytes.NewBuffer(requestData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	defer closeRequest(resp)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}

func encryptX25519(message []byte) ([]byte, error) {
	// Generate a private key
	privateKey := make([]byte, scalarSize)
	_, err := io.ReadFull(rand.Reader, privateKey)
	if err != nil {
		return nil, fmt.Errorf("read private key error: %w", err)
	}

	// Calculate public key from private key
	publicKey, err := curve25519.X25519(privateKey, curve25519.Basepoint)
	if err != nil {
		return nil, fmt.Errorf("encode public key: %w", err)
	}

	// Choose a random nonce
	nonce := make([]byte, scalarSize)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, fmt.Errorf("read nonce: %w", err)
	}

	// Encrypt the message
	encryptedMessage, err := curve25519.X25519(message, publicKey)
	if err != nil {
		return nil, fmt.Errorf("encode message: %w", err)
	}

	return encryptedMessage, nil
}
