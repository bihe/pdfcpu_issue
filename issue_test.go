package main

import (
	"io/ioutil"
	"testing"
)

func Test_PDF_Encryption(t *testing.T) {
	payload, err := ioutil.ReadFile("./unencrypted.pdf")
	if err != nil {
		t.Fatalf("cannot read payload: %v", err)
	}

	err = encryptPdfPayload("12345", payload)
	if err != nil {
		t.Errorf("could not encrypt PDF payload with password '12345': %v", err)
	}
}
