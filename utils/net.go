package utils

import "net/http"

func CheckURLLiveness(url string) error {
	_, err := http.Get(url)
	return err
}