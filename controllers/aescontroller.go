package controllers

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/beaverulf/goapi/encryption/aes"
	"github.com/beaverulf/goapi/encryption/utils"
	"github.com/beaverulf/goapi/types"
)

//EncryptAes128ECBHandler is an endpoint for encrypting with aes 128 ecb.
func EncryptAes128ECBHandlerFunc(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "Can't read body", http.StatusBadRequest)
		return
	}

	encryptionKey := r.Header.Get("EncryptionKey")

	ciphertext, err := aes.EncryptAesEcb128(body, []byte(encryptionKey))
	if err != nil {
		log.Printf("Encryption error: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var cipher = types.Cipher{Algorithm: "AES", Bits: "128", Mode: "ECB"}

	var cipherObject = types.CipherObject{Cipher: cipher, Key: encryptionKey, Encoding: "base64", Ciphertext: base64.StdEncoding.EncodeToString(ciphertext)}
	json.NewEncoder(w).Encode(cipherObject)

}

//DecryptAes128ECBHandlerFunc is an endpoint for encrypting with aes 128 ecb.
func DecryptAes128ECBHandlerFunc(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	key := r.Header.Get("EncryptionKey")

	//base64 decode
	log.Printf("decoded body len: %v", base64.StdEncoding.DecodedLen(len(body)))
	bodyDecoded := make([]byte, base64.StdEncoding.DecodedLen(len(body)))

	base64.StdEncoding.Decode(bodyDecoded, body)
	bodyPadded := utils.Pkcs7Padding(bodyDecoded, len(key))

	plaintext, err := aes.DecryptAesEcb128(bodyPadded, []byte(key))
	if err != nil {
		log.Printf("Error decrypting: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte(plaintext))
}
