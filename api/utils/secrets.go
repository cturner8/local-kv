package utils

import "os"

func ReadSecretFile(filePath string) []byte {
	secretFile, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return secretFile
}
