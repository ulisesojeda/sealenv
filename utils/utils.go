package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	b64 "encoding/base64"
	_ "encoding/hex"
	"fmt"
	"golang.org/x/term"
	"os"
	"strings"
)

func encrypt(plaintext string, password string) []byte {
	aes, err := aes.NewCipher([]byte(password))
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		panic(err)
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	return ciphertext
}

func decrypt(ciphertext []byte, password string) (string, bool) {
	aes, err := aes.NewCipher([]byte(password))
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		panic(err)
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
	if err != nil {
		return "INVALID_PASS", false
	}

	return string(plaintext), true
}

func DecryptVariable(variable string, password string) ([]string, bool) {
	name_value := strings.SplitN(variable, "=", 2)
	value_bin, _ := b64.StdEncoding.DecodeString(name_value[1])
	plaintext, ok := decrypt(value_bin, password)
	if ok {
		return []string{name_value[0], plaintext}, true
	} else {
		return []string{}, false
	}
}

func readPassword() string {
	stdin_fd := int(os.Stdin.Fd())
	state, _ := term.GetState(stdin_fd)
	term.MakeRaw(stdin_fd)
	t := term.NewTerminal(os.Stdin, "")
	pass, _ := t.ReadPassword("ENV Password: ")
	term.Restore(stdin_fd, state)
	return pass
}

func SealVariables(vars_file, out_file, password string) {
	raw_p, _ := os.ReadFile(vars_file)
	lines_p := strings.Split(string(raw_p), "\n")

	out_f, _ := os.Create(out_file)
	defer out_f.Close()

	for _, line_p := range lines_p {
		if len(line_p) > 0 {
			entry := strings.SplitN(line_p, "=", 2)
			key, value := entry[0], entry[1]

			cipher := encrypt(value, password)
			enc_cipher := b64.StdEncoding.EncodeToString(cipher)
			key_entry := fmt.Sprintf("%s=%s\n", key, enc_cipher)
			out_f.WriteString(key_entry)
		}
	}
	out_f.Sync()
}

func UnsealVariables(vars_file, password string) {
	raw_p, _ := os.ReadFile(vars_file)
	lines_p := strings.Split(string(raw_p), "\n")

	for _, line_p := range lines_p {
		if len(line_p) > 0 {
			entry := strings.SplitN(line_p, "=", 2)
			key, value := entry[0], entry[1]

			ciphertext, _ := b64.StdEncoding.DecodeString(value)
			if plain, ok := decrypt(ciphertext, password); ok {
				key_entry := fmt.Sprintf("%s=%s", key, plain)
				fmt.Println(key_entry)
			} else {
				fmt.Println(plain)
				os.Exit(1)
			}
		}
	}
}

func GetPassword() string {
	pass := readPassword()
	password := pass + (strings.Repeat("0", 32-len(pass)))
	return password
}
