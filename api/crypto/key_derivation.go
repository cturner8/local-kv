package crypto

import (
	"golang.org/x/crypto/argon2"
)

var (
	KEY_LENGTH uint32 = 32
	TIME       uint32 = 1
	MEMORY     uint32 = 64 * 1024
	THREADS    uint8  = 4
)

func DeriveKey(password, salt []byte) []byte {
	return argon2.Key(password, salt, TIME, MEMORY, THREADS, KEY_LENGTH)
}
