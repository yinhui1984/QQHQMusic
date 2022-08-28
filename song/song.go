package song

import (
	"QQHQMusic/cracker"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

type Song struct {
	Title          string
	Singer         string
	MID            string
	SongMID        string
	BestFormatCode string
	BestFormat     string
	Album          string
	BestSize       float64
	downloadLink   string
}

func (s *Song) GetDownloadLink() string {
	if "" == s.downloadLink {
		s.downloadLink = s.getDownloadLinkInner()
	}
	return s.downloadLink
}

func (s *Song) Play() {
	//log.Println("TODO: play: " + s.Title)
	//log.Println(s.GetDownloadLink())

	link := s.GetDownloadLink()
	log.Printf("准备播放 %s-%s, 获取链接:   %s\n", s.Title, s.Singer, link)

	//ffplay -hide_banner
	cmd := exec.Command("ffplay", link)
	err := cmd.Run()
	if err != nil {
		log.Println("ERROR when play: " + err.Error())
	}
}

func (s *Song) getDownloadLinkInner() string {
	link := s.getDownloadLinkFirstTry()
	if strings.Contains(link, `"title":"Not Found"`) {
		link = s.getDownloadLinkSecondTry()
	}

	return link
}

func (s *Song) getDownloadLinkFirstTry() string {
	fileName := s.BestFormatCode + s.MID + "." + s.BestFormat
	key := ("QMD" + cracker.TheKey.QQ)[0:8]
	en := cracker.EncryptDES(fileName, key)
	data := fmt.Sprintf(`"%s"`, en)
	response, err := http.Post("http://8.136.185.193/api/MusicLink/link",
		"application/json;charset=utf-8", bytes.NewBufferString(data))
	if err != nil {
		log.Println("Error when GetDownloadLink: " + err.Error())
		return ""
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	all, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Error when GetDownloadLink, parse body: " + err.Error())
		return ""
	}
	return string(all)
}

func (s *Song) getDownloadLinkSecondTry() string {
	vkey := ""
	fileName := s.BestFormatCode + s.MID + "." + s.BestFormat
	songMID := s.SongMID
	url := "https://u.y.qq.com/cgi-bin/musicu.fcg"
	data := fmt.Sprintf(`{
			"comm": {"ct": "19", "cv": "1777"},
			"queryvkey": {"method": "CgiGetVkey",
							"module": "vkey.GetVkeyServer",
							"param": {
                            "uin": "%s",
                            "guid": "QMD50",
                            "referer": "y.qq.com",
                            "songtype": [1],
							"filename": ["%s"],
                            "songmid": ["%s"]}}}`, cracker.TheKey.QQ, fileName, songMID)
	bs := bytes.NewBufferString(data)

	cookie := fmt.Sprintf("qqmusic_key=%s;qqmusic_uin=%s;", cracker.TheKey.Key, cracker.TheKey.QQ)

	req, err := http.NewRequest(http.MethodPost, url, bs)
	req.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36")
	req.Header.Set("content-type", "application/json; charset=UTF-8")
	req.Header.Set("cookie", cookie)
	req.Header.Set("referer", "https://y.qq.com/portal/profile.html")
	if err != nil {
		log.Fatal("ERROR when make http request: " + err.Error())
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	response, err := client.Do(req)
	if err != nil {
		log.Fatal("ERROR when do search request: " + err.Error())
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	var result map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		log.Fatal("ERROR when decode search result: " + err.Error())
		return ""
	}

	code := result["code"].(float64)
	if code != 0 {
		log.Printf("ERROR when parse search result, return code should be 0, current value is: %v \n", code)
	}

	a := result["queryvkey"].(map[string]interface{})
	b := a["data"].(map[string]interface{})
	c := b["midurlinfo"].([]interface{})
	d := c[0].(map[string]interface{})
	vkey = d["purl"].(string)

	return fmt.Sprintf(`http://ws.stream.qqmusic.qq.com/%s&fromtag=140`, vkey)
}

func New(i map[string]interface{}) *Song {
	s := &Song{}

	//title
	title := i["title"].(string)
	s.Title = fixSongChars(title)

	//mid
	f := i["file"].(map[string]interface{})
	s.MID = f["media_mid"].(string)

	//song mid //获取下载SecondTry时使用
	s.SongMID = i["mid"].(string)

	//best
	s.BestFormatCode, s.BestFormat, s.BestSize = getBestFormat(f)

	//singer
	a := i["singer"].([]interface{})
	b := a[0].(map[string]interface{})
	singer := b["name"].(string)
	s.Singer = fixSongChars(singer)

	//album
	x := i["album"].(map[string]interface{})
	album := x["title"].(string)
	album = strings.TrimSpace(album)
	if "" == album {
		album = "未分类专辑"
	}
	s.Album = fixSongChars(album)

	return s
}

func fixSongChars(str string) string {
	r := strings.NewReplacer(
		"|", ",",
		"/", "-",
		"╲", "-",
		"、", "·",
		"“", "\"",
		"”", "\"",
		"*", "x",
		"？", "?",
		"《", "[",
		"》", "]",
		"【", "[",
		"】", "]",
		"’", "'",
		" ", "")

	return r.Replace(str)
}

// 获取最佳格式, 返回 格式代码, 格式名称, 文件大小
// 比如 "RS01", "flac", "1234555"
// 传入参数是 歌曲map的"file"
func getBestFormat(f map[string]interface{}) (string, string, float64) {
	size := f["size_hires"].(float64)
	if size != 0 {
		return "RS01", "flac", size
	}

	size = f["size_flac"].(float64)
	if size != 0 {
		return "F000", "flac", size
	}

	size = f["size_320mp3"].(float64)
	if size != 0 {
		return "M800", "mp3", size
	}

	size = f["size_192ogg"].(float64)
	if size != 0 {
		return "O600", "ogg", size
	}

	size = f["size_128mp3"].(float64)
	if size != 0 {
		return "M500", "mp3", size
	}

	//其它低品质, 忽略

	return "", "", 0
}
