package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	log.Printf("%s\n", msg)
	type errResp struct {
		Error string `json:"error"`
	}
	respondWithJson(w, code, errResp{
		Error: msg,
	})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error: failed marshalling to json: %s", err)
		w.WriteHeader(500)
		return
	}
	w.Write(data)
}

func decodeRequestParams(r *http.Request, param any) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&param); err != nil {
		return err
	}
	return nil
}

func reverse[T any](s []T) []T {
	for i := 0; i < len(s)/2; i++ {
		temp := s[i]
		s[i] = s[len(s)-i-1]
		s[len(s)-i-1] = temp
	}
	return s
}

func contains[T comparable](item T, array []T) bool {
	for _, e := range array {
		if item == e {
			return true
		}
	}
	return false
}

func mapArray[T, K any](a []T, f func(T) K) []K {
	r := make([]K, 0, len(a))
	for _, elem := range a {
		r = append(r, f(elem))
	}
	return r
}

func createApiKey() string {
	b := make([]byte, sha256.BlockSize)
	hash := sha256.Sum256(b)
	return hex.EncodeToString(hash[:])
}

func getHeaderValue(header, prefix string) (value string, err error) {
	value, found := strings.CutPrefix(header, prefix)
	if !found {
		return "", errors.New("failed to retrieve header value")
	}
	return value, nil
}

// most common date layouts to try and parse
var dateLayouts = [...]string{
	time.Layout,
	time.UnixDate,
	time.ANSIC,
	time.RFC1123,
	time.RFC1123Z,
	time.RFC822,
	time.RFC822Z,
	time.RFC850,
}

func tryParseDate(date string) (time.Time, error) {
	for _, layout := range dateLayouts {
		parsedTime, err := time.Parse(layout, date)
		if err == nil {
			return parsedTime, nil
		}
	}
	return time.Time{}, errors.New(fmt.Sprintf("unrecogized time format: %s", date))
}
