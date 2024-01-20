package crypto

import (
	"encoding/base64"
	"testing"
)

func TestGenerateKey(t *testing.T) {
	key := DeriveKey([]byte("password"), []byte("salt"))
	base64key := base64.StdEncoding.EncodeToString(key)

	if len(key) != 32 {
		t.Errorf("Expected key length of 32, got %d", len(key))
	}

	if base64key != "E804mZBsD1m5OQ4G8McMuQD2F0sVQD1vVdSPaOaz5+I=" {
		t.Errorf("Key mismatch")
	}
}
