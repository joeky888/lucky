package oor

import (
	"bytes"
	"github.com/helloh2o/lucky/cmm/utils"
	"github.com/helloh2o/lucky/log"
	"testing"
)

func TestNewXORCipher(t *testing.T) {
	cipher := NewXORCipher(utils.RandString(10))
	painText := bytes.Repeat([]byte("hello world �rrrr�"), 1)
	encrypt := cipher.Encode(painText)
	log.Debug(string(encrypt))
	dencrypt := cipher.Decode(encrypt)
	log.Debug(string(dencrypt))
}
