package main

import (
	"bytes"
	"fmt"

	pdfApi "github.com/pdfcpu/pdfcpu/pkg/api"
	pdf "github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

func encryptPdfPayload(newPwd string, payload []byte) error {
	var (
		conf *pdf.Configuration
		err  error
	)

	conf = pdf.NewRC4Configuration(newPwd, newPwd, 40)
	conf.Cmd = pdf.ENCRYPT
	conf.UserPWNew = &newPwd

	buff := new(bytes.Buffer)
	if err = pdfApi.Optimize(bytes.NewReader(payload), buff, conf); err != nil {
		return fmt.Errorf("could not process PDF payload: %v", err)
	}

	return nil
}
