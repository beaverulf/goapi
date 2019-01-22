package controllers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/beaverulf/goapi/encryption/aes"
	"github.com/beaverulf/goapi/encryption/utils"
	"github.com/beaverulf/goapi/models"
)

// swagger:route POST /aes aes GetAESEncryptionServicesList
// Lists all encryption modes for aes.
// responses:
//  200: swaggerCryptoServiceResp
//GetAESEncryptionServicesList shows the swagger index
func GetAESEncryptionServicesList(w http.ResponseWriter, r *http.Request) {
	var services []models.CryptoService
	services = append(services, models.CryptoService{Name: "AES", Functions: []string{"128-ECB", "192-ECB", "256-ECB"}})

	json.NewEncoder(w).Encode(services)
}

//EncryptAesECBHandlerFunc is an endpoint for encrypting with aes 128 ecb.
func EncryptAesECBHandlerFunc(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "Can't read body", http.StatusBadRequest)
		return
	}

	encryptionKey := r.Header.Get("EncryptionKey")

	if len(encryptionKey) > 32 {
		log.Printf("Keysize too large")
		http.Error(w, "Too large key", http.StatusBadRequest)
		return
	}
	paddedKey := utils.GetAESKeyPadding([]byte(encryptionKey))
	ciphertext, err := aes.EncryptAesEcb128(body, paddedKey)
	if err != nil {
		log.Printf("Encryption error: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var cipher = models.Cipher{
		Algorithm: "AES",
		Entropy:   fmt.Sprint(8 * len(paddedKey)),
		Mode:      "ECB"}

	var cipherObject = models.CipherObject{
		Cipher:     cipher,
		Key:        encryptionKey,
		Encoding:   "base64",
		Ciphertext: base64.StdEncoding.EncodeToString(ciphertext)}
	json.NewEncoder(w).Encode(cipherObject)
}

//DecryptAesECBHandlerFunc is an endpoint for encrypting with aes 128 ecb.
func DecryptAesECBHandlerFunc(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	key := r.Header.Get("EncryptionKey")
	paddedKey := utils.GetAESKeyPadding([]byte(key))

	//base64 decode
	log.Printf("decoded body len: %v", base64.StdEncoding.DecodedLen(len(body)))
	bodyDecoded := make([]byte, base64.StdEncoding.DecodedLen(len(body)))

	base64.StdEncoding.Decode(bodyDecoded, body)
	bodyPadded := utils.Pkcs7Padding(bodyDecoded, len(paddedKey))

	plaintext, err := aes.DecryptAesEcb128(bodyPadded, []byte(paddedKey))
	if err != nil {
		log.Printf("Error decrypting: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte(plaintext))
}
