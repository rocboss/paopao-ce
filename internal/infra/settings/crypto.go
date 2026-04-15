// Copyright 2026 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package settings

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"strings"

	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/pkg/xerror"
)

const encryptedValuePrefix = "enc:v1:"

type secretCodec struct {
	key []byte
}

func newSecretCodec() *secretCodec {
	if conf.AdminSettingsSetting == nil {
		return &secretCodec{}
	}
	seed := strings.TrimSpace(conf.AdminSettingsSetting.EncryptionKey)
	if seed == "" {
		return &secretCodec{}
	}
	sum := sha256.Sum256([]byte(seed))
	return &secretCodec{key: sum[:]}
}

func (c *secretCodec) enabled() bool {
	return len(c.key) == 32
}

func (c *secretCodec) Encrypt(value string) (string, error) {
	if value == "" {
		return "", nil
	}
	if !c.enabled() {
		return "", xerror.ServerError.WithDetails("admin settings encryption key is not configured")
	}
	block, err := aes.NewCipher(c.key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	ciphertext := gcm.Seal(nil, nonce, []byte(value), nil)
	payload := append(nonce, ciphertext...)
	return encryptedValuePrefix + base64.StdEncoding.EncodeToString(payload), nil
}

func (c *secretCodec) Decrypt(value string) (string, error) {
	if value == "" {
		return "", nil
	}
	if !strings.HasPrefix(value, encryptedValuePrefix) {
		return value, nil
	}
	if !c.enabled() {
		return "", xerror.ServerError.WithDetails("admin settings encryption key is not configured")
	}
	raw, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(value, encryptedValuePrefix))
	if err != nil {
		return "", fmt.Errorf("decode encrypted setting: %w", err)
	}
	block, err := aes.NewCipher(c.key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	if len(raw) < gcm.NonceSize() {
		return "", xerror.ServerError.WithDetails("invalid encrypted setting payload")
	}
	nonce, ciphertext := raw[:gcm.NonceSize()], raw[gcm.NonceSize():]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}
