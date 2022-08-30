package cracker

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"strconv"
)

func getCookie() (string, error) {

	uid := "822a3b85-a5c9-438e-a277-a8da412e8265"
	systemVersion := "1.7.2"
	versionCode := "76"
	deviceBrand := "360"
	deviceModel := "QK1707-A01"
	appVersion := "7.1.2"

	toEncrypt := uid + deviceModel + deviceBrand + systemVersion + appVersion + versionCode
	key := ("QMD" + "F*ckYou!")[0:8] //你没看错, 密码就是这么粗鲁!  https://github.com/QiuChenly/QQFlacMusicDownloader/blob/main/DecompileFiles/main.java#L32
	encIP := EncryptDES(toEncrypt, key)

	response, err := http.Post("http://8.136.185.193/api/Cookies",
		"application/json;  charset=UTF-8",
		bytes.NewBufferString(`{"appVersion":"7.1.2","deviceBrand":"360","deviceModel":"QK1707-A01","ip":"`+
			encIP+
			`","systemVersion":"1.7.2","uid":"822a3b85-a5c9-438e-a277-a8da412e8265","versionCode":"76"}`))

	if err != nil {
		return "", errors.New("POST ERROR when getCookie:" + err.Error())
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	if response.StatusCode != 200 {
		return "", errors.New("RESPONSE ERROR when getCookie:" + strconv.Itoa(response.StatusCode))
	}

	body, err := io.ReadAll(response.Body) //ioutil.ReadAll(response.Body)
	if err != nil {
		return "", errors.New("READ ERROR when getCookie:" + err.Error())
	}

	return string(body), nil

}
