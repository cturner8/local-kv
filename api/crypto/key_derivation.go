package crypto

import (
	"crypto/rand"

	"golang.org/x/crypto/argon2"
)

var (
	KEY_LENGTH uint32 = 32
	TIME       uint32 = 3
	MEMORY     uint32 = 32 * 1024
	THREADS    uint8  = 4
)

func DeriveKey(password, salt []byte) []byte {
	return argon2.Key(password, salt, TIME, MEMORY, THREADS, KEY_LENGTH)
}

func GenerateSalt() []byte {
	salt := make([]byte, 32)
	_, err := rand.Read(salt)
	if err != nil {
		panic(err)
	}
	return salt
}
