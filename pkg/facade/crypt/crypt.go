package crypt

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/dchest/uniuri"
	"github.com/forgoer/openssl"
	"github.com/lucidfy/lucid/pkg/errors"
	"github.com/techoner/gophp/serialize"
)

// https://developpaper.com/go-implements-laravels-encrypt-and-decrypt-methods/

//Encryption
func Encrypt(value interface{}) (string, *errors.AppError) {
	iv := make([]byte, 16)
	_, err := rand.Read(iv)
	if err != nil {
		return "", errors.InternalServerError("rand.Read() error", err)
	}

	//Deserialization
	message, err := serialize.Marshal(value)
	if err != nil {
		return "", errors.InternalServerError("serialize.Marshal() error", err)
	}

	key := getKey()

	//Encryptionvalue
	res, err := openssl.AesCBCEncrypt(message, []byte(key), iv, openssl.PKCS7_PADDING)
	if err != nil {
		return "", errors.InternalServerError("openssl.AesCBCEncrypt() error", err)
	}

	//Base64 encryption
	resVal := base64.StdEncoding.EncodeToString(res)
	resIv := base64.StdEncoding.EncodeToString(iv)

	//Generate MAC value
	data := resIv + resVal
	mac := computeHmacSha256(data, key)

	//Construct ticket structure
	ticket := make(map[string]interface{})
	ticket["iv"] = resIv
	ticket["mac"] = mac
	ticket["value"] = resVal

	//JSON serialization
	resTicket, err := json.Marshal(ticket)
	if err != nil {
		return "", errors.InternalServerError("json.Marshal() error", err)
	}
	//Base64 encryptionticket
	ticketR := base64.StdEncoding.EncodeToString(resTicket)

	return ticketR, nil
}

//Decryption
func Decrypt(value string) (string, *errors.AppError) {
	//Base64 decryption
	token, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return "", errors.InternalServerError("base64.StdEncoding.DecodeString() error", err)
	}

	//JSON deserialization
	tokenJson := make(map[string]string)
	err = json.Unmarshal(token, &tokenJson)
	if err != nil {
		return "", errors.InternalServerError("json.Unmarshal() error", err)
	}

	tokenJsonIv, okIv := tokenJson["iv"]
	tokenJsonValue, okValue := tokenJson["value"]
	tokenJsonMac, okMac := tokenJson["mac"]
	if !okIv || !okValue || !okMac {
		return "", errors.InternalServerError("crypt.Decrypt() error", fmt.Errorf("value is not complete"))
	}

	key := getKey()

	//Mac check to prevent data tampering
	data := tokenJsonIv + tokenJsonValue
	check := checkMAC(data, tokenJsonMac, key)
	if !check {
		return "", errors.InternalServerError("crypt.Decrypt() error", fmt.Errorf("mac is invalid"))
	}

	//Base64 decryptioniv和value
	tokenIv, err := base64.StdEncoding.DecodeString(tokenJsonIv)
	if err != nil {
		return "", errors.InternalServerError("base64.StdEnconding.DecodeString() error", err)
	}
	tokenValue, err := base64.StdEncoding.DecodeString(tokenJsonValue)
	if err != nil {
		return "", errors.InternalServerError("base64.StdEnconding.DecodeString() error", err)
	}
	//AES decryption value
	dst, err := openssl.AesCBCDecrypt(tokenValue, []byte(key), tokenIv, openssl.PKCS7_PADDING)
	// fmt.Println("dst", string(dst))
	if err != nil {
		return "", errors.InternalServerError("openssl.AesCBCDecrypt() error", err)
	}

	//Deserialization
	res, err := serialize.UnMarshal(dst)
	if err != nil {
		return "", errors.InternalServerError("serialize.UnMarshal() error", err)
	}
	return res.(string), nil
}

//Compare the expected hash with the actual hash
func checkMAC(message, msgMac, secret string) bool {
	expectedMAC := computeHmacSha256(message, secret)
	// fmt.Println(expectedMAC, msgMac)
	return hmac.Equal([]byte(expectedMAC), []byte(msgMac))
}

//Calculate MAC value
func computeHmacSha256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

//Processing key
func getKey() string {
	appKey := os.Getenv("APP_KEY")
	if strings.HasPrefix(appKey, "base64:") {
		split := appKey[7:]
		if key, err := base64.StdEncoding.DecodeString(split); err == nil {
			return string(key)
		}
		return split
	}
	return appKey
}

func GenerateRandomString(n int) string {
	return uniuri.NewLen(n)
}
