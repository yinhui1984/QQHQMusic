package cracker

import (
	"encoding/base64"
	DD "github.com/sunmi-OS/gocore/encryption/des"
	"log"
)

func EncryptDES(text, key string) string {
	e, err := DD.DesEncrypt(text, key, key)
	if err != nil {
		log.Fatal("ERROR when EncryptDES: " + err.Error())
		return ""
	}

	b64 := base64.StdEncoding.EncodeToString([]byte(e))
	return b64
}

func DecryptDES(base64Text, key string) string {
	bs, err := base64.StdEncoding.DecodeString(base64Text)
	if err != nil {
		log.Fatal("ERROR when DecryptDES: " + err.Error())
		return ""
	}
	str := string(bs)
	decrypt, err := DD.DesDecrypt(str, key, key)
	if err != nil {
		return ""
	}

	return decrypt
}
