package main

import (
	"testing"
)

func TestLoadCorrect(t *testing.T) {
	ao := AssetObject("H4sIAAAAAAAA/8pIzcnJVyjPL8pJ4QIEAAD//y07CK8MAAAA")
	data, err := ao.LoadAsString()
	if data != "hello world\n" || err != nil {
		t.Fatal("invalid string")
	}
}

func TestLoadIncorrect(t *testing.T) {
	ao := AssetObject("H4sIAAAAAAAA/8p")
	data, err := ao.LoadAsString()
	if data != "" || err == nil {
		t.Fatal("valid string")
	}
}
