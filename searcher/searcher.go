package searcher

import (
	"QQHQMusic/song"
	"QQHQMusic/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func buildSearchContent(keyWord string) string {
	keyWord = strings.TrimSpace(keyWord)
	content := fmt.Sprintf(`{
		"comm": {"ct": "19", "cv": "1845"},
		"music.search.SearchCgiService":{
			"method": "DoSearchForQQMusicDesktop",
			"module": "music.search.SearchCgiService",
			"param": {"query": "%s", "num_per_page": %d, "page_num": 1}}}`,
		keyWord, utils.SearchCount)

	return content
}

func Search(keyWord string) []*song.Song {
	url := "https://u.y.qq.com/cgi-bin/musicu.fcg"
	data := buildSearchContent(keyWord)
	bs := bytes.NewBufferString(data)
	req, err := http.NewRequest(http.MethodPost, url, bs)
	req.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36")
	req.Header.Set("content-type", "application/json; charset=UTF-8")
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
		return nil
	}

	code := result["code"].(float64)
	if code != 0 {
		log.Fatalf("ERROR when parse search result, return code should be 0, current value is: %v \n", code)
	}

	s := result["music.search.SearchCgiService"].(map[string]interface{})
	d := s["data"].(map[string]interface{})
	b := d["body"].(map[string]interface{})
	so := b["song"].(map[string]interface{})
	li := so["list"].([]interface{})

	var songList []*song.Song
	for i := range li {
		songList = append(songList, song.New(li[i].(map[string]interface{})))
	}

	var filteredSongList []*song.Song

	length := len(songList)
	for i, s := range songList {
		fmt.Printf("过滤掉低品质: %d/%d...\r", i, length)
		if "" == s.BestFormat || 1024*1024*10 > s.BestSize || "" == s.BestFormatCode {
			continue
		}

		filteredSongList = append(filteredSongList, s)
	}

	utils.CurrentKeyWord = keyWord
	return filteredSongList
}
