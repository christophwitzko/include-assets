package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
)

type AssetObject string

func (ao AssetObject) LoadAsString() (string, error) {
	data, gzipErr := gzip.NewReader(base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(string(ao))))
	if gzipErr != nil {
		return "", gzipErr
	}
	defer data.Close()
	output, err := ioutil.ReadAll(data)
	if err != nil {
		return "", err
	}
	return string(output), nil
}
