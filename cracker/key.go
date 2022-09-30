package cracker

import (
	"log"
	"strings"
)

type key struct {
	QQ  string
	Key string
}

var TheKey = key{}

func init() {
	cookie, err := getCookie()

	if err != nil {
		log.Fatal("ERROR when get cookie : ", err.Error())
	}

	//log.Println("cookie: " + cookie)
	cookie = strings.Replace(cookie, "-", "", -1)
	cookie = strings.Replace(cookie, "|", "", -1)
	//log.Println("cookie clean: " + cookie)

	if len(cookie) < 10 {
		log.Fatal("COOKIE ERROR, to short, should >= 10")
	}

	if !strings.Contains(cookie, "%") {
		log.Fatal("COOKIE FORMAT ERROR, should contains the % char")
	}

	subs := strings.Split(cookie, "%")

	qq := DecryptDES(subs[1], subs[0][0:8])
	if len(qq) == 0 {
		log.Fatal("DECRYPT QQ ERROR")
	}
	if len(qq) < 8 {
		qq += "QMD"
	}

	key := DecryptDES(subs[0], qq[0:8])
	if len(key) == 0 {
		log.Fatal("DECRYPT KEY ERROR")
	}

	TheKey.QQ = qq
	TheKey.Key = key
}
