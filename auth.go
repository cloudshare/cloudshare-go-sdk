package cloudshare

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func generateToken() string {
	return randSeq(10)
}

func hash(text string) string {
	s := sha1.Sum([]byte(text))
	return hex.EncodeToString(s[:])
}

func authToken(apiKey string, apiID string, url string) string {
	timestamp := time.Now().Unix()
	token := generateToken()
	hmac := hash(fmt.Sprintf("%s%s%d%s", apiKey, url, timestamp, token))
	ret := fmt.Sprintf("userapiid:%s;timestamp:%d;token:%s;hmac:%s",
		apiID, timestamp, token, hmac)
	return ret
}
