package utils

import (
	"io/ioutil"
	"os"
)

func ReadFile(path string) ([]byte, error) {
	fd, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	r, err := ioutil.ReadAll(fd)
	if err != nil {
		return nil, err
	}
	return r, nil
}