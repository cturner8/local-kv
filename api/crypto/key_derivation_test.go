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

	if base64key != "mgExwmqkNdLiDYfakZOi8LL8ehUr9MTWxkVe3pYX9jc=" {
		t.Errorf("Key mismatch")
	}
}

func TestGenerateSalt(t *testing.T) {
	salt := GenerateSalt()
	if len(salt) != 32 {
		t.Errorf("Expected salt length of 32, got %d", len(salt))
	}
}
